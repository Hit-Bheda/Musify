// Package postgres contains all postgres realted utilities
package postgres

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load env!")
	}

	dbs := os.Getenv("DB_URL")
	DB, err := gorm.Open(postgres.Open(dbs), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to databse!")
	}

	return DB
}
