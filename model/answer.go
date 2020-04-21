package model

type AnswerModel struct {
	Id         int    `json:"-" gorm:"column:id;primary_key;"`
	GroupId    int    `json:"group_id" gorm:"column:groupId;"`
	StudentId  int    `json:"student_id" gorm:"column:studentId;"`
	ProblemId  int    `json:"problem_id" gorm:"column:problemId"`
	Status     int    `json:"status" gorm:"column:status"`
	Score      int    `json:"score" gorm:"column:score"`
	Submit     string `json:"submit" gorm:"column:submit"`
	Error      string `json:"error" gorm:"column:error"`
	Correct    bool   `json:"correct" gorm:"column:correct"`
	UpdateTime int64  `json:"update_time" gorm:"column:updateTime"`
}

func (a *AnswerModel) TableName() string {
	return "answer"
}

func (a *AnswerModel) GetProblemScore(groupId, studentId int) ([]ProblemScore, error) {
	res := make([]ProblemScore, 0)
	db := DB.Self.Table(a.TableName()).Where("groupId = ? and studentId = ?", groupId, studentId).Scan(&res)
	return res, db.Error
}

func (a *AnswerModel) Create() error {
	return DB.Self.Create(&a).Error
}

func (a *AnswerModel) Detail(groupId, studentId, problemId int) error {
	return DB.Self.Where("studentId = ? and groupId = ? and problemId = ?", studentId, groupId, problemId).Find(&a).Error
}

func (a *AnswerModel) DetailById(id int) error {
	return DB.Self.Where("id = ?", id).Find(&a).Error
}

func (a *AnswerModel) Update(data map[string]interface{}, id int) error {
	return DB.Self.Table(a.TableName()).Where("id = ?", id).Update(data).Find(&a).Error
}
