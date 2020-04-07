package model

import "gopkg.in/go-playground/validator.v9"

type ProblemModel struct {
	ProblemId   int    `json:"problemId" gorm:"column:problem;primary_key;AUTO_INCREMENT"`
	Title       string `json:"title" gorm:"column:title;not null" validate:"gt=0"`
	Description string `json:"description" gorm:"column:description;not null" validate:"gt=0"`
	Example     []byte `json:"example" gorm:"column:example;not null"`
	Data        string `json:"data" gorm:"column:data;not null"`
	Solution    string `json:"solution" gorm:"column:Solution;not null" validate:"gt=0"`
	Output      []byte `json:"output" gorm:"column:output;not null"`
	Poster      string `json:"poster" gorm:"column:poster;not null" validate:"gt=0"`
}

func (p *ProblemModel) TableName() string {
	return "problem"
}

// 创建
func (p *ProblemModel) Create() error {
	return DB.Self.Create(&p).Error
}

func (p *ProblemModel) Update(problemId string, data interface{}) error {
	return DB.Self.Model(&p).Where("problemId = ?", problemId).Update(data).Error
}

// 验证字段
func (p *ProblemModel) Validate() error {
	validate := validator.New()
	return validate.Struct(p)
}
