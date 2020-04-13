package experiment

import (
	"encoding/json"
	. "experiment/handler"
	"experiment/model"
	"experiment/pkg/errno"
	"experiment/service"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"strconv"
	"sync"
)

type createRequest struct {
	Problems  []int  `json:"problems" binding:"required"`
	GroupName string `json:"group_name" binding:"required"`
}

type ListResponse struct {
	GroupId   int          `json:"group_id"`
	GroupName string       `json:"group_name"`
	List      []ProblemTag `json:"list"`
}

type ProblemTag struct {
	ProblemId int    `json:"problem_id"`
	Title     string `json:"title"`
}

type DistributedRequest struct {
	GroupId   int      `json:"group_id" binding:"required"`
	ClassList []string `json:"class_list" binding:"required"`
}

// 教师创建实验
func Create(c *gin.Context) {
	var request createRequest
	if err := c.Bind(&request); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		log.Error("", err)
		return
	}

	userName, _ := c.Get("username")
	problems, err := json.Marshal(request.Problems)
	if err != nil {
		SendResponse(c, errno.ErrJsonMarshal, nil)
		return
	}
	experiment := model.ExperimentModel{
		GroupName: request.GroupName,
		Poster:    userName.(string),
		Problems:  problems,
	}

	if err = experiment.Validate(); err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		log.Error("", err)
		return
	}

	if err = experiment.Create(); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	SendResponse(c, errno.OK, map[string]interface{}{
		"group_id": experiment.GroupId,
	})
}

func ProblemList(c *gin.Context) {
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

	e := model.ExperimentModel{}
	if err := e.Search([]string{"problems", "groupName"}, id); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	problems := make([]int, 0)
	err = json.Unmarshal(e.Problems, &problems)
	if err != nil {
		SendResponse(c, errno.ErrJsonUnmarshal, nil)
		log.Error("", err)
		return
	}

	list := make([]ProblemTag, len(problems))

	// 并行化，提高响应速度
	w := sync.WaitGroup{}
	for index, problemId := range problems {
		w.Add(1)
		go func(index, problemId int) {
			defer w.Done()
			p := model.ProblemModel{}
			if err := p.Search([]string{"title"}, problemId); err != nil {
				// 未找到
				return
			}
			list[index] = ProblemTag{
				ProblemId: problemId,
				Title:     p.Title,
			}
		}(index, problemId)
	}
	w.Wait()

	SendResponse(c, errno.OK, ListResponse{
		GroupId:   id,
		GroupName: e.GroupName,
		List:      list,
	})
}

func ClassDetail(c *gin.Context) {
	classId := c.Query("classId")
	groupIdSrc := c.Query("groupId")
	if len(classId) == 0 || len(groupIdSrc) == 0 {
		SendResponse(c, errno.ErrBind, nil)
		return
	}
	groupId, _ := strconv.Atoi(groupIdSrc)
	// 获取改班级学生列表
	sc := model.StudentClassModel{}
	studentList, err := sc.GetStudent(classId)
	if err != nil || len(studentList) == 0 {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	// 获取班级信息
	cla := model.ClassModel{}
	classInfos := cla.Search([]string{classId})
	if len(classInfos) != 1 {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	res := model.ExperimentClassDetail{
		Class:   classInfos[0].Class,
		Grade:   classInfos[0].Grade,
		Major:   classInfos[0].Major,
		Number:  classInfos[0].Number,
		GroupId: groupId,
		List:    nil,
	}

	// 获取学生详情信息
	list := make([]model.StudentDetail, len(studentList))

	wg := sync.WaitGroup{}
	for i, student := range studentList {
		wg.Add(1)
		go func(i, sid int) {
			defer wg.Done()
			// 获取学生基本信息
			s := model.UserModel{}
			if err := s.Detail(sid); err != nil {
				SendResponse(c, errno.ErrUserNotFound, nil)
				log.Infof("can`t find this student id: %d", student)
				return
			}
			// 获取每题得分信息
			answers := model.AnswerModel{}
			problemScore, err := answers.GetProblemScore(groupId, sid)
			if err != nil {
				SendResponse(c, errno.ErrDatabase, nil)
				return
			}
			// 获取学生总分
			score := 0
			for _, v := range problemScore {
				score += v.Score
			}
			list[i] = model.StudentDetail{
				UserId:   s.UserId,
				Number:   s.Number,
				Name:     s.Username,
				Score:    score,
				Problems: problemScore,
			}
		}(i, student)
	}
	wg.Wait()

	res.List = list
	SendResponse(c, errno.OK, res)
}

func Distributed(c *gin.Context) {
	var request DistributedRequest
	if err := c.Bind(&request); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}
	if len(request.ClassList) == 0 {
		SendResponse(c, errno.ErrParam, nil)
		return
	}
	// class - experiment 中添加数据

	/* student - experiment 中添加数据
		 1. 找出对应 classId 的所有学生 id， student - class
	   2. 更新目标表

		 channel[string](size) 通信机制
	   1. 插入一个 class - experiment, 往 channel 中写入数据
	   2. 等待从 channel 中获取 classId，完成 student - experiment 的更新
	*/

	teacher := service.New()
	err := teacher.Distributed(request.GroupId, request.ClassList)
	if err != nil {
		SendResponse(c, err, nil)
	}
	SendResponse(c, errno.OK, nil)
}
