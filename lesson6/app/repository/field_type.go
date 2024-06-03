package repository

import (
	"fmt"
	"log"

	"example.com/database"
	models "model.go"
)

func GetAllFieldType() ([]models.Field_Type, error) {
	db := database.ConnectToDB()

	//close db connection
	defer db.Close()
	var field_type_list []models.Field_Type
	// execute the sql statement
	rows, err := db.Query(`SELECT "field_type_id","field_type_name","price_per_hour" FROM "field_type"`)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	// close the statement
	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var fieldtype models.Field_Type

		// unmarshal the row object to topic
		err = rows.Scan(&fieldtype.FieldTypeId, &fieldtype.FieldTypeName, &fieldtype.PricePerHour)
		fmt.Println(fieldtype)
		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		// append the topic in the topics slice
		field_type_list = append(field_type_list, fieldtype)

	}
	defer db.Close()
	// return empty topic on error
	return field_type_list, err
}
