package initialize

import (
	"encoding/json"
	"fmt"

	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	"mxshop_api/user_web/global"
)

// InitConfig 通过先通过viper获取本地nacos配置在获取服务配置
func init() {
	v := viper.New()
	debug := GetEnvInfo("MXSHOP_DEBUG")
	configFilePrefix := "config"
	configFileName := fmt.Sprintf("user_web/%s_pro.yaml", configFilePrefix)
	if debug {
		configFileName = fmt.Sprintf("user_web/%s_prev.yaml", configFilePrefix)
	}
	v.SetConfigFile(configFileName)
	if err := v.ReadInConfig(); err != nil {
		zap.S().Panicw("read config failed", "err", err)
	}
	// 实例化配置
	if err := v.Unmarshal(global.NacosConf); err != nil {
		zap.S().Panicw("unmarshal config failed", "err", err)
	}
	zap.S().Infof("consul config :%v", global.NacosConf)

	// 创建clientConfig
	cc := constant.ClientConfig{
		NamespaceId:         global.NacosConf.Namespace, // 如果需要支持多namespace，我们可以创建多个client,它们有不同的NamespaceId。当namespace是public时，此处填空字符串。
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "./user_web/consul/log",
		CacheDir:            "./user_web/consul/cache",
		LogLevel:            "debug",
	}
	// 至少一个Cfgig
	sc := []constant.ServerConfig{
		{
			IpAddr:      global.NacosConf.Host,
			ContextPath: "/consul",
			Port:        uint64(global.NacosConf.Port), // 使用 HTTP 端口
			Scheme:      global.NacosConf.Scheme,       // 强制使用 HTTP 协议
		},
	}

	// 创建动态配置客户端的另一种方式 (推荐)
	configClient, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)
	if err != nil {
		zap.S().Panicw("create config client failed", "err", err)
	}
	// 获取配置

	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: global.NacosConf.DataId,
		Group:  global.NacosConf.Group,
	})
	if err != nil {
		zap.S().Panicw("get config failed", "err", err)
	}
	// 监听配置变化
	if err = configClient.ListenConfig(vo.ConfigParam{
		DataId: global.NacosConf.DataId,
		Group:  global.NacosConf.Group,
		OnChange: func(namespace, group, dataId, data string) {
			fmt.Println("group:" + group + ", dataId:" + dataId + ", data:" + data)
		},
	}); err != nil {
		zap.S().Panicw("config listen failed", "err", err)
	}

	// 将json序列化为struct
	// 实例化配置对象
	Cfgig := global.Cfg
	if err = json.Unmarshal([]byte(content), &Cfgig); err != nil {
		zap.S().Panicw("unmarshal config failed", "err", err)
	}
	zap.S().Infow("server config", "Cfgig", Cfgig)
}

// 通过viper监听本地配置文件
//func InitConfig() {
//	debug := GetEnvInfo("MXSHOP_DEBUG")
//	configFilePrefix := "config"
//	configFileName := fmt.Sprintf("user_web/%s_pro.yaml", configFilePrefix)
//	if debug {
//		configFileName = fmt.Sprintf("user_web/%s_debug.yaml", configFilePrefix)
//	}
//	v := viper.New()
//	// 设置文件路径
//	v.SetConfigFile(configFileName)
//	if err := v.ReadInConfig(); err != nil {
//		panic(err)
//	}
//	// 实例化配置结构体
//	Cfgig := global.Cfg
//	if err := v.Unmarshal(Cfgig); err != nil {
//		panic(err)
//	}
//	// viper 动态监听变化
//	v.WatchConfig()
//	v.OnConfigChange(func(e fsnotify.Event) {
//		// 重新读取
//		_ = v.ReadInConfig()
//		// 重新解析
//		_ = v.Unmarshal(Cfgig)
//		fmt.Printf("%+v\n", Cfgig)
//
//	})
//}

// GetEnvInfo 获取环境变量
func GetEnvInfo(env string) bool {
	viper.AutomaticEnv()
	return viper.GetBool(env)
}
