package model

import "gopkg.in/go-playground/validator.v9"

type ExperimentModel struct {
	GroupId   int    `json:"group_id" gorm:"column:groupId;primary_key;AUTO_INCREMENT"`
	GroupName string `json:"group_name" gorm:"column:groupName; not null" validate:"gt=0"`
	Poster    string `json:"poster" gorm:"column:poster; not null" validate:"gt=0"`
	Problems  []byte `json:"problems" gorm:"column:problems; not null"`
}

func (e *ExperimentModel) TableName() string {
	return "experiment"
}

func (e *ExperimentModel) Create() error {
	return DB.Self.Create(&e).Error
}

func (e *ExperimentModel) Search(str []string, experimentId int) error {
	return DB.Self.Select(str).Where("groupId = ?", experimentId).Find(&e).Error
}

func (e *ExperimentModel) GetDetail(experimentId int) error {
	return DB.Self.Where("groupId = ?", experimentId).Find(&e).Error
}

func (e *ExperimentModel) Validate() error {
	validate := validator.New()
	return validate.Struct(e)
}
