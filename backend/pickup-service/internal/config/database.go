package config

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() *gorm.DB {
	host := os.Getenv("DB_HOST")
	if host == "" {
		host = "localhost"
	}
	password := os.Getenv("DB_PASS")
	if password == "" {
		password = "admin123"
	}
	dbname := os.Getenv("DB_NAME")
	if dbname == "" {
		dbname = "pickup_db"
	}

	dsn := fmt.Sprintf(
		"host=%s user=postgres password=%s dbname=%s port=5432 sslmode=disable",
		host, password, dbname,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}

	DB = db
	return db
}
