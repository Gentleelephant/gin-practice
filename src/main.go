package main

import (
	"gin-practice/src/controller"
	"gin-practice/src/routers"
	"github.com/gin-gonic/gin"
)

func main() {

	engine := gin.Default()

	// 注册路由
	routers.RegisterMiddlware(engine)

	engine.POST("/login", controller.Login)

	engine.GET("/hello", controller.Hello)

	engine.Run(":8080")

}
