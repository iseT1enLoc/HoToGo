package usermodel

type ReqCreateUser struct {
	Name     string `json:"name" gorm:"column:name"`
	Email    string `json:"email" gorm:"column:email unique"`
	Password string `json:"pass_word" gorm:"column:pass_word"`
}
