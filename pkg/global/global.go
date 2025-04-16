package global

import (
	"order/internal/util"
	"order/pkg/config"

	"github.com/dubbogo/gost/log/logger"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	ConfigCenterConfig *config.ConfigCenterConfig

	Config *config.Config
	DB     *gorm.DB
	Redis  redis.UniversalClient

	Log logger.Logger // 这个地方特殊，要使用和框架相同的配置，需要在 main 中初始化
)

func InitGlobal(configPath string) {
	var err error

	ConfigCenterConfig = config.GetEnvConfigCenterConfig()

	Config = &config.Config{}
	err = config.GetConfig(configPath, Config)
	if err != nil {
		panic(err)
	}

	DB, err = config.GetMysql(&Config.MySQL)
	if err != nil {
		panic(err)
	}

	// 临时操作（正式环境不要这样）
	if Config.System.Env != config.EnvProd {
		if err := util.MigrateTables(DB); err != nil {
			panic(err)
		}
	}

	Redis, err = config.GetRedis(&Config.Redis)
	if err != nil {
		panic(err)
	}
}
