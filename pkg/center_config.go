package pkg

import (
	"dubbo.apache.org/dubbo-go/v3"
	"dubbo.apache.org/dubbo-go/v3/client"
	"dubbo.apache.org/dubbo-go/v3/config_center"
	"dubbo.apache.org/dubbo-go/v3/server"
)

// NacosConfig Nacos 配置
type NacosConfig struct {
	Address  string `yaml:"address" json:"address"`
	Username string `yaml:"username" json:"username"`
	Password string `yaml:"password" json:"password"`
}

// CenterConfig 配置中心配置
type CenterConfig struct {
	Namespace string `yaml:"namespace" json:"namespace"`
	Group     string `yaml:"group" json:"group"`
	DataID    string `yaml:"data_id" json:"data_id"`
}

// GetConfigCenterOption 返回配置中心选项
func GetConfigCenterOption(nacosConfig *NacosConfig, centerConfig *CenterConfig) dubbo.InstanceOption {
	return dubbo.WithConfigCenter(
		config_center.WithNacos(),
		config_center.WithAddress(nacosConfig.Address),
		config_center.WithUsername(nacosConfig.Username),
		config_center.WithPassword(nacosConfig.Password),
		config_center.WithNamespace(centerConfig.Namespace),
		config_center.WithGroup(centerConfig.Group),
		config_center.WithDataID(centerConfig.DataID),
	)
}

// GetDubboInstance 获取 dubbo 实例
func GetDubboInstance(option dubbo.InstanceOption) *dubbo.Instance {
	instance, err := dubbo.NewInstance(
		option,
	)
	if err != nil {
		panic(err)
	}
	return instance
}

// GetServer 获取服务端
func GetServer(instance *dubbo.Instance) *server.Server {
	server, err := instance.NewServer()
	if err != nil {
		panic(err)
	}
	return server
}

// GetClient 获取客户端
func GetClient(instance *dubbo.Instance) *client.Client {
	client, err := instance.NewClient()
	if err != nil {
		panic(err)
	}
	return client
}
