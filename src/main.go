package main

import (
	"gin-practice/src/auth/rbac"
	"gin-practice/src/cache"
	"gin-practice/src/global"
	"gin-practice/src/initialize"
	"gin-practice/src/routers"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {

	engine := gin.Default()

	// 注册中间件
	routers.RegisterMiddlware(engine)

	// 注册路由
	routers.RegisterRouter(engine)

	err := initialize.LoadConfig()
	if err != nil {
		return
	}
	// 初始化redis
	cache.InitRedis()

	// 初始化Casbin
	rbac.InitCasbin()

	// 初始化数据库
	initialize.InitDB()

	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	err = engine.Run(global.GlobalConfig.Server.Host + ":" + global.GlobalConfig.Server.Port)
	if err != nil {
		return
	}

}
