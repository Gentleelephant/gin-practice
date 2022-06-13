package middleware

import (
	"github.com/gin-gonic/gin"
)

//TODO 错误处理中间件
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Call next handler
		c.Next()

		// Handle gin Context errors
		ginErr := c.Errors.Last()
		if ginErr == nil {
			return
		}

	}
}
