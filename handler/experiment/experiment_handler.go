package experiment

import (
	"encoding/json"
	. "experiment/handler"
	"experiment/model"
	"experiment/pkg/errno"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
)

type createRequest struct {
	Problems  []string `json:"problems" binding:"required"`
	GroupName string   `json:"group_name" binding:"required"`
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
