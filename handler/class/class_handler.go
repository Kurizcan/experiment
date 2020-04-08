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
