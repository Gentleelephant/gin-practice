package main

import (
	"gin-practice/src/config"
	"gin-practice/src/dao"
	"gin-practice/src/routers"
	"github.com/gin-gonic/gin"
)

func main() {

	engine := gin.Default()

	// 初始化配置
	config.InitConfig()

	// 初始化数据库
	dao.InitDB()

	// 注册中间件
	routers.RegisterMiddlware(engine)

	// 注册路由
	routers.RegisterRouter(engine)

	err := engine.Run(config.GolbalConfig.Server.Host + ":" + config.GolbalConfig.Server.Port)
	if err != nil {
		return
	}

}
