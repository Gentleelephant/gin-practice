package nologin

import (
	"gin-practice/src/auth"
	"gin-practice/src/config"
	"gin-practice/src/model"
	"github.com/gin-gonic/gin"
	"log"
)

func Login(c *gin.Context) {
	user := model.LoginUser{}
	err := c.Bind(&user)
	log.Println(user)
	if err != nil {
		c.JSON(200, gin.H{
			"code":    4001,
			"message": "参数错误",
		})
		return
	}
	// 创建连接池，本来连接吃因嘎嘎i是个公用属性
	pool, err := auth.NewPool(config.GlobalConfig.LDAP.NewLdapClient, 10)
	if err != nil {
		return
	}
	// 判断是否启用LDAP登陆
	if config.GlobalConfig.LDAP.Enabled {
		checkUser := auth.CheckUser{
			pool,
			config.GlobalConfig.LDAP.ManagerDN,
			config.GlobalConfig.LDAP.ManagerPassword,
			config.GlobalConfig.LDAP.UserSearchBase,
			config.GlobalConfig.LDAP.LoginAttribute,
			config.GlobalConfig.LDAP.MailAttribute,
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
	// 启用数据库登陆

}
