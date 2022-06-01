package dao

import (
	"gin-practice/src/config"
	"gin-practice/src/entity"
)

func AddUser(user *entity.User) error {
	return config.DB.Create(user).Error
}
