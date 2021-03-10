package models

import (
	"time"

	"gorm.io/gorm"
)

// User model
type User struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Username  string         `gorm:"uniqueIndex;not null" json:"username"`
	Password  string         `gorm:"not null" json:"-"`
	Email     string         `gorm:"not null" json:"email"`
	Name      string         `gorm:"not null" json:"name"`
	Picture   string         `gorm:"column=picture" json:"picture"`
}
