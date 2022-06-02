package dao

import (
	"gin-practice/src/config"
	"gin-practice/src/entity"
	"gorm.io/gorm"
)

func AddUser(user *entity.User) error {
	return config.DB.Create(user).Error
}
func AddUserDn(dn *entity.UserDn) error {
	return config.DB.Create(dn).Error
}

func AddUserAndUserDn(user *entity.User, dn *entity.UserDn) error {
	return config.DB.Transaction(func(tx *gorm.DB) error {
		err := tx.Create(user).Error
		if err != nil {
			return err
		}
		return tx.Create(dn).Error
	})
}

func CheckUser(user *entity.User) int64 {
	first := config.DB.Where("username = ?", user.Username).First(user)
	return first.RowsAffected
}
