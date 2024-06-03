package userstorage

import (
	usermodel "lesson8/modules/users/model"

	"gorm.io/gorm"
)

func ListUser(db *gorm.DB) ([]usermodel.User, error) {
	var users []usermodel.User

	if err := db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}
