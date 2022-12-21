package models

import "time"

type Article struct {
	ID      int          `json:"id" gorm:"primary_key:auto_increment"`
	Image   string       `json:"image" gorm:"type: varchar(255)"`
	Title   string       `json:"title" gorm:"type: TEXT"`
	Body    string       `json:"body" gorm:"type: TEXT"`
	Created time.Time    `json:"created"`
	UserID  int          `json:"user_id"`
	User    UserResponse `json:"user"`
}
