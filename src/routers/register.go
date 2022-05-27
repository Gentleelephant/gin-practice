package routers

import (
	"gin-practice/src/middleware"
	"github.com/gin-gonic/gin"
)

// RegisterMiddlware 注册路由中间件
func RegisterMiddlware(engine *gin.Engine) {

	// 注册验证登录中间件
	engine.Use(middleware.CheckSessions)

}

// 注册路由
