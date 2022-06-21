package global

import (
	"gin-practice/src/cache"
	"gin-practice/src/config"
	"gin-practice/src/dao"
	"github.com/go-redis/redis"
	"github.com/jordan-wright/email"
	"gorm.io/gorm"
)

var (
	DB           *gorm.DB
	PolicyDao    *dao.CasbinRuleManager
	UserDao      *dao.UserManager
	GlobalConfig = &config.Config{}
	ConfigPath   = "C:\\work\\code\\goPro\\gin-practice\\src\\config.yaml"
	Pool         *email.Pool
	Rdb          *redis.Client
	RedisClient  *cache.RedisWrapper
	EmailLists   chan *email.Email
)
