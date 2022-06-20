package global

import (
	"gin-practice/src/config"
	"gin-practice/src/dao"
	"gorm.io/gorm"
)

var (
	DB           *gorm.DB
	PolicyDao    *dao.CasbinRuleManager
	UserDao      *dao.UserManager
	GlobalConfig = &config.Config{}
	ConfigPath   = "C:\\work\\code\\goPro\\gin-practice\\src\\config.yaml"
)
