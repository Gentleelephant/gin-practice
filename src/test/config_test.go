package test

import (
	"fmt"
	"gin-practice/src/config"
	"gin-practice/src/dao"
	"testing"
)

func TestLoadConfig(t *testing.T) {

	loadConfig, err := config.LoadConfig("C:\\work\\code\\goPro\\gin-practice\\src\\config\\config.yaml")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Print(*loadConfig)

}

func TestInitDB(t *testing.T) {

	config.InitConfig()
	db := dao.InitDB()
	fmt.Print(db)

}
