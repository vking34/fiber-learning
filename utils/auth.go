package utils

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret string

func init() {
	jwtSecret = os.Getenv("JWT_SECRET")
}

// HashPassword hash pass by bcrypt
func HashPassword(pass string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), 10)
	return string(bytes), err
}

// CheckPasswordHash compare hashed pass vs request pass
func CheckPasswordHash(pass, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
	return err == nil
}

// GenerateJWTToken generates token
func GenerateJWTToken(userID uint, username string, expireHours int) (string, error) {
	jwtGenerator := jwt.New(jwt.SigningMethodHS256)

	claims := jwtGenerator.Claims.(jwt.MapClaims)
	claims["user_id"] = username
	claims["username"] = userID
	claims["exp"] = time.Now().Add(time.Duration(int(time.Hour) * expireHours)).Unix()

	return jwtGenerator.SignedString([]byte(jwtSecret))
}
