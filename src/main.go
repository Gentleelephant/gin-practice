package main

import (
	"gin-practice/src/config"
	"gin-practice/src/routers"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {

	engine := gin.Default()

	// 注册中间件
	routers.RegisterMiddlware(engine)

	// 注册路由
	routers.RegisterRouter(engine)

	// 加载配置
	loadConfig, err := config.LoadConfig("C:\\work\\code\\goPro\\gin-practice\\src\\config\\config.yaml")
	if err != nil {
		log.Fatalln("load config error:", err)
		return
	}

	err = engine.Run(loadConfig.Server.Host + ":" + loadConfig.Server.Port)
	if err != nil {
		log.Fatalln("run server error:", err)
		return
	}

}
