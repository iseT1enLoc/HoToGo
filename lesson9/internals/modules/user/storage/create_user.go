package userstorage

import (
	usermodel "lesson9/internals/modules/user/models"

	"gorm.io/gorm"
)

func AddUser(db *gorm.DB, user *usermodel.User) error {
	if err := db.Create(user).Error; err != nil {
		return err
	}
	return nil
}
