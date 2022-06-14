package utils

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestName(t *testing.T) {

	pasword := "zhang"

	bytes, err := bcrypt.GenerateFromPassword([]byte(pasword), bcrypt.DefaultCost)

	//bytes1, err := bcrypt.GenerateFromPassword([]byte(pasword), bcrypt.DefaultCost)

	err = bcrypt.CompareHashAndPassword(bytes, []byte("zhang"))
	if err != nil {
		fmt.Println("error:", err)
	}

}
