package models

import "time"

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

func (User) TableName() string {
	return "users"
}
