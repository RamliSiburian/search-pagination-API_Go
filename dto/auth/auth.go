package auth

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required" `
}

type LoginResponse struct {
	ID       int    `json:"id"`
	Fullname string `json:"fullname"`
	Token    string `json:"token"`
}
type CheckAuthResponse struct {
	ID       int    `gorm:"type: int" json:"id"`
	Fullname string `json:"fullname"`
}
