package database

import (
	"log"
	"mentor/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	var err error
	dsn := "root:root@tcp(localhost:3306)/mentor_application?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	DB = database
	// Auto-migrate the User model
	DB.AutoMigrate(&models.User{})
}
