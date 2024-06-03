package userstorage

import (
	usermodel "lesson8/modules/users/model"

	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB, data *usermodel.User) error {
	if err := db.Create(data).Error; err != nil {
		return err
	}

	return nil
}
