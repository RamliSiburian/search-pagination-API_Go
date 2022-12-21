package Databases

import (
	"MisterAladin/models"
	"MisterAladin/pkg/mysql"
	"fmt"
)

func Migration() {
	err := mysql.DB.AutoMigrate(&models.Article{}, &models.User{})

	if err != nil {
		panic("migration filed")
	}

	fmt.Println("migration success")
}
