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
type Field_Type struct {
	FieldTypeId   string `json:"field_type_id"`
	FieldTypeName string `json:"field_type_name"`
	PricePerHour  string `json:"price_per_hour"`
}
