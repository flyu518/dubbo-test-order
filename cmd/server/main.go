package main

import (
	"os"
	"os/signal"
	"syscall"

	"user/api"
	"user/internal/config"
	"user/internal/service"

	"dubbo.apache.org/dubbo-go/v3"
	_ "dubbo.apache.org/dubbo-go/v3/imports" // 导入dubbo-go的依赖，必须的
	"github.com/dubbogo/gost/log/logger"
)

// 启动应用
func main() {
	api.SetProviderUserService(new(service.UserService))
	if err := dubbo.Load(); err != nil {
		panic(err)
	}

	config.InitConfig("config/config.yaml") // 地址待处理

	logger.Infof("用户服务已启动")

	waitForShutdown()
}

// 优雅退出处理
func waitForShutdown() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	sig := <-quit
	logger.Infof("收到信号: %s, 退出应用", sig)
}
