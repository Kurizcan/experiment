package model

type StudentInfo struct {
	UserId   int    `json:"userId"`
	Username string `json:"username"`
	College  string `json:"college"`
	Grade    int    `json:"password"`
	Major    string `json:"major"`
	Class    int    `json:"class"`
}

type TeacherInfo struct {
	UserId   int    `json:"userId"`
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
