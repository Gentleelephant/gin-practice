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
			"code":    -1,
			"message": "参数错误",
		})
		return
	}
	// 判断是否启用LDAP登陆
	if config.GlobalConfig.LDAP.Enabled {
		provider := auth.GetProvider()
		check := provider.Authentication(user.Username, user.Password)
		if !check {
			c.JSON(200, gin.H{
				"code":    -1,
				"message": "用户名或密码错误",
			})
			return
		} else {
			c.SetCookie("token", user.Username, 3600, "/", "", false, false)
			c.JSON(200, gin.H{
				"code":    0,
				"message": "登陆成功",
			})
			return
		}
	}
}
