package routers

import (
	"gin-practice/src/controller/needlogin"
	"gin-practice/src/controller/nologin"
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
		v2.GET("/method", nologin.Method)
		v2.POST("/login", nologin.Login)
		v2.GET("/hello", needlogin.Hello)
		v2.POST("/register", nologin.Register)
	}

}
