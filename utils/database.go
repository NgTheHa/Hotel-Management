package utils

import (
	"fmt"
	"hotel-management/models"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbHost, dbUser, dbPassword, dbName, dbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Migrate models
	DB = db
	err = DB.AutoMigrate(
		//&models.Room{},
		//&models.RoomKey{},
		//&models.RoomHousekeeping{},
		&models.Booking{},
		&models.Employee{},
		&models.Service{},
		&models.Notification{},
		&models.Payment{},
		&models.Customer{},
		&models.Hotel{},
		&models.Facility{},
		&models.Review{},
	)
	if err != nil {
		log.Fatalf("Failed to auto migrate models: %v", err)
	}

	fmt.Println("Database connected!")
}
