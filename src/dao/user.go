package dao

import (
	"gin-practice/src/common"
	"gin-practice/src/entity"
	"gin-practice/src/utils"
	"gorm.io/gorm"
)

type UserManager struct {
	DB *gorm.DB
}

func (m *UserManager) GetUserByName(username string) (*entity.User, error) {
	user := &entity.User{}
	user.Username = username
	return user, nil
}

func (m *UserManager) CreateUser(user *entity.User) error {
	randomString, err := utils.GenerateRandomString(5)
	user.UserId = common.ACCOUNT_PREFIX + randomString
	err = m.DB.Transaction(func(tx *gorm.DB) error {
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
