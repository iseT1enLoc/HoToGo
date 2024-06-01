package repository

import (
	"database/sql"
	"fmt"
	"log"

	"example.com/database"
	"golang.org/x/crypto/bcrypt"
	models "model.go"
)

func GetAllCustomer() ([]models.Customer, error) {
	db := database.ConnectToDB()

	//close db connection
	defer db.Close()
	var cuslist []models.Customer
	// execute the sql statement
	rows, err := db.Query(`SELECT "customer_id","customer_name","gmail","date_of_birth","phone_number","pass_word","customer_type" FROM "customer"`)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	// close the statement
	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var customer models.Customer

		// unmarshal the row object to topic
		err = rows.Scan(&customer.CustomerId, &customer.CustomerName, &customer.Gmail, &customer.DateOfBirth, &customer.PhoneNumber, &customer.PassWord, &customer.CustomerType)
		fmt.Println(customer)
		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		// append the topic in the topics slice
		cuslist = append(cuslist, customer)

	}
	defer db.Close()
	// return empty topic on error
	return cuslist, err
}

func GetCustomerById(cus_id string) (models.Customer, error) {
	// create the postgres db connection
	db := database.ConnectToDB()

	// close the db connection
	defer db.Close()

	// create a word of models.word type
	var customer models.Customer

	// create the select sql query
	sqlStatement := `SELECT "customer_id","customer_name","gmail","date_of_birth","phone_number","pass_word","customer_type" FROM "customer" WHERE "customer_id" = $1 `

	// execute the sql statement
	rows := db.QueryRow(sqlStatement, cus_id)

	// unmarshal the row object to word
	err := rows.Scan(&customer.CustomerId, &customer.CustomerName, &customer.Gmail, &customer.DateOfBirth, &customer.PhoneNumber, &customer.PassWord, &customer.CustomerType)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return customer, nil
	case nil:
		return customer, nil
	default:
		log.Fatalf("Unable to scan the row. %v", err)
	}

	// return empty word on error
	return customer, err
}

func InsertNewCustomer(customer models.Customer) (models.Customer, error) {
	db := database.ConnectToDB()
	defer db.Close()
	sqlStatement := `INSERT INTO "customer"("customer_name","gmail","date_of_birth","phone_number","pass_word","customer_type") VALUES ($1, $2, $3, $4, $5,$6) RETURNING "customer_id"`

	var id string
	password := string(string(customer.PassWord))
	// Hashing the passowrd using Salt with SALT_SECRET env var
	salt := []byte(password + "dfghjkshfghldflskf")
	hashedPasswordBytes, err := bcrypt.GenerateFromPassword([]byte(salt), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalln(err)
	}
	hashedPassword := string(hashedPasswordBytes)

	err = db.QueryRow(sqlStatement, customer.CustomerName, customer.Gmail, customer.DateOfBirth, customer.PhoneNumber, hashedPassword, customer.CustomerType).Scan(&id)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	customer.CustomerId = id

	fmt.Printf("Inserted a single record %v", id)
	return customer, nil
}

func UpdateCustomer(id string, customer models.Customer) int64 {
	// create the postgres db connection
	db := database.ConnectToDB()

	// close the db connection
	defer db.Close()

	// create the update sql query
	sqlStatement := `UPDATE "customer" SET "customer_name"=$2, "gmail"=$3,"phone_number"=$4,"date_of_birth"=$5,"customer_type"=$6 WHERE "customer_id"=$1`

	// execute the sql statement
	res, err := db.Exec(sqlStatement, id, customer.CustomerName, customer.Gmail, customer.PhoneNumber, customer.DateOfBirth, customer.CustomerType)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	fmt.Printf("Total rows/record affected %v", rowsAffected)

	return rowsAffected
}

func DeleteCustomer(id string) int64 {
	//connect to database
	db := database.ConnectToDB()
	// create the update sql query
	sqlStatement := `DELETE FROM "customer" WHERE "customer_id"=$1`

	// execute the sql statement
	res, err := db.Exec(sqlStatement, id)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	fmt.Printf("Total rows/record affected %v", rowsAffected)

	return rowsAffected
}
