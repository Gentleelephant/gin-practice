package middleware

import (
	"gin-practice/src/auth/rbac"
	"github.com/gin-gonic/gin"
)

func RBAC(c *gin.Context) {

	path := c.FullPath()
	method := c.Request.Method
	// 根据cookie的到用户
	user, t := c.Get("user")
	if t != true {
		user = "default"
	}
	e := rbac.Enforcer

	ok, err := e.Enforce(user.(string), path, method)
	if err != nil {
		c.Abort()
		return
	}
	if !ok {
		c.JSON(200, gin.H{
			"code":    4009,
			"message": "没有权限",
		})
		c.Abort()
	}

}
