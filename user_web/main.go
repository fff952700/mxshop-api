package main

import (
	"fmt"
	"go.uber.org/zap"
	"mxshop_api/user_web/global"
	"mxshop_api/user_web/initialize"
	"mxshop_api/user_web/utils"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// 1. 初始化zap日志
	initialize.InitLogger()

	// 2.获取配置文件
	initialize.InitConfig()

	// 3.加载时区，在jwt验证token使用
	initialize.InitTimeZone()
	// 3. 初始consul 获取svc客户端连接地址
	initialize.InitConsul()
	// 3. 初始化svc客户端连接
	initialize.InitUserClient()

	// 4.初始化router
	routers := initialize.Routers()

	// 5.初始化翻译器
	if err := initialize.InitValidator("zh"); err != nil {
		zap.L().Panic("init validator failed", zap.Error(err))
	}

	// 6.初始化redis
	initialize.InitRedis()

	// 7.服务发现
	go func() {
		// 7.模拟本地服务多部署，使用随机端口，正式使用k8s则无需使用 通过yapi测试不方便所有不开启
		debug := initialize.GetEnvInfo("MXSHOP_DEBUG")
		if !debug {
			// 使用随机可用端口
			port, err := utils.GetFreePort()
			if err != nil {
				zap.S().Panicw("service not port", "msg", err.Error())
			}
			global.ServerConf.ServerPort = port
		}
		zap.S().Infof("user-api 服务器成功启动在 :%d", global.ServerConf.ServerPort)
		if err := routers.Run(fmt.Sprintf(":%d", global.ServerConf.ServerPort)); err != nil {
			zap.S().Panicw("service start error", "msg", err.Error())
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
}
