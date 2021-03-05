package utils

import "golang.org/x/crypto/bcrypt"

// HashPassword hash pass by bcrypt
func HashPassword(pass string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), 12)
	return string(bytes), err
}
