package service

type LoginRequest struct {
	Username string `json:"username" binding:"max=20"`
	Password string `json:"password" binding:"max=20"`
}
