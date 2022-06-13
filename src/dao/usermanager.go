package dao

import (
	"gin-practice/src/config"
	"gin-practice/src/entity"
)

// TODO
func CreateUser(user *entity.User) error {
	return config.DB.Create(user).Error
}

// TODO
func UpdateUser(user *entity.User) error {
	return nil
}

// TODO
func DeleteUser(user *entity.User) error {
	return nil
}

// TODO
func GetUser(user *entity.User) error {
	return nil
}

// TODO
func CheckUserExist(username string) bool {
	user := &entity.User{}
	user.Username = username
	if checkUser(user) == 0 {
		return false
	}
	return true
}

func checkUser(user *entity.User) int64 {
	first := config.DB.Where("username = ?", user.Username).First(user)
	return first.RowsAffected
}
