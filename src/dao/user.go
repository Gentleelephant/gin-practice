package dao

import (
	"gin-practice/src/common"
	"gin-practice/src/entity"
	"gin-practice/src/utils"
	"gorm.io/gorm"
)

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
	randomString, err := utils.GenerateRandomString(5)
	user.UserId = common.ACCOUNT_PREFIX + randomString
	err = m.db.Transaction(func(tx *gorm.DB) error {
		err := tx.Create(user).Error
		if err != nil {
			return err
		}
		err = tx.Create(&entity.CasbinRule{
			PType: "g",
			V0:    user.Username,
			V1:    user.Role,
		}).Error
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (m *UserManager) checkUser(user *entity.User) int64 {
	first := DB.Where("username = ?", user.Username).First(user)
	return first.RowsAffected
}
