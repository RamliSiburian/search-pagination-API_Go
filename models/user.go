package models

type User struct {
	ID       int    `json:"id" gorm:"primary_key:auto_increment"`
	Fullname string `json:"fullname" gorm:"type: varchar(255)"`
	Username string `json:"username" gorm:"type: varchar(255)"`
	Password string `json:"password" gorm:"type: varchar(255)"`
}
type UserResponse struct {
	ID       int    `json:"id"`
	Fullname string `json:"fullname"`
}

func (UserResponse) TableName() string {
	return "users"
}
