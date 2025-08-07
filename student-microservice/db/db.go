package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/gajare/swager_golang/models"
)

var DB *gorm.DB

func InitDB() {
	dsn := "host=localhost user=postgres password=Amol dbname=books port=5432 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto migrate Student model
	DB.AutoMigrate(&models.Student{})
}
