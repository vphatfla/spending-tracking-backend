package util

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(passwordStr string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(passwordStr), bcrypt.DefaultCost)
	return string(bytes), err
}
