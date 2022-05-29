package dao

import (
	"gin-practice/src/config"
	"gin-practice/src/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func InitDB() {

	host := config.GolbalConfig.Mysql.Host
	port := config.GolbalConfig.Mysql.Port
	user := config.GolbalConfig.Mysql.User
	password := config.GolbalConfig.Mysql.Password
	database := config.GolbalConfig.Mysql.Database

	dsn := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + database + "?charset=utf8&parseTime=True&loc=Local"
	open, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("init db error:", err)
	}
	err = open.AutoMigrate(&entity.User{})
	if err != nil {
		log.Println("create table + "+"user"+"error:", err)
	}
	config.GolbalConfig.DB = open
}
