package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/darksuei/chat-kit/config"
)

var DB *gorm.DB

func Connect() {
	DB_HOST, _ := config.ReadEnv("DB_HOST")
	DB_PORT, _ := config.ReadEnv("DB_PORT")
	DB_USERNAME, _ := config.ReadEnv("DB_USERNAME")
	DB_PASSWORD, _ := config.ReadEnv("DB_PASSWORD")
	DB_NAME, _ := config.ReadEnv("DB_NAME")
	DB_SSL_MODE, _ := config.ReadEnv("DB_SSL_MODE")
	
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", DB_HOST, DB_PORT, DB_USERNAME, DB_PASSWORD, DB_NAME, DB_SSL_MODE)
	
	var err error

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to connect to the database: %s", err)
	}

	log.Printf("Database connection successful..")
}
  