package db

import (
	"fmt"
	"lesson9/internals/common/appconfig"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func LoadConfig() (*appconfig.AppConfig, error) {
	err := godotenv.Load("../.env")

	if err != nil {
		return &appconfig.AppConfig{}, err
	}
	return &appconfig.AppConfig{
		PORT:     os.Getenv("PORT"),
		HOSTNAME: os.Getenv("HOSTNAME"),
		USERNAME: os.Getenv("USERNAME"),
		PASSWORD: os.Getenv("PASSWORD"),
		DBNAME:   os.Getenv("DBNAME"),
		SECRET:   os.Getenv("SECRET"),
	}, nil
}

func GetConnectionToDatabaseInboundedTime(cfg *appconfig.AppConfig) (*gorm.DB, error) {
	timetry := 5 * time.Second
	//define function to connect database
	var connectDatabase = func(cfg *appconfig.AppConfig) (*gorm.DB, error) {
		dbsourcename := fmt.Sprintf("host=%v user=postgres password=%v dbname=%v port=%v sslmode=disable", cfg.HOSTNAME, cfg.PASSWORD, cfg.DBNAME, cfg.PORT)
		db, err := gorm.Open(postgres.Open(dbsourcename), &gorm.Config{})

		if err != nil {
			return nil, fmt.Errorf("Can not connect to database. Error happenned %v", err)
		}
		return db, nil
	}

	var db *gorm.DB
	var err error

	deadline := time.Now().Add(timetry)
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
