package main

import (
	"fmt"
	"go.uber.org/zap"
	"mxshop-api/user-web/global"
	"mxshop-api/user-web/initialize"
)

func main() {
	// 启动服务
	//ip := flag.String("ip", "127.0.0.1", "ip address")
	//port := flag.Int("port", 8082, "port number")
	//flag.Parse()

	// 1. 初始化zap日志
	initialize.InitLogger()

	// 2.获取配置文件
	initialize.InitConfig()

	// 3. 初始化svc客户端连接
	initialize.InitUserClient()

	// 4.初始化router
	initialize.Routers()

	// 5.初始化翻译
	if err := initialize.InitValidator("zh"); err != nil {
		zap.S().Panicw("init validator failed", "msg", err.Error())
	}

	// 6.初始化redis
	initialize.InitRedis()

	routers := initialize.Routers()
	if err := routers.Run(fmt.Sprintf(":%d", global.ServerConf.ServerPort)); err != nil {
		zap.S().Panicw("service start error", "msg", err.Error())
	}
}
