package config

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

// RedisConfig Redis配置结构
type RedisConfig struct {
	Mode     string       `mapstructure:"mode" yaml:"mode" json:"mode"`
	Password string       `mapstructure:"password" yaml:"password" json:"password"`
	Single   RedisSingle  `mapstructure:"single" yaml:"single" json:"single"`
	Cluster  RedisCluster `mapstructure:"cluster" yaml:"cluster" json:"cluster"`
}

// RedisSingle 单节点配置
type RedisSingle struct {
	Addr string `mapstructure:"addr" yaml:"addr" json:"addr"`
	DB   int    `mapstructure:"db" yaml:"db" json:"db"`
}

// RedisCluster 集群配置
type RedisCluster struct {
	Addrs []string `mapstructure:"addrs" yaml:"addrs" json:"addrs"`
}

// GetRedis 获取Redis客户端（懒加载）
func GetRedis(config *RedisConfig) (redis.UniversalClient, error) {
	var client redis.UniversalClient

	switch config.Mode {
	case "single":
		// 单节点模式
		client = redis.NewClient(&redis.Options{
			Addr:     config.Single.Addr,
			Password: config.Password,
			DB:       config.Single.DB,
		})
	case "cluster":
		// 集群模式
		client = redis.NewClusterClient(&redis.ClusterOptions{
			Addrs:    config.Cluster.Addrs,
			Password: config.Password,
		})
	default:
		return nil, fmt.Errorf("不支持的Redis模式: %s", config.Mode)
	}

	// 测试连接
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := client.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("Redis连接失败: %s", err)
	}

	return client, nil
}
