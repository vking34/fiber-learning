package user_service

import (
	"bytes"
	"encoding/gob"
	"errors"
	"strconv"
	"time"

	"github.com/vking34/fiber-messenger/db"
	"github.com/vking34/fiber-messenger/models"
	"gorm.io/gorm"
)

//
func findPgUserByID(id int) (*models.User, error) {
	var user models.User
	if err := db.Pg.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	userID := strconv.Itoa(id)
	isExist := db.Redis.Exists(db.Ctx, userID).Val()
	if isExist == 0 {
		db.SetCache(userID, &user, 10*time.Second)
	}

	return &user, nil
}

// FindUserByID returns user
func FindUserByID(id int) (*models.User, error) {
	cmd := db.Redis.Get(db.Ctx, strconv.Itoa(id))
	userInBinary, err := cmd.Bytes()
	if err != nil {
		return findPgUserByID(id)
	}

	userByteReader := bytes.NewReader(userInBinary)
	var user *models.User
	if err = gob.NewDecoder(userByteReader).Decode(&user); err != nil {
		return findPgUserByID(id)
	}

	return user, nil
}

func FindUsersByUsernameOrEmail(username string, email string) (*gorm.DB, []models.User) {
	var users []models.User
	result := db.Pg.Where("username =?", username).Or("email =?", email).Find(&users)

	return result, users
}

// FindUser find user by username
func FindUser(username string) (*models.User, error) {
	var user models.User
	if err := db.Pg.Where(&models.User{Username: username}).Find(&user).Error; err != nil {
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

	if err := db.Pg.Where(&models.User{Email: email}).Find(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

// CreateUser create user
func CreateUser(user *models.User) error {
	return db.Pg.Create(&user).Error
}
