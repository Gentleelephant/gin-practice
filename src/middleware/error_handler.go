package middleware

import (
	"gin-practice/src/common"
	"github.com/gin-gonic/gin"
	"net/http"
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
		var err *common.CommonError
		switch et := ginErr.Err.(type) {
		case common.CommonError:
			err = &et
		case *common.CommonError:
			err = et
		default:
			err = common.ServerError
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
	}
}
