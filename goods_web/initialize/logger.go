package initialize

import "go.uber.org/zap"

func init() {
	// 1.初始化全局zap
	logger, _ := zap.NewDevelopment()
	zap.ReplaceGlobals(logger)
}
