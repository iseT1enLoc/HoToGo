package userstorage

import (
	usermodel "lesson9/internals/modules/user/models"

	"gorm.io/gorm"
)

func UpdateOneUser(db *gorm.DB, user *usermodel.User) (int, error) {
	if err := db.First(&user).Error; err != nil {
		return 0, err
	}
	db.Save(&user)
	return int(db.RowsAffected), nil
}
