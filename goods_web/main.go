package main

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/v2/inner/uuid"
	"mxshop_api/goods_web/utils/register/consul"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"mxshop_api/goods_web/global"
	"mxshop_api/goods_web/initialize"
	"mxshop_api/goods_web/utils"
)

func main() {
	// 设置 Gin 为 Release 模式
	gin.SetMode(gin.ReleaseMode)
	// 初始化router
	routers := initialize.Routers()
	svc := global.Cfg.ServerInfo
	consulInfo := global.Cfg.ConsulInfo
	// consul client
	client := consul.NewRegisterClient(consulInfo.Host, consulInfo.Port)
	// 服务发现
	go func() {
		// 模拟本地服务多部署，使用随机端口，正式使用k8s则无需使用 通过yapi测试不方便所有不开启
		debug := initialize.GetEnvInfo("MXSHOP_DEBUG")
		if !debug {
			// 使用随机可用端口
			port, err := utils.GetFreePort()
			if err != nil {
				zap.S().Panicf("%s service not port err %v", svc.Name, err.Error())
			}
			global.Cfg.ServerInfo.Port = port
		}
		id, err := uuid.NewV4()
		if err != nil {
			zap.S().Panic(err)
		}
		global.Cfg.ServerInfo.Id = id.String()
		err = client.Register()
		if err != nil {
			zap.S().Panicf("%s register err %v", svc.Name, err.Error())
		}
		zap.S().Infof("%s service start succees %s:%d", svc.Name, svc.Host, svc.Port)
		if err = routers.Run(fmt.Sprintf(":%d", svc.Port)); err != nil {
			zap.S().Panicf("%s service start error %v", svc.Name, err.Error())
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	err := client.Deregister()
	if err != nil {
		zap.S().Errorf("%s Failed to deregister service %T", svc.Name, zap.Error(err))
	} else {
		zap.S().Infof("%s Service deregistered successfully", svc.Name)
	}
}
