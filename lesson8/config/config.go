package config

import (
	"fmt"
	"lesson8/component/appctx"
	"log"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func LoadConfig() (*appctx.AppCongfig, error) {
	env, err := godotenv.Read()
	if err != nil {
		return nil, fmt.Errorf("Cen not load environment file %v", err)
	}
	return &appctx.AppCongfig{
		Host:     env["HOST"],
		Password: env["PASSWORD"],
		User:     env["USER"],
		Dbname:   env["DBNAME"],
		Pord:     env["PORT"],
	}, nil
}

func ConnectDatabaseInBoundedTime(cfg *appctx.AppCongfig) (*gorm.DB, error) {
	const timeRetry = 20 * time.Second
	//define function to connect database
	var connectDatabase = func(cfg *appctx.AppCongfig) (*gorm.DB, error) {
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
