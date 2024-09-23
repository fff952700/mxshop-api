package initialize

import (
	"context"
	"fmt"
	"github.com/hashicorp/consul/api"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"mxshop-api/user-web/global"
	"mxshop-api/user-web/proto"
)

func InitUserClient() {
	// 通过consul获取
	cfg := api.DefaultConfig()
	cfg.Address = fmt.Sprintf("%s:%d", global.ServerConf.ConsulInfo.Host, global.ServerConf.ConsulInfo.Port)
	client, err := api.NewClient(cfg)
	if err != nil {
		zap.S().Panicw("Init Consul Client Failed", "msg", err.Error())
	}
	data, err := client.Agent().ServicesWithFilter(fmt.Sprintf("Service == \"%s\"", global.ServerConf.ConsulInfo.SrvName))
	if err != nil {
		zap.S().Panicw("Init Consul filter Failed", "msg", err.Error())
	}
	userSrvHost := ""
	userSrvPort := 0
	for _, v := range data {
		userSrvHost = v.Address // 需要在服务注册的时候传入address和port 否则读取为nil
		userSrvPort = v.Port
		break
	}
	if userSrvHost == "" {
		zap.S().Fatalf("Init Consul filter Failed")
		return
	}
	// 使用 gRPC 客户端 API 创建连接
	conn, err := grpc.NewClient(
		fmt.Sprintf("%s:%d", userSrvHost, userSrvPort),
		grpc.WithTransportCredentials(insecure.NewCredentials()), // 不加密连接
	)

	if err != nil {
		zap.S().Panicw("Init UserClient Failed", "msg", err.Error())
		return
	}
	// 连接成功后，设置全局的 UserClient
	global.UserClient = proto.NewUserClient(conn)
	// 尝试调用一个简单的 RPC 方法来确认连接成功
	_, err = global.UserClient.GetUserByMobile(context.Background(), &proto.MobileRequest{
		Mobile: "13666666666", // 使用一个测试手机号
	})

	if err != nil {
		zap.S().Panicw("Init UserClient Failed", "msg", err.Error())
		return
	}

}
