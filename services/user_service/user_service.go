package user_service

import (
	"errors"

	"github.com/vking34/fiber-messenger/db"
	"github.com/vking34/fiber-messenger/models"
	"gorm.io/gorm"
)

// FindUser find user by username
func FindUser(username string) (*models.User, error) {
	var user models.User
	if err := db.DB.Where(&models.User{Username: username}).Find(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &user, nil
}

// FindUserByEmail find user by email
func FindUserByEmail(email string) (*models.User, error) {
	var user models.User

	if err := db.DB.Where(&models.User{Email: email}).Find(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &user, nil
}

// CreateUser create user
func CreateUser(user *models.User) error {
	return db.DB.Create(&user).Error
}
