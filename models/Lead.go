package models

import "time"

type Lead struct {
	ID        uint      `gorm:"primaryKey;column:id" json:"id"`
	Name      string    `gorm:"type:varchar(100);column:name" json:"name"`
	Email     string    `gorm:"uniqueIndex;type:varchar(100);column:email" json:"email"`
	Mobile    string    `gorm:"type:varchar(50);column:mobile" json:"mobile"`
	Location  string    `gorm:"type:varchar(100);column:location" json:"location"`
	Message   string    `gorm:"column:message" json:"message"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}
