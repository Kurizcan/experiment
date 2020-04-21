package model

type StudentInfo struct {
	UserId   int    `json:"user_id"`
	Username string `json:"username"`
	College  string `json:"college"`
	Grade    int    `json:"password"`
	Major    string `json:"major"`
	Class    int    `json:"class"`
}

type TeacherInfo struct {
	UserId   int    `json:"user_id"`
	Username string `json:"username"`
	College  string `json:"college"`
}

type Output struct {
	Headers []string `json:"headers"`
	Rows    []data   `json:"rows"`
}

type Example struct {
	Headers TableField  `json:"headers"`
	Rows    []tableData `json:"rows"`
}

type TableField struct {
	Table []string `json:"table"`
	Field []field  `json:"field"`
}

// 一行数据
type data []string

// 字段
type field []string

// 一张数据库中的数据
type tableData []data

type ExperimentClassDetail struct {
	Class   int             `json:"class_id"`
	Grade   int             `json:"grade"`
	Major   string          `json:"major"`
	Number  string          `json:"number"`
	GroupId int             `json:"group_id"`
	List    []StudentDetail `json:"list"`
}

type StudentDetail struct {
	UserId   int            `json:"user_id"`
	Number   string         `json:"number"`
	Name     string         `json:"name"`
	Score    int            `json:"score"`
	Problems []ProblemScore `json:"problems"`
}

type ProblemScore struct {
	ProblemId int `json:"problem_id" gorm:"column:problemId"`
	Score     int `json:"score" gorm:"column:score"`
	Status    int `json:"status" gorm:"column:status"`
}

type StudentExperiment struct {
	GroupId   int    `json:"group_id"`
	GroupName string `json:"group_name"`
	Status    int    `json:"status"`
	Score     int    `json:"score"`
}

type ExperimentResult struct {
	Score int                 `json:"score"`
	List  []ProblemResultList `json:"list"`
}

type ProblemResultList struct {
	ProblemId int    `json:"problem_id"`
	Score     int    `json:"score"`
	Title     string `json:"title"`
}
