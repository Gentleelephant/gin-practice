package v1

import (
	"gin-practice/src/auth"
	"gin-practice/src/cache"
	"gin-practice/src/common"
	"gin-practice/src/global"
	"gin-practice/src/model"
	"gin-practice/src/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Login(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			_ = c.Error(err)
		}
	}()

	user := model.LoginUserDTO{}
	err = c.Bind(&user)
	if err != nil {
		return
	}
	// 创建连接池
	pool, err := auth.NewPool(global.GlobalConfig.LDAP.NewLdapClient, 10)
	if err != nil {
		return
	}
	// 判断是否启用LDAP登陆
	if global.GlobalConfig.LDAP.Enabled {
		checkUser := auth.CheckUser{
			Pool:            pool,
			ManagerDN:       global.GlobalConfig.LDAP.ManagerDN,
			ManagerPassword: global.GlobalConfig.LDAP.ManagerPassword,
			UserSearchBase:  global.GlobalConfig.LDAP.UserSearchBase,
			LoginAttribute:  global.GlobalConfig.LDAP.LoginAttribute,
			MailAttribute:   global.GlobalConfig.LDAP.MailAttribute,
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
	checkUser, err := global.UserDao.GetUserByName(user.Username)
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
	randomString, err := utils.GenerateRandomString(64)
	if err != nil {
		return
	}
	err = cache.RedisClient.SetSession(string(checkUser.ID), randomString, time.Second*360)
	if err != nil {
		return
	}
	err = cache.RedisClient.SetUserInfo(randomString, checkUser, time.Second*360)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    2000,
		"message": "登陆成功",
	})
}
