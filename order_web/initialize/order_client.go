package initialize

import (
	"context"
	"fmt"

	_ "github.com/mbobakov/grpc-consul-resolver" // It's important
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"mxshop_api/order_web/global"
	"mxshop_api/order_web/proto"
)

func init() {
	cfg := global.Cfg.ConsulInfo
	conn, err := grpc.NewClient(fmt.Sprintf("consul://%s:%d/%s?wait=%s&tag=%s", cfg.Host, cfg.Port,
		cfg.TargetServerName, "14s", cfg.Target),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`),
	)
	if err != nil {
		zap.S().Panicw("init client conn err", "error", err.Error())
		return
	}
	orderClient := proto.NewOrderClient(conn)
	_, err = orderClient.OrderDetail(context.Background(), &proto.OrderRequest{Id: 2})

	if err != nil {
		zap.S().Panicw("init orderClient err", "error", err.Error())
		return
	}
	global.OrderClient = orderClient

}
