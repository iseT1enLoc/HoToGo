package userstorage

import (
	usermodel "lesson9/internals/modules/user/models"

	"gorm.io/gorm"
)

func DeleteUser(db *gorm.DB, id string) (int, error) {
	var user usermodel.User
	if err := db.Delete(&user, id).Error; err != nil {
		return 0, err
	}
	return int(db.RowsAffected), nil
}
