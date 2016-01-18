package services

import (
	"PJApp/app/models"

	_ "github.com/go-sql-driver/mysql"
	"PJApp/app/constant"
)

// Select
func GetUser(username string) *models.User {

	users := []models.User{}
	DB.Debug().Where("Username = ? and Status = ?", username, constant.STATUS_ARI).Find(&users)

	if len(users) != 1 {
		return nil
	}

	return &users[0]
}

// Insert
func InsertUser(user models.User) error {

	user.Status = constant.STATUS_ARI
	return 	DB.Debug().Create(&user).Error
}