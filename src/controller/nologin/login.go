package nologin

import (
	"gin-practice/src/auth"
	"gin-practice/src/config"
	"gin-practice/src/dao"
	"gin-practice/src/entity"
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
		authentication, dn, err := provider.Authentication(user.Username, user.Password)
		check := authentication
		if !check {
			c.JSON(200, gin.H{
				"code":    -1,
				"message": err,
			})
			return
		} else {
			c.SetCookie("token", user.Username, 3600, "/", "", false, false)
			c.JSON(200, gin.H{
				"code":    2000,
				"message": "登陆成功",
			})
			// 判断用户是否存在，不存在则创建
			// 判断表中是否以及包含该用户
			user := &entity.User{
				Username: user.Username,
			}

			take := dao.CheckUser(user)
			if take != 1 {
				log.Println("user not exist, insert user......")
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
			return
		}
	}
	// 启用数据库登陆

}
