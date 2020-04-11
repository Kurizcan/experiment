package model

import "fmt"

type StudentClassModel struct {
	Id        int    `json:"-" gorm:"column:id;primary_key;"`
	StudentId int    `json:"student_id" gorm:"column:studentId;"`
	ClassId   string `json:"class_id" gorm:"column:classId"`
}

func (s *StudentClassModel) TableName() string {
	return "student_class"
}

func (s *StudentClassModel) GetStudent(classId string) ([]int, error) {
	org := make([]StudentClassModel, 0)
	res := DB.Self.Where("classId = ?", classId).Find(&org)
	fmt.Println(org)
	var list []int
	if res.Error == nil {
		list = make([]int, len(org))
		for i, v := range org {
			list[i] = v.StudentId
		}
	}
	fmt.Print(list)
	return list, res.Error
}
