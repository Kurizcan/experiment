package class

import (
	. "experiment/handler"
	"experiment/model"
	"experiment/pkg/errno"
	"github.com/gin-gonic/gin"
	"strconv"
)

type TeacherClassResponse struct {
	TeacherId int                `json:"teacher_id"`
	List      []model.ClassModel `json:"list"`
}

type TeacherClassDetail struct {
	ClassId string                       `json:"class_id"`
	Grade   int                          `json:"grade"`
	Class   int                          `json:"class"`
	Major   string                       `json:"major"`
	Number  string                       `json:"number"`
	List    []model.ClassExperimentModel `json:"list"`
}

func GetClassByTid(c *gin.Context) {
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

	t := model.TeacherClassModel{}
	classIds := t.GetClassByTeacher(id)
	if len(classIds) == 0 {
		SendResponse(c, errno.ErrClassFound, nil)
		return
	}

	cla := model.ClassModel{}
	list := cla.Search(classIds)
	if len(list) == 0 {
		SendResponse(c, errno.ErrClassFound, nil)
		return
	}

	SendResponse(c, errno.OK, TeacherClassResponse{
		TeacherId: id,
		List:      list,
	})
}

func GetClassDetail(c *gin.Context) {
	classId := c.Param("id")
	if len(classId) == 0 {
		SendResponse(c, errno.ErrParam, nil)
		return
	}
	// 获取 class 信息
	cla := model.ClassModel{}
	if err := cla.Detail(classId); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}
	// 获取 experiment 信息
	exs := model.ClassExperimentModel{}
	list, err := exs.SearchByClassId(classId)
	if err != nil || len(list) == 0 {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	SendResponse(c, errno.OK, TeacherClassDetail{
		ClassId: cla.ClassId,
		Grade:   cla.Grade,
		Class:   cla.Class,
		Major:   cla.Major,
		Number:  cla.Number,
		List:    list,
	})
}
