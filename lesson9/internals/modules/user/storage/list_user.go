package userstorage

import (
	usermodel "lesson9/internals/modules/user/models"

	"gorm.io/gorm"
)

func ListUser(db *gorm.DB) ([]usermodel.User, error) {
	//declare a variable
	var users []usermodel.User

	if err := db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
