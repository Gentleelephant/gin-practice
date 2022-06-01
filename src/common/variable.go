package common

import (
	"gin-practice/src/config"
	"gorm.io/gorm"
)

var (
	DB           = &gorm.DB{}
	GlobalConfig = &config.Config{}
	ConfigPath   = "C:\\work\\code\\goPro\\gin-practice\\src\\config\\config.yaml"
)
