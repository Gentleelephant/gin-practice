package routers

import (
	"gin-practice/src/controller"
	"gin-practice/src/middleware"
	"github.com/gin-gonic/gin"
)

// RegisterMiddlware 注册路由中间件
func RegisterMiddlware(engine *gin.Engine) {

	// 注册验证登录中间件
	engine.Use(middleware.CheckSessions)

}

// 注册路由
func RegisterRouter(engine *gin.Engine) {

	v2 := engine.Group("/v2")
	{
		v2.POST("/login", controller.Login)
		v2.GET("/hello", controller.Hello)
		v2.POST("/register", controller.Register)
	}

}
