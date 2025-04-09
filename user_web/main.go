package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"mxshop_api/user_web/global"
	"mxshop_api/user_web/initialize"
	"mxshop_api/user_web/utils"
)

func main() {
	// 设置 Gin 为 Release 模式
	gin.SetMode(gin.ReleaseMode)
	// 初始化router
	routers := initialize.Routers()
	initialize.InitValidator("zh")
	// 服务发现
	go func() {
		// 模拟本地服务多部署，使用随机端口，正式使用k8s则无需使用 通过yapi测试不方便所有不开启
		debug := initialize.GetEnvInfo("MXSHOP_DEBUG")
		serverName := global.Cfg.ServerInfo.Name
		if !debug {
			// 使用随机可用端口
			port, err := utils.GetFreePort()
			if err != nil {
				zap.S().Panicf("%s service not port err %v", serverName, err.Error())
			}
			global.Cfg.ServerInfo.Port = port
		}
		svc := global.Cfg.ServerInfo
		zap.S().Infof("%s service start succees %s:%d", svc.Name, svc.Host, svc.Port)
		if err := routers.Run(fmt.Sprintf(":%d", svc.Port)); err != nil {
			zap.S().Panicf("%s service start error %v", serverName, err.Error())
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
}
