package usermodel

type ReqCreateUser struct {
	Name string `json:"name" gorm:"column:name"`
}
