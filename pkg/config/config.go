package config

import (
	"errors"
	"fmt"
	"sync"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var (
	viperOnce sync.Once
)

type Config struct {
	System SystemConfig `mapstructure:"system" yaml:"system" json:"system"`
	MySQL  MysqlConfig  `mapstructure:"mysql" yaml:"mysql" json:"mysql"`
	Redis  RedisConfig  `mapstructure:"redis" yaml:"redis" json:"redis"`
}

// GetConfig 获取配置
// TODO::暂时使用配置文件方式，配置中心有点问题
func GetConfig(configPath string, config *Config) error {
	var innerErr error
	viperOnce.Do(func() {
		v := viper.New()
		v.SetConfigFile(configPath)
		v.SetConfigType("yaml")
		err := v.ReadInConfig()
		if err != nil {
			innerErr = errors.New(fmt.Sprintf("读取配置文件失败: %s", err))
			return
		}

		// 监控配置文件变化
		v.WatchConfig()

		v.OnConfigChange(func(e fsnotify.Event) {
			fmt.Println("config file changed:", e.Name)
			if err = v.Unmarshal(config); err != nil {
				fmt.Println(err)
			}
		})

		// 解析配置
		if err := v.Unmarshal(config); err != nil {
			innerErr = fmt.Errorf("解析配置文件失败: %s", err)
			return
		}
	})

	return innerErr
}
