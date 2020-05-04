package student

import (
	. "experiment/handler"
	"experiment/model"
	"experiment/pkg/constvar"
	"experiment/pkg/errno"
	"experiment/pkg/message"
	"experiment/pkg/redis"
	"experiment/service"
	"experiment/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"sync"
	"time"
)

type SubmitRequest struct {
	GroupId   int    `json:"group_id" binding:"required"`
	ProblemId int    `json:"problem_id" binding:"required"`
	Submit    string `json:"submit" binding:"required"`
}

func MyExperiments(c *gin.Context) {
	studentId, exists := c.Get("userId")
	if !exists {
		SendResponse(c, errno.ErrParam, nil)
		return
	}
	fmt.Println(studentId)

	se := model.StudentExperimentModel{}
	groups, err := se.GetGroups(int(studentId.(float64)))
	if err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	res := make([]model.StudentExperiment, len(groups))

	wg := sync.WaitGroup{}
	for i, group := range groups {
		wg.Add(1)
		go func(index int, se model.StudentExperimentModel) {
			defer wg.Done()
			experiment := model.ExperimentModel{}
			if err := experiment.GetDetail(se.GroupId); err != nil {
				SendResponse(c, errno.ErrDatabase, nil)
				return
			}
			res[index] = model.StudentExperiment{
				GroupId:   se.GroupId,
				GroupName: experiment.GroupName,
				Status:    se.Status,
				Score:     se.Score,
			}
		}(i, group)
	}
	wg.Wait()

	SendResponse(c, errno.OK, res)
}

func ProblemSubmit(c *gin.Context) {
	var request SubmitRequest
	if err := c.Bind(&request); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	if len(request.Submit) == 0 {
		SendResponse(c, errno.ErrParam, nil)
		return
	}

	studentId := util.GetUserId(c)
	if studentId == constvar.EMPTY {
		SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}

	var err error
	answer := model.AnswerModel{}
	if err = answer.Detail(request.GroupId, studentId, request.ProblemId); err != nil {
		answer.GroupId = request.GroupId
		answer.ProblemId = request.ProblemId
		answer.Status = constvar.ProblemSubmitStatus[constvar.RUNNING]
		answer.StudentId = studentId
		answer.Submit = request.Submit
		answer.UpdateTime = time.Now().Unix()
		err = answer.Create()
	} else {
		err = answer.Update(map[string]interface{}{
			"Submit":     request.Submit,
			"Status":     constvar.ProblemSubmitStatus[constvar.RUNNING],
			"UpdateTime": time.Now().Unix(),
		}, answer.Id)
	}

	if err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	//TODO 提交至判题队列中进行判题
	msg := message.TopicAnswerMessage{
		AnswerId:   answer.Id,
		ProblemId:  answer.ProblemId,
		Submit:     answer.Submit,
		UpdateTime: answer.UpdateTime,
	}
	realMsg, err := util.MsgEncode(msg)
	if err != nil {
		SendResponse(c, errno.ErrJsonMarshal, nil)
		return
	}
	client := message.GetKafkaClient()
	err = client.Produce(message.TopicAnswer, realMsg)
	if err != nil {
		SendResponse(c, errno.ErrSendMsgFail, nil)
		return
	}

	// 可以先返回，目前整个请求是同步，需要等待发送到 mq 成功为止
	SendResponse(c, errno.OK, map[string]int{
		"run_id": answer.Id,
	})
}

func GetStatus(c *gin.Context) {
	idSrc := c.Param("id")
	if len(idSrc) == 0 {
		SendResponse(c, errno.ErrParam, nil)
		return
	}

	id, err := strconv.Atoi(idSrc)
	if err != nil {
		SendResponse(c, errno.ErrParam, nil)
		return
	}

	userId := util.GetUserId(c)
	if userId == constvar.EMPTY {
		SendResponse(c, errno.ErrAuthority, nil)
		return
	}
	//TODO 可以将 run_id - userId, status 关系存储在缓存中，避免权限校验过程需要读取数据库 Done
	data, err := getRunIdStatus(id, userId)
	SendResponse(c, err, data)
	return
}

func GetProblemDetail(c *gin.Context) {
	grouIdSrc := c.Query("group")
	problemSrc := c.Query("problem")
	if len(grouIdSrc) == 0 || len(problemSrc) == 0 {
		SendResponse(c, errno.ErrParam, nil)
		return
	}

	problemId, err := strconv.Atoi(problemSrc)
	grouId, err := strconv.Atoi(grouIdSrc)
	if err != nil {
		SendResponse(c, errno.ErrParam, nil)
		return
	}

	userId := util.GetUserId(c)
	if userId == constvar.EMPTY {
		SendResponse(c, errno.ErrUserNotFound, nil)
	}

	answer := model.AnswerModel{}
	if err = answer.Detail(grouId, userId, problemId); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	// 提交的答案运行中，无法查看详情
	if answer.Status == constvar.ProblemSubmitStatus[constvar.RUNNING] {
		SendResponse(c, errno.ErrSubmitRunning, nil)
		return
	}

	SendResponse(c, errno.OK, map[string]interface{}{
		"submit":  answer.Submit,
		"correct": answer.Correct,
		"error":   answer.Error,
		"score":   answer.Score,
	})

}

// 提交实验
func ExperimentSubmit(c *gin.Context) {
	groupIdSrc := c.Param("id")
	if len(groupIdSrc) == 0 {
		SendResponse(c, errno.ErrParam, nil)
		return
	}

	groupId, err := strconv.Atoi(groupIdSrc)
	if err != nil {
		SendResponse(c, errno.ErrParam, nil)
		return
	}

	userId := util.GetUserId(c)
	if userId == constvar.EMPTY {
		SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}

	student := service.NewStudent()
	res, err := student.Submit(groupId, userId)
	if err != errno.OK {
		SendResponse(c, err, nil)
		return
	}

	SendResponse(c, errno.OK, res)
}

func getRunIdStatus(runId, userId int) (map[string]interface{}, error) {
	key := redis.GetRunIdStatusKey(runId, userId)
	data := make(map[string]interface{})
	statusSrc := redis.Client.Get(key)
	status, err := strconv.Atoi(statusSrc)
	if err == nil && status == constvar.ProblemSubmitStatus[constvar.RUNNING] {
		data["status"] = status
		data["correct"] = nil
		data["error"] = nil
		data["score"] = nil
		return data, nil
	} else {
		// 缓存失效, 查数据库拿数据
		answer := model.AnswerModel{}
		if err := answer.DetailById(runId); err != nil {
			return nil, errno.ErrDatabase
		}
		data["status"] = answer.Status
		data["correct"] = answer.Correct
		data["error"] = answer.Error
		data["score"] = answer.Score
		// 更新缓存，设置过期时间，避免更新数据库删除缓存操作失败
		err = redis.Client.Set(key, fmt.Sprintf("%d", answer.Status), time.Minute*5)
	}
	return data, nil
}
