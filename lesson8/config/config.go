package config

import (
	"fmt"
	"lesson8/component/appconfig"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func LoadConfig() (*appconfig.AppCongfig, error) {
	fmt.Println("Enter load config")
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	return &appconfig.AppCongfig{
		Host:     os.Getenv("HOST"),
		Password: os.Getenv("PASSWORD"),
		User:     os.Getenv("USER"),
		Dbname:   os.Getenv("DBNAME"),
		Pord:     os.Getenv("PORT"),
	}, nil
}

func ConnectDatabaseInBoundedTime(cfg *appconfig.AppCongfig) (*gorm.DB, error) {
	const timeRetry = 5 * time.Second
	//define function to connect database
	var connectDatabase = func(cfg *appconfig.AppCongfig) (*gorm.DB, error) {
		dbsourcename := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable", cfg.Host, cfg.User, cfg.Password, cfg.Dbname, cfg.Pord)
		db, err := gorm.Open(postgres.Open(dbsourcename), &gorm.Config{})

		if err != nil {
			return nil, fmt.Errorf("Can not connect to database. Error happenned %v", err)
		}
		return db, nil
	}
	var db *gorm.DB
	var err error

	deadline := time.Now().Add(timeRetry)

	for time.Now().Before(deadline) {
		log.Println("Connecting to database...")
		db, err = connectDatabase(cfg)
		if err == nil {
			return db, nil
		}
		time.Sleep(time.Second)
	}

	return nil, fmt.Errorf("failed to connect to database after retrying for 20 seconds: %w", err)

}
