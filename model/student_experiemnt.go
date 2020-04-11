package model

type StudentExperimentModel struct {
	Id        int `json:"-" gorm:"column:id;primary_key;"`
	GroupId   int `json:"group_id" gorm:"column:groupId;"`
	StudentId int `json:"student_id" gorm:"column:studentId;"`
	Score     int `json:"score" gorm:"column:score"`
	Status    int `json:"status" gorm:"column:status"`
}

func (s *StudentExperimentModel) TableName() string {
	return "student_experiment"
}

func (s *StudentExperimentModel) GetGroups(studentId int) ([]StudentExperimentModel, error) {
	res := make([]StudentExperimentModel, 0)
	db := DB.Self.Where("studentId = ?", studentId).Find(&res)
	return res, db.Error
}
