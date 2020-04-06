package user

import (
	. "experiment/handler"
	"experiment/model"
	"experiment/pkg/auth"
	"experiment/pkg/constvar"
	"experiment/pkg/errno"
	"experiment/pkg/token"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
)

// Create creates a new user account.
func Create(c *gin.Context) {
	var u model.UserModel
	if err := c.Bind(&u); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	// 验证数据
	if err := u.Validate(); err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}

	// 加密密码
	if err := u.Encrypt(); err != nil {
		SendResponse(c, errno.ErrEncrypt, nil)
		return
	}
	// 插入数据库
	if err := u.Create(); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	SendResponse(c, errno.OK, nil)
}

func Login(c *gin.Context) {
	var u model.UserModel
	if err := c.Bind(&u); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	d, err := model.GetUser(u.Number)
	if err != nil {
		SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}
	log.Infof("user %v", d)

	if err := auth.Compare(d.Password, u.Password); err != nil {
		SendResponse(c, errno.ErrPasswordIncorrect, nil)
		return
	}

	// 签发 token
	t, err := token.Sign(c, token.Context{Number: d.Number, Username: d.Username, Type: float64(d.Type)}, "")
	if err != nil {
		SendResponse(c, errno.ErrToken, nil)
		return
	}

	SendResponse(c, nil, LoginResponse{
		UserId:   d.UserId,
		UserName: d.Username,
		Token:    t,
	})
}

func Get(c *gin.Context) {
	userNumber, _ := c.Get("number")
	userType, _ := c.Get("type")
	number := c.Param("id")
	// 学生非用户本人，无法查询他人的资料
	if userType.(float64) == float64(constvar.Student) && userNumber.(string) != number {
		SendResponse(c, errno.ErrAuthority, nil)
		log.Infof("user: %s want to search user: %s", userNumber, number)
		return
	}
	user, err := model.GetUser(number)
	if err != nil {
		SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}

	switch user.Type {
	case constvar.Student:
		SendResponse(c, nil, model.StudentInfo{
			UserId:   user.UserId,
			Username: user.Username,
			College:  user.College,
			Grade:    user.Grade,
			Major:    user.Major,
			Class:    user.Class,
		})
	case constvar.Teacher:
		SendResponse(c, nil, model.TeacherInfo{
			UserId:   user.UserId,
			Username: user.Username,
			College:  user.College,
		})
	}
}
