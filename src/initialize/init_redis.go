package initialize

import (
	"fmt"
	"gin-practice/src/cache"
	"gin-practice/src/global"
	"github.com/go-redis/redis"
)

func InitRedis() {
	redisConfig := global.GlobalConfig.Redis
	// 初始化redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", redisConfig.RedisHost, redisConfig.RedisPort),
		Password: redisConfig.RedisPassword, // no password set
		DB:       0,                         // use default DB
	})
	global.Rdb = rdb
	wrapper := cache.RedisWrapper{
		Client: rdb,
	}
	global.RedisClient = &wrapper
}
