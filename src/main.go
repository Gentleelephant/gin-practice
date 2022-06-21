package main

import (
	"gin-practice/src/auth/rbac"
	"gin-practice/src/global"
	"gin-practice/src/initialize"
	"gin-practice/src/routers"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"time"
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
	initialize.InitRedis()

	initialize.InitEmailPool()

	// 初始化Casbin
	rbac.InitCasbin()

	// 初始化数据库
	initialize.InitDB()

	// 持续监听邮件列表
	go func() {
		for {
			select {
			case mail := <-global.EmailLists:
				// 发送邮件
				err = global.Pool.Send(mail, time.Second*10)
				log.Println(err)
				log.Println("email:", mail)
			default:
			}
		}

	}()

	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	err = engine.Run(global.GlobalConfig.Server.Host + ":" + global.GlobalConfig.Server.Port)
	if err != nil {
		return
	}

}
