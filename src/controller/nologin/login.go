package nologin

import (
	"gin-practice/src/auth"
	"gin-practice/src/config"
	"gin-practice/src/model"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			_ = c.Error(err)
		}
	}()
	user := model.LoginUser{}
	err = c.Bind(&user)
	if err != nil {
		return
	}
	// 创建连接池
	pool, err := auth.NewPool(config.GlobalConfig.LDAP.NewLdapClient, 10)
	if err != nil {
		return
	}
	// 判断是否启用LDAP登陆
	if config.GlobalConfig.LDAP.Enabled {
		checkUser := auth.CheckUser{
			Pool:            pool,
			ManagerDN:       config.GlobalConfig.LDAP.ManagerDN,
			ManagerPassword: config.GlobalConfig.LDAP.ManagerPassword,
			UserSearchBase:  config.GlobalConfig.LDAP.UserSearchBase,
			LoginAttribute:  config.GlobalConfig.LDAP.LoginAttribute,
			MailAttribute:   config.GlobalConfig.LDAP.MailAttribute,
		}
		authentication, err := checkUser.Authentication(user.Username, user.Password)
		if err != nil {
			return
		}
		c.JSON(200, gin.H{
			"code":    0,
			"message": "登陆成功",
			"data":    authentication,
		})
	}
	//TODO 启用数据库登陆

}
