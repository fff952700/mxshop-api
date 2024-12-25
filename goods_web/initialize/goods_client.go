package initialize

import (
	"context"
	"fmt"

	_ "github.com/mbobakov/grpc-consul-resolver" // It's important
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"mxshop_api/goods_web/global"
	"mxshop_api/goods_web/proto"
)

func init() {
	cfg := global.Cfg.ConsulInfo
	conn, err := grpc.NewClient(fmt.Sprintf("consul://%s:%d/%s?wait=%s&tag=%s", cfg.Host, cfg.Port,
		cfg.TargetServerName, "14s", cfg.Target),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`),
	)
	if err != nil {
		zap.S().Panicw("Init Client Conn Err", "error", err.Error())
		return
	}
	goodsClient := proto.NewGoodsClient(conn)
	_, err = goodsClient.GetGoodsDetail(context.Background(), &proto.GoodInfoRequest{Id: 421})

	if err != nil {
		zap.S().Panicw("Init UserClient Err", "error", err.Error())
		return
	}
	global.GoodsClient = goodsClient

}
