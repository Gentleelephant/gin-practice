package dao

import (
	"gin-practice/src/config"
	"gin-practice/src/entity"
)

func AddUser(user *entity.User) error {
	return config.DB.Create(user).Error
}

func CheckUser(user *entity.User) int64 {
	first := config.DB.Where("username = ?", user.Username).First(user)
	return first.RowsAffected
}
