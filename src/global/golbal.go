package global

import (
	"gin-practice/src/dao"
	"gorm.io/gorm"
)

var (
	DB        *gorm.DB
	PolicyDao *dao.CasbinRuleManager
	UserDao   *dao.UserManager
)
