package cache

import (
	"fmt"
	"gin-practice/src/entity"
	"github.com/go-redis/redis"
	"testing"
	"time"
)

func TestCache(t *testing.T) {

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	wrapper := RedisWrapper{
		Client: rdb,
	}

	user := entity.User{
		ID:        0,
		Username:  "",
		Password:  "",
		Email:     "",
		Phone:     "",
		CreatedAt: 0,
		UpdatedAt: 0,
		DeletedAt: 0,
	}

	err := wrapper.SetUserInfo("user123456", &user, time.Second*360)

	if err != nil {
		t.Error(err)
	}

	get, _ := wrapper.GetUserInfo("user123456")
	fmt.Println(get)

	err = wrapper.SetSession("usertest", "testetestets", time.Second*360)
	if err != nil {
		t.Error(err)
	}
	session, err := wrapper.GetSession("usertest")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(session)

}
