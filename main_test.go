package main

import (
	"context"
	"order/pkg/global"
	"order/pkg/util"
	"os"
	"testing"

	"dubbo.apache.org/dubbo-go/v3"
	"dubbo.apache.org/dubbo-go/v3/client"
	_ "dubbo.apache.org/dubbo-go/v3/imports" // 导入dubbo-go的依赖，必须的
	"github.com/flyu518/dubbo-test-sdk/order/api"
	"github.com/stretchr/testify/assert"
)

var instance *dubbo.Instance
var cli *client.Client
var srv api.OrderService

func TestMain(m *testing.M) {
	// 初始化全局变量
	global.InitGlobal("./config/config.yaml")

	// 获取 dubbo 实例和服务端
	instance = util.GetDubboInstance(global.ConfigCenterConfig)
	cli = util.GetDubboClient(instance)

	// 获取服务
	var err error
	srv, err = api.NewOrderService(cli)
	if err != nil {
		panic(err)
	}

	// 调用 m.Run 执行测试
	code := m.Run()

	os.Exit(code)
}

func TestOrder(t *testing.T) {
	t.Run("获取订单", func(t *testing.T) {
		res, err := srv.GetOrder(context.Background(), &api.GetOrderRequest{
			OrderId: "123456",
		})

		assert.NoError(t, err)

		assert.Equal(t, "123456", res.Order.OrderId)
	})
}
