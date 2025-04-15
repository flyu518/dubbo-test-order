package main

import (
	"order/internal/handler"
	"order/pkg"

	_ "dubbo.apache.org/dubbo-go/v3/imports" // 导入dubbo-go的依赖，必须的
	"github.com/dubbogo/gost/log/logger"
	"github.com/flyu518/dubbo-test-sdk/order/api"
	userApi "github.com/flyu518/dubbo-test-sdk/user/api"
)

// 启动应用
func main() {
	// 获取 dubbo 实例和服务端
	instance := pkg.GetDubboInstance(pkg.ParseEnvCenterConfig())
	srv := pkg.GetServer(instance)

	// 获取用户服务
	userClient := pkg.GetClient(instance)
	userService, err := userApi.NewUserService(userClient)
	if err != nil {
		panic(err)
	}

	// 注册服务
	if err := api.RegisterOrderServiceHandler(srv, &handler.OrderHandler{
		UserService: userService,
	}); err != nil {
		panic(err)
	}

	logger.Infof("订单服务已启动")

	// 启动服务
	if err := srv.Serve(); err != nil {
		logger.Error(err)
	}
}
