package util

import (
	"order/pkg/config"

	"dubbo.apache.org/dubbo-go/v3"
	"dubbo.apache.org/dubbo-go/v3/client"
	"dubbo.apache.org/dubbo-go/v3/config_center"
	"dubbo.apache.org/dubbo-go/v3/server"
)

// GetDubboInstance 获取 dubbo 实例
func GetDubboInstance(centerConfig *config.ConfigCenterConfig) *dubbo.Instance {
	instance, err := dubbo.NewInstance(
		dubbo.WithConfigCenter(
			config_center.WithNacos(),
			config_center.WithAddress(centerConfig.Address),
			config_center.WithUsername(centerConfig.Username),
			config_center.WithPassword(centerConfig.Password),
			config_center.WithNamespace(centerConfig.Namespace),
			config_center.WithGroup(centerConfig.Group),
			config_center.WithDataID(centerConfig.DataID),
		),
	)
	if err != nil {
		panic(err)
	}
	return instance
}

// GetDubboServer 获取 dubbo 服务端
func GetDubboServer(instance *dubbo.Instance) *server.Server {
	server, err := instance.NewServer()
	if err != nil {
		panic(err)
	}
	return server
}

// GetDubboClient 获取 dubbo 客户端
func GetDubboClient(instance *dubbo.Instance) *client.Client {
	client, err := instance.NewClient()
	if err != nil {
		panic(err)
	}
	return client
}
