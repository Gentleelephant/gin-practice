package middleware

import (
	"gin-practice/src/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func CheckSessions(c *gin.Context) {

	if strings.HasPrefix(c.Request.URL.Path, "/v2/login") && c.Request.Method == "POST" {
		c.Next()
		return
	}

	if strings.HasPrefix(c.Request.URL.Path, "/swagger") && c.Request.Method == "GET" {
		c.Next()
		return
	}

	if strings.HasPrefix(c.Request.URL.Path, "/v2/register") && c.Request.Method == "POST" {
		c.Next()
		return
	}

	if strings.HasPrefix(c.Request.URL.Path, "/v2/method") && c.Request.Method == "GET" {
		c.Next()
		return
	}

	//

	// 后面来认证
	_, err := c.Cookie("token")
	if err != nil {
		c.JSON(http.StatusOK, model.CustomResp{
			Code: 4000,
			Msg:  "请先登录",
			Data: err.Error(),
		})
		c.Abort()
	}
}
