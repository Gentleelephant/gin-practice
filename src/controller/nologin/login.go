package nologin

import (
	"gin-practice/src/auth"
	"gin-practice/src/common"
	"gin-practice/src/config"
	"gin-practice/src/dao"
	"gin-practice/src/model"
	"gin-practice/src/utils"
	"github.com/gin-gonic/gin"
	"net/http"
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
		return
	}
	//TODO 启用数据库登陆
	checkUser, err := dao.UserDao.GetUserByName(user.Username)
	if err != nil {
		return
	}
	checkPassword := checkUser.Password
	if err != nil {
		return
	}
	verify := utils.Verify(checkPassword, user.Password)
	if !verify {
		err = common.PasswordWrongError
		return
	}

	// 登陆成功向cookie中写入sessionid，想redis中写入sessionid
	c.JSON(http.StatusOK, gin.H{
		"code":    2000,
		"message": "登陆成功",
	})
}
