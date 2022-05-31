package dao

import (
	"gin-practice/src/config"
	"gin-practice/src/entity"
)

var db = config.GolbalConfig.DB

func AddUser(user *entity.User) error {
	return db.Create(user).Error
}
