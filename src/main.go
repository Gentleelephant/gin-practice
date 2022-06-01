package main

import (
	"fmt"
	"gin-practice/src/config"
	"gin-practice/src/dao"
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

	err := config.LoadConfig(config.ConfigPath)
	if err != nil {
		return
	}

	fmt.Println("config", config.GlobalConfig)

	db := dao.InitDB()
	config.DB = db

	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	err = engine.Run(config.GlobalConfig.Server.Host + ":" + config.GlobalConfig.Server.Port)
	if err != nil {
		return
	}

}
