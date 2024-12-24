package initialize

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"

	"mxshop_api/user_web/global"
	"mxshop_api/user_web/storage"
)

func init() {
	Client := redis.NewClient(
		&redis.Options{
			Addr: fmt.Sprintf("%s:%d", global.ServerConf.RedisInfo.Host, global.ServerConf.RedisInfo.Port),
			DB:   global.ServerConf.RedisInfo.DB,
		})
	_, err := Client.Ping(context.Background()).Result()
	if err != nil {
		zap.S().Panicw("ping redis failed", "err", err.Error())
	}
	global.RedisClient = Client
	// 初始化存储桶
	// 解析时间
	ExpirationTime, err := time.ParseDuration(global.ServerConf.RedisInfo.ExpirationTime)
	if err != nil {
		zap.S().Panicw("parse expiration time failed", "err", err.Error())
	}
	redisStore := storage.NewRedisStore(Client, ExpirationTime)
	global.RedisStore = redisStore
}
