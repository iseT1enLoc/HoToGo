package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func ConnectToDB() *sql.DB {

	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Access the environment variables
	var url = os.Getenv("postgres_url")
	// open database
	db, err := sql.Open("postgres", url)
	CheckError(err)

	// close database
	//defer db.Close()

	// check db
	err = db.Ping()
	CheckError(err)

	fmt.Println("Connected!")
	return db
}
func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
