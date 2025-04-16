package config

import (
	"encoding/json"
	"fmt"
	"os"
)

// ConfigCenterConfig 配置中心配置
type ConfigCenterConfig struct {
	Address   string `yaml:"address" json:"address"`
	Username  string `yaml:"username" json:"username"`
	Password  string `yaml:"password" json:"password"`
	Namespace string `yaml:"namespace" json:"namespace"`
	Group     string `yaml:"group" json:"group"`
	DataID    string `yaml:"data_id" json:"data_id"`
}

const CONFIG_CENTER_CONFIG_ENV_KEY = "LM_CONFIG_CENTER_CONFIG"

// GetEnvConfigCenterConfig 获取环境变量里面的配置中心配置
func GetEnvConfigCenterConfig() *ConfigCenterConfig {
	center_config := os.Getenv(CONFIG_CENTER_CONFIG_ENV_KEY)
	if center_config == "" {
		panic(fmt.Sprintf("环境变量 %s 未设置", CONFIG_CENTER_CONFIG_ENV_KEY))
	}

	configCenterConfig := ConfigCenterConfig{}
	err := json.Unmarshal([]byte(center_config), &configCenterConfig)
	if err != nil {
		panic(fmt.Sprintf("环境变量 %s 解析失败: %s", CONFIG_CENTER_CONFIG_ENV_KEY, err.Error()))
	}

	return &configCenterConfig
}
