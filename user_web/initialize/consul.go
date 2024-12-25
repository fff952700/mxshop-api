package initialize

// InitConsul 从consul获取service user-srv 是用lb过后则不需要这部分
//func init() {
//	config := api.DefaultConfig()
//	ConsulInfo := global.Cfg.ConsulInfo
//	config.Address = fmt.Sprintf("%s:%d", ConsulInfo.Host, ConsulInfo.Port)
//	client, err := api.NewClient(config)
//	if err != nil {
//		zap.S().Panicw("init consul fail", "err", err)
//	}
//	// 通过agent filter获取service
//	data, err := client.Agent().ServicesWithFilter(fmt.Sprintf(`Service == "%s"`, ConsulInfo.TargetServerName))
//	if err != nil {
//		zap.S().Panicw("Filter consul Service fail", "err", err)
//	}
//	zap.S().Infow("Filter consul Service success", "data", data)
//	for _, value := range data {
//		// 通过consul 获取grpc的地址
//		global.Cfg.UserServerInfo.Host = value.Address
//		global.Cfg.UserServerInfo.Port = value.Port
//	}
//}
