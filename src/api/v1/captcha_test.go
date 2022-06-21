package v1

import (
	"fmt"
	"gin-practice/src/utils"
	"testing"
)

func TestGetCaptcha(t *testing.T) {
	_, err := GetCaptchaFile()
	if err != nil {
		t.Error(err)
	}
}

func TestName(t *testing.T) {
	num := utils.RandNum(7)
	fmt.Println(num)
}
