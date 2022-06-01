package dao

import (
	"gin-practice/src/config"
	"gin-practice/src/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type MySQL struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

func InitDB() *gorm.DB {

	host := config.GlobalConfig.Mysql.Host
	port := config.GlobalConfig.Mysql.Port
	user := config.GlobalConfig.Mysql.User
	password := config.GlobalConfig.Mysql.Password
	database := config.GlobalConfig.Mysql.Database

	dsn := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + database + "?charset=utf8&parseTime=True&loc=Local"
	open, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("init db error:", err)
	}
	err = open.AutoMigrate(&entity.User{})
	if err != nil {
		log.Println("create table + "+"user"+"error:", err)
	}
	return open
}
