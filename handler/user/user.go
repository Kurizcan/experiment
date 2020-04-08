package user

type LoginRequest struct {
	Number   string `json:"number"`
	Password string `json:"password"`
}

type LoginResponse struct {
	UserId   int    `json:"user_id"`
	UserName string `json:"userName"`
	Token    string `json:"token"`
}
