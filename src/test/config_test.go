package test

import (
	"fmt"
	"gin-practice/src/config"
	"testing"
)

func TestLoadConfig(t *testing.T) {

	loadConfig, err := config.LoadConfig("C:\\work\\code\\goPro\\gin-practice\\src\\config\\config.yaml")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Print(*loadConfig)

}
