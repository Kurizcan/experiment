package model

import "experiment/util"

type ClassModel struct {
	ClassId string `json:"class_id" gorm:"column:classId;primary_key;"`
	Grade   int    `json:"grade" gorm:"column:grade;"`
	Class   int    `json:"class" gorm:"column:class;"`
	College string `json:"college" gorm:"column:college;"`
	Major   string `json:"major" gorm:"column:major;"`
	Number  string `json:"number" gorm:"column:number;"`
}

func (c *ClassModel) TableName() string {
	return "class"
}

func (c *ClassModel) Create() error {
	c.ClassId, _ = util.GenShortId()
	return DB.Self.Create(&c).Error
}

func (c *ClassModel) Search(str []string) (res []ClassModel) {
	DB.Self.Where("classId in (?)", str).Find(&res)
	return
}
