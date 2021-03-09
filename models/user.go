package models

import "gorm.io/gorm"

// User model
type User struct {
	gorm.Model
	Username string `gorm:"uniqueIndex;not null" json:"username"`
	Password string `gorm:"not null" json:"password"`
	Email    string `gorm:"not null" json:"email"`
	Name     string `gorm:"not null"`
}
