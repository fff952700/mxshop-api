package initialize

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"mxshop-api/user-web/global"
	"mxshop-api/user-web/proto"
)

func InitUserClient() {
	// 使用 gRPC 客户端 API 创建连接
	conn, err := grpc.NewClient(
		fmt.Sprintf("%s:%d", global.ServerConf.UserServerInfo.Host, global.ServerConf.UserServerInfo.Port),
		grpc.WithTransportCredentials(insecure.NewCredentials()), // 不加密连接
	)

	if err != nil {
		zap.S().Panicw("init UserClient failed", "msg", err.Error())
		return
	}
	// 连接成功后，设置全局的 UserClient
	global.UserClient = proto.NewUserClient(conn)
	// 尝试调用一个简单的 RPC 方法来确认连接成功
	_, err = global.UserClient.GetUserByMobile(context.Background(), &proto.MobileRequest{
		Mobile: "13666666666", // 使用一个测试手机号
	})

	if err != nil {
		zap.S().Panicw("init UserClient failed", "msg", err.Error())
		return
	}

}
