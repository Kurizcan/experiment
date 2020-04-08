package experiment

import (
	"encoding/json"
	. "experiment/handler"
	"experiment/model"
	"experiment/pkg/errno"
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
