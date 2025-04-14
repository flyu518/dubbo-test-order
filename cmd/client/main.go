package main

import (
	"context"
	"order/pkg"

	_ "dubbo.apache.org/dubbo-go/v3/imports" // 导入dubbo-go的依赖，必须的
	"github.com/dubbogo/gost/log/logger"
	"github.com/flyu518/dubbo-test-sdk/order/api"
)

func main() {
	// 通过环境变量获取配置
	nacosConfig := pkg.GetNacosConfig()
	centerConfig := pkg.GetCenterConfig()

	// 获取 dubbo 实例和客户端
	instance := pkg.GetDubboInstance(pkg.GetConfigCenterOption(nacosConfig, centerConfig))
	client := pkg.GetClient(instance)

	// 获取服务
	srv, err := api.NewOrderService(client)
	if err != nil {
		panic(err)
	}

	logger.Infof("订单客户端已启动")

	res, err := srv.GetOrder(context.Background(), &api.GetOrderRequest{
		OrderId: "123456",
	})
	if err != nil {
		panic(err)
	}

	logger.Infof("订单结果: %v", res)
}
