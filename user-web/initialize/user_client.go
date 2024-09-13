package initialize

import (
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"mxshop-api/user-web/global"
	"mxshop-api/user-web/proto"
)

func InitUserClient() {
	conn, err := grpc.NewClient(fmt.Sprintf("%s:%d", global.ServerConf.UserServerInfo.Host, global.ServerConf.UserServerInfo.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		zap.S().Panicw("init UserClient failed", "msg", err.Error())
	}
	global.UserClient = proto.NewUserClient(conn)
}
