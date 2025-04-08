package consul

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"go.uber.org/zap"
	"mxshop_api/order_web/global"
)

// 定义consul注册
type RegisterClient struct {
	Host string
	Port int
}

func NewRegisterClient(host string, port int) *RegisterClient {
	return &RegisterClient{
		Host: host,
		Port: port,
	}
}

// 生成注册接口
type Register interface {
	Register() error
	Deregister() error
}

// 实现接口方法
func (r *RegisterClient) Register() error {
	serverInfo := global.Cfg.ServerInfo
	// 实例化consul对象
	config := api.DefaultConfig()
	config.Address = fmt.Sprintf("%s:%d", r.Host, r.Port)
	client, err := api.NewClient(config)
	if err != nil {
		zap.S().Panicw("[InitConsul] init consul fail")

	}

	zap.S().Infow("host %s port %d", serverInfo.Host, serverInfo.Port)
	// 健康检查
	check := &api.AgentServiceCheck{
		HTTP:     fmt.Sprintf("http://%s:%d/health", serverInfo.Host, serverInfo.Port), //
		Timeout:  "5s",                                                                 // 超时时间
		Interval: "5s",                                                                 // 运行检查的频率
		// 指定时间后自动注销不健康的服务节点
		DeregisterCriticalServiceAfter: "15s",
	}
	// 注册consul中的信息 id相同在consul会认为是一个
	registration := &api.AgentServiceRegistration{
		ID:      serverInfo.Id,   // 服务唯一ID
		Name:    serverInfo.Name, // 服务名称
		Tags:    serverInfo.Tag,  // 为服务打标签
		Address: serverInfo.Host,
		Port:    serverInfo.Port,
		Check:   check,
	}
	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		zap.S().Panicw("[InitConsul] register service fail", err)
	}
	global.ConsulClient = client
	return nil
}

func (r *RegisterClient) Deregister() error {
	err := global.ConsulClient.Agent().ServiceDeregister(global.Cfg.ServerInfo.Id)
	if err != nil {
		zap.S().Panicw("[InitConsul] deregister service fail", err)
	}
	return nil
}
