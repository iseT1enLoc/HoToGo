package usermodel

type ReqUser struct {
	UserName string `json:"user_name" gorm:"column:name"`
	Email    string `json:"email" gorm:"column:"email"`
	Password string `json:"password" gorm:"column:pass_word"`
}
