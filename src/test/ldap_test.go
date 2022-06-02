package test

import (
	"gin-practice/src/auth"
	"gin-practice/src/config"
	"gin-practice/src/dao"
	"gin-practice/src/entity"
	"testing"
)

func TestLdap(t *testing.T) {

	err := config.LoadConfig(config.ConfigPath)
	config.DB = dao.InitDB()
	if err != nil {
		return
	}
	provider := auth.GetProvider()
	authentication, dn, _ := provider.Authentication("zhang", "zhang")
	if authentication == false {
		t.Error("authentication failed")
		return
	}
	// 判断表中是否以及包含该用户
	user := &entity.User{
		Username: "lihua",
		Email:    "1132960613@qq.com",
	}

	take := dao.CheckUser(user)
	t.Log(take)
	if take != 1 {
		t.Log("user not exist.......")
		t.Log("insert user......")
		// 插入用户
		err := dao.AddUserAndUserDn(user, &entity.UserDn{
			Username: user.Username,
			Dn:       dn,
		})
		if err != nil {
			return
		}
		return
	}
	t.Log("user exist.......")

}
