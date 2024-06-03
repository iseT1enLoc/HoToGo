package usermodel

type User struct {
	ID   int    `json:"id" gorm:"column:id"`
	Name string `json:"name" gorm:"column:name"`
}

func (User) TableName() string {
	return "users"
}
