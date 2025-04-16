package config

import (
	"strings"

	"gorm.io/gorm/logger"
)

type DsnProvider interface {
	Dsn() string
}

type GeneralDB struct {
	Host         string `mapstructure:"host" json:"host" yaml:"host"`                               // 数据库地址
	Port         string `mapstructure:"port" json:"port" yaml:"port"`                               // 数据库端口
	DbName       string `mapstructure:"db-name" json:"db-name" yaml:"db-name"`                      // 数据库名
	Username     string `mapstructure:"username" json:"username" yaml:"username"`                   // 数据库账号
	Password     string `mapstructure:"password" json:"password" yaml:"password"`                   // 数据库密码
	Prefix       string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`                         // 数据库前缀
	Config       string `mapstructure:"config" json:"config" yaml:"config"`                         // 高级配置
	Engine       string `mapstructure:"engine" json:"engine" yaml:"engine" default:"InnoDB"`        // 数据库引擎，默认InnoDB
	MaxIdleConns int    `mapstructure:"max-idle-conns" json:"max-idle-conns" yaml:"max-idle-conns"` // 空闲中的最大连接数
	MaxOpenConns int    `mapstructure:"max-open-conns" json:"max-open-conns" yaml:"max-open-conns"` // 打开到数据库的最大连接数
	Singular     bool   `mapstructure:"singular" json:"singular" yaml:"singular"`                   // 是否开启全局禁用复数，true表示开启
	LogMode      string `mapstructure:"log-mode" json:"log-mode" yaml:"log-mode"`                   // 是否开启Gorm全局日志
	LogZap       bool   `mapstructure:"log-zap" json:"log-zap" yaml:"log-zap"`                      // 是否通过zap写入日志文件
}

func (c GeneralDB) LogLevel() logger.LogLevel {
	switch strings.ToLower(c.LogMode) {
	case "silent", "Silent":
		return logger.Silent
	case "error", "Error":
		return logger.Error
	case "warn", "Warn":
		return logger.Warn
	case "info", "Info":
		return logger.Info
	default:
		return logger.Info
	}
}
