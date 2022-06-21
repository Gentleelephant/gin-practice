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
	//engine.Use(middleware.CheckSessions)
	engine.Use(middleware.JWTAuth())

}

// 注册路由
func RegisterRouter(engine *gin.Engine) {

	docs.SwaggerInfo.BasePath = "/v1"
	g1 := engine.Group("/v1")
	{
		g1.GET("/method", v1.Method)
		g1.POST("/login", v1.Login)
		g1.GET("/hello", v1.Hello)
		g1.POST("/register", v1.Register)
		g1.GET("/getUser", v1.GetUser)
		g1.POST("/check_email", v1.EmailCaptcha)
		g1.GET("/captcha", v1.GetImageCaptcha)
	}

}
