package dao

import (
	"gin-practice/src/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func InitDB() *gorm.DB {

	host := config.GolbalConfig.Mysql.Host
	port := config.GolbalConfig.Mysql.Port
	user := config.GolbalConfig.Mysql.User
	password := config.GolbalConfig.Mysql.Password
	database := config.GolbalConfig.Mysql.Database

	dsn := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + database + "?charset=utf8&parseTime=True&loc=Local"
	open, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("init db error:", err)
	}
	return open
}
