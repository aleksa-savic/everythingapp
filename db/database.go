package database

import (
	"everythingapp/model"

	"gorm.io/gorm"
)

// DB gorm connector
var DB *gorm.DB

func DeleteUser(userID uint) error {
	result := DB.Delete(&model.User{}, userID)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
