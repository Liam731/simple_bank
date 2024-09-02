package config

import (
	"SimpleBank/helper"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "postgres"
	port     = 5432
	user     = "root"
	password = "secret"
	dbname   = "root"
)

func DatabaseConnection() *gorm.DB {
	sqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := gorm.Open(postgres.Open(sqlInfo), &gorm.Config{})
	helper.ErrorPanic(err)
	return db
}

func CloseDatabaseConnection(db *gorm.DB) {
	sqlDB, err := db.DB()
	helper.ErrorPanic(err)
	sqlDB.Close()
	err = sqlDB.Close()
	if err != nil {
		log.Fatalf("Failed to close database connection: %v", err)
	} else {
		log.Println("Database connection successfully closed.")
	}
}
