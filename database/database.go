package database

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	var err error
	dsn := "admin:Demo#$12345@tcp(api-db.czguk2c0a24d.eu-north-1.rds.amazonaws.com:3306)/mentor?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	DB = database
	// Auto-migrate the User model
	//DB.AutoMigrate(&models.User{})
}
