package nologin

import (
	"gin-practice/src/auth/rbac"
	"github.com/gin-gonic/gin"
)

func Method(c *gin.Context) {

	e := rbac.Enforcer
	ok, err := e.Enforce("test", "/v1/test", "GET")
	if err != nil {
		return
	}
	if !ok {
		c.JSON(200, gin.H{
			"code":    4001,
			"message": "没有权限",
		})
		return
	}

	path := c.FullPath()
	c.JSON(200, gin.H{
		"fullPath": path,
	})

}
