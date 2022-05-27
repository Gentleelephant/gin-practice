package main

import (
	"gin-practice/src/routers"
	"github.com/gin-gonic/gin"
)

func main() {

	engine := gin.Default()

	// 注册中间件
	routers.RegisterMiddlware(engine)

	// 注册路由
	routers.RegisterRouter(engine)

	engine.Run(":8080")

}
