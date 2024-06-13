package database

import (
	"amish-directory/database/models"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := os.Getenv("DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to the Database!")
	}
	db.AutoMigrate(&models.Amish{})
	db.AutoMigrate(&models.Drivers{})
	db.AutoMigrate(&models.Events{})
	db.AutoMigrate(&models.Roles{})
	db.AutoMigrate(&models.Users{})
	db.AutoMigrate(&models.Vehicles{})

	DB = db
}