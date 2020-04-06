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

// Token represents a JSON web token.
type Token struct {
	Token string `json:"token"`
}
