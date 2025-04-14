package pkg

import (
	"encoding/json"
	"os"
)

// GetNacosConfig 获取 nacos 配置
func GetNacosConfig() *NacosConfig {
	nacos_config := os.Getenv("DUBBOGO_NACOS_CONFIG")
	if nacos_config == "" {
		panic("DUBBOGO_NACOS_CONFIG 未设置")
	}

	// 解析配置
	nacosConfig := NacosConfig{}
	err := json.Unmarshal([]byte(nacos_config), &nacosConfig)
	if err != nil {
		panic("DUBBOGO_NACOS_CONFIG 解析失败: " + err.Error())
	}

	return &nacosConfig
}

// GetCenterConfig 获取中心配置
func GetCenterConfig() *CenterConfig {
	center_config := os.Getenv("DUBBOGO_CENTER_CONFIG")
	if center_config == "" {
		panic("DUBBOGO_CENTER_CONFIG 未设置")
	}

	centerConfig := CenterConfig{}
	err := json.Unmarshal([]byte(center_config), &centerConfig)
	if err != nil {
		panic("DUBBOGO_CENTER_CONFIG 解析失败: " + err.Error())
	}

	return &centerConfig
}
