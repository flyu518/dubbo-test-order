package pkg

import (
	"encoding/json"
	"fmt"
	"os"

	"dubbo.apache.org/dubbo-go/v3"
	"dubbo.apache.org/dubbo-go/v3/client"
	"dubbo.apache.org/dubbo-go/v3/config_center"
	"dubbo.apache.org/dubbo-go/v3/server"
)

// CenterConfig 配置中心配置
type CenterConfig struct {
	Address   string `yaml:"address" json:"address"`
	Username  string `yaml:"username" json:"username"`
	Password  string `yaml:"password" json:"password"`
	Namespace string `yaml:"namespace" json:"namespace"`
	Group     string `yaml:"group" json:"group"`
	DataID    string `yaml:"data_id" json:"data_id"`
}

const CENTER_CONFIG_ENV_KEY = "LM_CENTER_CONFIG"

// ParseEnvCenterConfig 解析环境变量里面的配置中心配置
func ParseEnvCenterConfig() *CenterConfig {
	center_config := os.Getenv(CENTER_CONFIG_ENV_KEY)
	if center_config == "" {
		panic(fmt.Sprintf("环境变量 %s 未设置", CENTER_CONFIG_ENV_KEY))
	}

	centerConfig := CenterConfig{}
	err := json.Unmarshal([]byte(center_config), &centerConfig)
	if err != nil {
		panic(fmt.Sprintf("环境变量 %s 解析失败: %s", CENTER_CONFIG_ENV_KEY, err.Error()))
	}

	return &centerConfig
}

// GetDubboInstance 获取 dubbo 实例
func GetDubboInstance(centerConfig *CenterConfig) *dubbo.Instance {
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
