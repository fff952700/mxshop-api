package tests

import (
	"go.uber.org/zap"
	"testing"
	"time"
)

func TestUseZap(t *testing.T) {
	logger, _ := zap.NewProduction()
	//logger, _ := zap.NewDevelopment()
	// 生成环境和开发环境输出的日志级别和格式不一样
	defer logger.Sync() // flushes buffer, if any
	// 使用sugar底层做了映射。性能没有直接使用好
	//sugar := logger.Sugar()
	url := "www.baidu.com"
	//sugar.Infow("failed to fetch URL",
	//	// Structured context as loosely typed key-value pairs.
	//	"url", url,
	//	"attempt", 3,
	//	"backoff", time.Second,
	//)
	//sugar.Infof("Failed to fetch URL: %s", url)
	logger.Info("failed to fetch URL",
		// Structured context as strongly typed Field values.
		zap.String("url", url),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)
}

func TestZapStdFile(t *testing.T) {
	// 自定义配置文件
	cfg := zap.NewProductionConfig()
	// 切片类型可以输出到多个地方
	cfg.OutputPaths = []string{
		"stdout",
		"./myZap.log",
	}
	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}
	url := "www.baidu.com"
	logger.Info("failed to fetch URL",
		zap.String("url", url),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second))
}
