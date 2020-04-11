package model

type StudentExperimentModel struct {
	Id        int `json:"-" gorm:"column:id;primary_key;"`
	GroupId   int `json:"group_id" gorm:"column:groupId;"`
	StudentId int `json:"student_id" gorm:"column:studentId;"`
	Score     int `json:"score" gorm:"column:score"`
}

func (s *StudentExperimentModel) TableName() string {
	return "student_experiment"
}
