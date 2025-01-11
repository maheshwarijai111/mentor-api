package models

import (
	"mentor/database"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID           uint      `gorm:"primaryKey;column:id" json:"id"`
	FirstName    string    `gorm:"type:varchar(100);column:first_name" json:"first_name"`
	LastName     string    `gorm:"type:varchar(100);column:last_name" json:"last_name"`
	Email        string    `gorm:"uniqueIndex;type:varchar(100);column:email" json:"email"`
	PasswordHash string    `gorm:"type:varchar(255);column:password_hash" json:"password_hash"`
	Role         string    `gorm:"type:varchar(50);column:role" json:"role"`
	Location     string    `gorm:"type:varchar(100);column:location" json:"location"`
	CreatedAt    time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at" json:"updated_at"`
}

type UserList struct {
	ID        uint   `gorm:"column:user_id" json:"id"`
	FirstName string `gorm:"type:varchar(100);column:first_name" json:"first_name"`
	LastName  string `gorm:"type:varchar(100);column:last_name" json:"last_name"`
	Email     string `gorm:"uniqueIndex;type:varchar(100);column:email" json:"email"`
	Role      string `gorm:"type:varchar(50);column:role" json:"role"`
	Location  string `gorm:"type:varchar(100);column:location" json:"location"`
}

func VerifyUserInDB(username, password string) bool {
	var user User
	result := database.DB.Where("email = ?", username).Find(&user)
	if password != "" {
		err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
		if err != nil {
			return false
		}
	}
	if result.Error == nil && result.RowsAffected > 0 {
		return true
	} else {
		return false
	}
}

func (User) TableName() string {
	return "users"
}
