package test

import (
	"fmt"
	"gin-practice/src/config"
	"gin-practice/src/dao"
	"gin-practice/src/entity"
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
	dao.InitDB()
	db := config.GolbalConfig.DB
	user := entity.User{}
	db.Where(&entity.User{Username: "test3", Password: "test"}, "username", "password").Find(&user)
	fmt.Print(user)

}
