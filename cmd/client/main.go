package main

import (
	"context"
	"order/pkg"

	_ "dubbo.apache.org/dubbo-go/v3/imports" // 导入dubbo-go的依赖，必须的
	"github.com/dubbogo/gost/log/logger"
	"github.com/flyu518/dubbo-test-sdk/order/api"
)

func main() {

	// api.SetConsumerOrderService(srv)
	// if err := dubbo.Load(); err != nil {
	// 	panic(err)
	// }

	// logger.Infof("订单客户端已启动")

	// res, err := srv.GetOrder(context.Background(), &api.GetOrderRequest{
	// 	OrderId: "123456",
	// })
	// if err != nil {
	// 	panic(err)
	// }

	// logger.Infof("订单结果: %v", res)

	nacosConfig := pkg.NacosConfig{
		Address:  "127.0.0.1:8848",
		Username: "",
		Password: "",
	}
	centerConfig := pkg.CenterConfig{
		Namespace: "dev",
		Group:     "order",
		DataID:    "client.yaml",
	}

	// 获取 dubbo 实例和客户端
	instance := pkg.GetDubboInstance(pkg.GetConfigCenterOption(&nacosConfig, &centerConfig))
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
