package routers

import (
	"gin-practice/src/api/v1"
	"gin-practice/src/docs"
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

	docs.SwaggerInfo.BasePath = "/v2"
	v2 := engine.Group("/v2")
	{
		v2.GET("/method", v1.Method)
		v2.POST("/login", v1.Login)
		v2.GET("/hello", v1.Hello)
		v2.POST("/register", v1.Register)
	}

}
