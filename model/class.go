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

func (c *ClassModel) Detail(classId string) error {
	return DB.Self.Where("classId = ?", classId).Find(&c).Error
}

type ClassExperimentModel struct {
	Id        int    `json:"-" gorm:"column:id;primary_key;"`
	GroupId   int    `json:"group_id" gorm:"column:groupId;"`
	ClassId   string `json:"class_id" gorm:"column:classId;"`
	GroupName string `json:"group_name" gorm:"column:groupName;"`
}

func (ce *ClassExperimentModel) TableName() string {
	return "class_experiment"
}

func (ce *ClassExperimentModel) SearchByClassId(classId string) ([]ClassExperimentModel, error) {
	res := make([]ClassExperimentModel, 0)
	s := DB.Self.Table(ce.TableName()).Where("classId = ?", classId).Find(&res)
	return res, s.Error
}
