package config

import (
	"errors"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// MysqlConfig Mysql配置结构
type MysqlConfig struct {
	GeneralDB `yaml:",inline" mapstructure:",squash"`
}

// Dsn 构建MySQL连接字符串
func (c *MysqlConfig) Dsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s",
		c.Username,
		c.Password,
		c.Host,
		c.Port,
		c.DbName,
		c.Config,
	)
}

// GetMysql 获取MySQL连接（懒加载）
func GetMysql(config *MysqlConfig) (*gorm.DB, error) {
	if config.DbName == "" {
		return nil, errors.New("数据库名不能为空")
	}

	mysqlConfig := mysql.Config{
		DSN:                       config.Dsn(), // DSN data source name
		DefaultStringSize:         191,          // string 类型字段的默认长度
		SkipInitializeWithVersion: false,        // 根据版本自动配置
	}

	gormConfig := &gorm.Config{
		Logger: logger.Default.LogMode(config.LogLevel()),
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   config.Prefix,
			SingularTable: config.Singular,
		},
	}

	db, err := gorm.Open(mysql.New(mysqlConfig), gormConfig)
	if err != nil {
		return nil, err
	}

	db.InstanceSet("gorm:table_options", "ENGINE="+config.Engine)
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(config.MaxIdleConns)
	sqlDB.SetMaxOpenConns(config.MaxOpenConns)
	return db, nil
}
