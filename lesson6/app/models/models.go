package models

type Customer struct {
	CustomerId   string `json:"customer_id"`
	CustomerName string `json:"customer_name"`
	Gmail        string `json:"gmail"`
	PhoneNumber  string `json:"phone_number"`
	DateOfBirth  string `json:"date_of_birth"`
	PassWord     string `json:"pass_word"`
	CustomerType string `json:"customer_type"`
}
