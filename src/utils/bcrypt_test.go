package utils

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestName(t *testing.T) {

	password := "zhang"

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		t.Error(err)
	}
	//bytes1, err := bcrypt.GenerateFromPassword([]byte(pasword), bcrypt.DefaultCost)

	err = bcrypt.CompareHashAndPassword(bytes, []byte("zhang"))
	if err != nil {
		fmt.Println("error:", err)
	}
	//verify := Verify(string(bytes), "zhang"))
	verify := Verify("$2a$10$Cm8Q838bzoPmKYkM60xivuLqhfauU7uGCh/CcUxqZq.8118eDS.Pm", "test")
	fmt.Println(verify)

}
