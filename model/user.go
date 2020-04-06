package model

import (
	"experiment/pkg/auth"
	validator "gopkg.in/go-playground/validator.v9"
)

// 用户信息
type UserModel struct {
	UserId   int    `json:"user_id" gorm:"column:userId;primary_key;AUTO_INCREMENT"`
	Username string `json:"username" gorm:"column:username;not null"`
	Password string `json:"password" gorm:"column:password;not null" binding:"required"`
	Number   string `json:"number" gorm:"column:number;not null" binding:"required"`
	Type     int    `json:"type" gorm:"column:type;not null"`
	Grade    int    `json:"grade" gorm:"column:grade"`
	Class    int    `json:"class" gorm:"column:class"`
	Major    string `json:"major" gorm:"column:major"`
	College  string `json:"college" gorm:"column:college"`
}

func (c *UserModel) TableName() string {
	return "user"
}

// 创建
func (u *UserModel) Create() error {
	return DB.Self.Create(&u).Error
}

// 更新
func (u *UserModel) Update() error {
	return DB.Self.Save(u).Error
}

// 根据学号获取用户信息
func GetUser(number string) (*UserModel, error) {
	u := &UserModel{}
	d := DB.Self.Where("number = ?", number).Find(&u)
	return u, d.Error
}

// 比较密码是否是正确的密码
func (u *UserModel) Compare(pwd string) (err error) {
	err = auth.Compare(u.Password, pwd)
	return
}

// 加密密码
func (u *UserModel) Encrypt() (err error) {
	u.Password, err = auth.Encrypt(u.Password)
	return
}

// 验证字段
func (u *UserModel) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}
