package main

import (
	"order/internal/handler"
	"order/pkg/global"
	"order/pkg/util"

	_ "dubbo.apache.org/dubbo-go/v3/imports" // 导入dubbo-go的依赖，必须的
	"github.com/dubbogo/gost/log/logger"
	"github.com/flyu518/dubbo-test-sdk/order/api"
)

// 启动应用
func main() {
	// 初始化全局变量
	global.InitGlobal("./config/config.yaml") // 考虑设置个绝对地址

	// 获取 dubbo 实例和服务端
	instance := util.GetDubboInstance(global.ConfigCenterConfig)
	srv := util.GetDubboServer(instance)

	// 注册服务
	if err := api.RegisterOrderServiceHandler(srv, handler.GetOrderHandler(instance)); err != nil {
		panic(err)
	}

	global.Log = logger.GetLogger() // 实例化之后设置，不要在实例化之前设置

	global.Log.Info("订单服务已启动")

	// 启动服务
	if err := srv.Serve(); err != nil {
		global.Log.Error(err)
	}
}
