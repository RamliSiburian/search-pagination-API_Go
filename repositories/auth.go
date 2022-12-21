package repositories

import (
	"MisterAladin/models"

	"gorm.io/gorm"
)

type users struct {
	db *gorm.DB
}

type AuthRepository interface {
	Login(username string) (models.User, error)
	GetUsers(ID int) (models.User, error)
	CheckEmail(email string) (models.User, error)
}

func RepositoryAuth(db *gorm.DB) *users {
	return &users{db}
}

func (r *users) Login(username string) (models.User, error) {
	var user models.User
	err := r.db.First(&user, "username=?", username).Error
	return user, err
}

func (r *users) GetUsers(ID int) (models.User, error) {
	var user models.User
	err := r.db.First(&user, ID).Error

	return user, err
}

func (r *users) CheckEmail(email string) (models.User, error) {
	var user models.User
	err := r.db.First(&user, "email=?", email).Error

	return user, err
}
