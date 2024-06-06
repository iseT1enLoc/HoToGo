package usermodel

type User struct {
	ID       int    `json:"id" gorm:"column:id"`
	UserName string `json:"user_name" gorm:"column:name"`
	Email    string `json:"email" gorm:"column:"email,unique"`
	Password string `json:"password" gorm:"column:pass_word"`
}
