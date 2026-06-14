package config

import (
	"payment-service/internal/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {

	dsn := "host=localhost user=postgres password=admin123 dbname=payment_db port=5432 sslmode=disable"

	db, err := gorm.Open(
		postgres.Open(dsn),
		&gorm.Config{},
	)

	if err != nil {
		panic(err)
	}

	DB = db
	err = DB.AutoMigrate(
	&entity.Payment{},
)

if err != nil {
	panic(err)
}
}