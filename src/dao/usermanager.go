package dao

import (
	"gin-practice/src/common"
	"gin-practice/src/config"
	"gin-practice/src/entity"
	"gorm.io/gorm"
)

var (
	UserDao *UserManager
)

func init() {
	UserDao = &UserManager{
		db: config.DB,
	}
}

type UserManager struct {
	db *gorm.DB
}

func (m *UserManager) GetUserByName(username string) (*entity.User, error) {
	user := &entity.User{}
	user.Username = username
	if m.checkUser(user) == 0 {
		return nil, common.UserNotFoundError
	}
	return user, nil
}

func (m *UserManager) CreateUser(user *entity.User) error {
	checkUser := m.checkUser(user)
	if checkUser != 0 {
		return common.UsernameAlreadyExistError
	}
	return config.DB.Create(user).Error
}

func (m *UserManager) checkUser(user *entity.User) int64 {
	first := config.DB.Where("username = ?", user.Username).First(user)
	return first.RowsAffected
}
