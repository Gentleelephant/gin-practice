package cache

import (
	"encoding/json"
	"gin-practice/src/common"
	"github.com/go-redis/redis"
	"time"
)

type RedisWrapper struct {
	Client *redis.Client
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

func (w *RedisWrapper) SetEMailCaptcha(email string, code string) error {
	return w.Client.Set(common.EMAIL_CAPTCHA+common.SEPARATOR+email, code, time.Minute*5).Err()
}

func (w *RedisWrapper) GetEMailCaptcha(email string) (string, error) {
	value := w.Client.Get(common.EMAIL_CAPTCHA + common.SEPARATOR + email)
	return value.Result()
}
