package middleware

import (
	"gin-practice/src/jwt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// JWTAuth 中间件，检查token

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {

		if c.Request.URL.Path == "/v1/login" && c.Request.Method == "POST" {
			c.Next()
			return
		}

		if c.Request.URL.Path == "/v1/hello" && c.Request.Method == "GET" {
			c.Next()
			return
		}

		if c.Request.URL.Path == "/v1/register" && c.Request.Method == "POST" {
			c.Next()
			return
		}

		if c.Request.URL.Path == "/v1/captcha" && c.Request.Method == "GET" {
			c.Next()
			return
		}

		if c.Request.URL.Path == "/v1/health_check" && c.Request.Method == "GET" {
			c.Next()
			return
		}

		// 通过http header中的token解析来认证
		token := c.Request.Header.Get("Token")
		if token == "" || len(token) == 0 {
			c.JSON(http.StatusOK, gin.H{
				"status": -1,
				"msg":    "请求未携带token，无权限访问",
				"data":   nil,
			})
			c.Abort()
			return
		}

		log.Print("get token: ", token)

		// 初始化一个JWT对象实例，并根据结构体方法来解析token
		j := jwt.NewJWT()
		// 解析token中包含的相关信息(有效载荷)
		claims, err := j.ParserToken(token)

		if err != nil {
			switch err.Error() {
			case "Token is expired":
				c.JSON(http.StatusOK, gin.H{
					"status": -1,
					"msg":    "token授权已过期，请重新申请授权",
					"data":   nil,
				})
				c.Abort()
				return
			case "Token is invalid":
				c.JSON(http.StatusOK, gin.H{
					"status": -1,
					"msg":    "无效的token，请重新申请授权",
					"data":   nil,
				})
				c.Abort()
				return
			case "Token is not valid":
				c.JSON(http.StatusOK, gin.H{
					"status": -1,
					"msg":    "token不可用，请重新申请授权",
					"data":   nil,
				})
				c.Abort()
				return
			default:
				// 其他错误
				c.JSON(http.StatusOK, gin.H{
					"status": -1,
					"msg":    err.Error(),
					"data":   nil,
				})
				c.Abort()
				return
			}
		}
		// 将解析后的有效载荷claims重新写入gin.Context引用对象中
		c.Set("claims", claims)

	}
}
