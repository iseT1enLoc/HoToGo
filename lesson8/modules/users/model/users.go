package usermodel

type User struct {
	ID       int    `json:"id" gorm:"column:id"`
	Name     string `json:"name" gorm:"column:name"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"pass_word" gorm:"column:pass_word"`
}

func (User) TableName() string {
	return "users"
}
