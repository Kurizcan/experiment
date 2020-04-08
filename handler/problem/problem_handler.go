package problem

import (
	"encoding/json"
	. "experiment/handler"
	"experiment/model"
	"experiment/pkg/errno"
	"experiment/util"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
)

func Create(c *gin.Context) {
	var request UploadRequest
	if err := c.Bind(&request); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		log.Error("", err)
		return
	}

	example, err := json.Marshal(request.Example)
	if err != nil {
		SendResponse(c, errno.ErrJsonMarshal, nil)
		return
	}
	output, err := json.Marshal(request.Output)
	if err != nil {
		SendResponse(c, errno.ErrJsonMarshal, nil)
		return
	}

	username, _ := c.Get("username")
	p := model.ProblemModel{
		Title:       request.Title,
		Description: request.Description,
		Example:     example,
		Solution:    request.Solution,
		Output:      output,
		Poster:      username.(string),
	}

	if err := p.Validate(); err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}

	if err := p.Create(); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	SendResponse(c, errno.OK, map[string]int{
		"problemId": p.ProblemId,
	})
}

func UploadData(c *gin.Context) {
	problemId := c.Param("id")
	if len(problemId) == 0 {
		SendResponse(c, errno.ErrFileInit, nil)
		return
	}

	file, _, err := c.Request.FormFile("data")
	if err != nil {
		SendResponse(c, errno.ErrFileInit, nil)
		return
	}

	data, err := util.StoreFile(file)
	if err != nil {
		SendResponse(c, errno.ErrFileInit, nil)
		return
	}

	p := &model.ProblemModel{}
	if err := p.Update(problemId, map[string]interface{}{"data": data}); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	SendResponse(c, errno.OK, nil)
}

func Detail(c *gin.Context) {
	id := c.Param("id")
	if len(id) == 0 {
		SendResponse(c, errno.ErrParam, nil)
		return
	}

	p := model.ProblemModel{}
	if err := p.Detail(id); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	example := model.Example{}
	err := json.Unmarshal(p.Example, &example)
	if err != nil {
		SendResponse(c, errno.ErrJsonUnmarshal, nil)
		return
	}
	output := model.Output{}
	err = json.Unmarshal(p.Output, &output)
	if err != nil {
		SendResponse(c, errno.ErrJsonUnmarshal, nil)
		return
	}

	SendResponse(c, errno.OK, DetailResponse{
		ProblemId:   p.ProblemId,
		Title:       p.Title,
		Description: p.Description,
		Example:     example,
		Output:      output,
		Poster:      p.Poster,
	})

}
