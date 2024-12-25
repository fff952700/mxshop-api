package tests

import (
	"fmt"
	"testing"

	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
)

func TestNacos(t *testing.T) {
	// 创建clientConfig
	clientConfig := constant.ClientConfig{
		NamespaceId:         "09d89590-bb9a-4218-8b67-db9ac8364bba", // 如果需要支持多namespace，我们可以创建多个client,它们有不同的NamespaceId。当namespace是public时，此处填空字符串。
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "nacos/log",
		CacheDir:            "nacos/cache",
		LogLevel:            "debug",
	}
	// 至少一个Cfgig
	Cfgigs := []constant.ServerConfig{
		{
			IpAddr:      "192.168.2.105",
			ContextPath: "/nacos",
			Port:        8848,   // 使用 HTTP 端口
			Scheme:      "http", // 强制使用 HTTP 协议
		},
	}

	// 创建动态配置客户端的另一种方式 (推荐)
	configClient, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: Cfgigs,
		},
	)
	if err != nil {
		panic(err)
	}
	// 获取配置

	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: "user-api",
		Group:  "dev",
	})
	if err != nil {
		t.Fatalf("Failed to get config: %v", err) // 使用t.Fatalf便于调试，输出详细错误信息
	}
	fmt.Println("content:", content)

}
