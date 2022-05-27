package middleware

import (
	"gin-practice/src/entity"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func CheckSessions(c *gin.Context) {

	if strings.HasPrefix(c.Request.URL.Path, "/login") && c.Request.Method == "POST" {
		c.Next()
		return
	}

	// 目前并不需要cookie来干什么
	// 后面来认证
	_, err := c.Cookie("token")
	if err != nil {
		c.JSON(http.StatusOK, entity.CustomResp{
			Code: 4000,
			Msg:  "请先登录",
			Data: err.Error(),
		})
		c.Abort()
	}
}
