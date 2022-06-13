package cache

import (
	"encoding/json"
	"fmt"
	"gin-practice/src/common"
	"gin-practice/src/config"
	"github.com/go-redis/redis"
	"time"
)

var (
	Rdb *redis.Client
)

type RedisWrapper struct {
	Client *redis.Client
}

func InitRedis() {
	redisConfig := config.GlobalConfig.Redis
	// 初始化redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", redisConfig.RedisHost, redisConfig.RedisPort),
		Password: redisConfig.RedisPassword, // no password set
		DB:       0,                         // use default DB
	})
	Rdb = rdb
}

func (w *RedisWrapper) SetSession(userId string, value string, expiration time.Duration) error {
	return w.Client.Set(common.SESSION+common.SEPARATOR+userId, value, expiration).Err()
}

func (w *RedisWrapper) GetSession(userId string) (string, error) {
	value := w.Client.Get(common.SESSION + common.SEPARATOR + userId)
	return value.Result()
}

func (w *RedisWrapper) SetUserInfo(sessionId string, value interface{}, expiration time.Duration) error {
	marshal, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return w.Client.Set(common.SESSION+common.SEPARATOR+sessionId, marshal, expiration).Err()
}

func (w *RedisWrapper) GetUserInfo(sessionId string) (interface{}, error) {
	value := w.Client.Get(common.SESSION + common.SEPARATOR + sessionId)
	return value.Result()
}
