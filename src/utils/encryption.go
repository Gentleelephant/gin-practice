package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func Encryption(str string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func Verify(str, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(str), []byte(hash))
	if err != nil {
		return false
	}
	return true
}
