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
	// 判断是否启用LDAP登陆
	if config.GlobalConfig.LDAP.Enabled {
		provider := auth.GetProvider()
		authentication, err := provider.Authentication(user.Username, user.Password)
		check := authentication
		if !check {
			c.JSON(200, gin.H{
				"code":    -1,
				"message": err.Error(),
			})
			return
		} else {
			c.SetCookie("token", user.Username, 3600, "/", "", false, false)
			c.JSON(200, gin.H{
				"code":    2000,
				"message": "登陆成功",
			})
			return
		}
	}
	// 启用数据库登陆

}
