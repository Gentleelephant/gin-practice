package dao

import (
	"gin-practice/src/config"
	"gin-practice/src/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
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

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             time.Second,   // 慢 SQL 阈值
			LogLevel:                  logger.Silent, // 日志级别
			IgnoreRecordNotFoundError: true,          // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  true,          // 禁用彩色打印
		},
	)

	dsn := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + database + "?charset=utf8&parseTime=True&loc=Local"
	open, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Println("init db error:", err)
	}
	err = open.AutoMigrate(&entity.User{})
	if err != nil {
		log.Println("create table + "+"user"+"error:", err)
	}
	err = open.AutoMigrate(&entity.CasbinRule{})
	if err != nil {
		log.Println("create table + "+"casebin_rule"+"error:", err)
	}
	return open
}
