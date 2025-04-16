package handler

import (
	"context"

	"order/internal/service"
	"order/pkg/util"

	"dubbo.apache.org/dubbo-go/v3"
	"github.com/dubbogo/gost/log/logger"
	"github.com/flyu518/dubbo-test-sdk/order/api"
	userApi "github.com/flyu518/dubbo-test-sdk/user/api"
)

func GetOrderHandler(instance *dubbo.Instance) *OrderHandler {
	// 获取用户服务
	client := util.GetDubboClient(instance)
	userService, err := userApi.NewUserService(client)
	if err != nil {
		panic(err)
	}
	return &OrderHandler{
		UserService: userService,
	}
}

// OrderHandler 实现订单服务
type OrderHandler struct {
	UserService userApi.UserService
}

// GetOrder 实现获取订单服务
func (h *OrderHandler) GetOrder(ctx context.Context, req *api.GetOrderRequest) (*api.GetOrderResponse, error) {

	logger.Infof("收到获取订单请求: %v", req.OrderId)

	// 查询用户信息
	user, err := h.UserService.GetUser(ctx, &userApi.GetUserRequest{
		Username: "hahah",
	})
	if err != nil {
		return nil, err
	}

	logger.Infof("查询用户信息: %v", user)

	return service.OrderService.GetOrder(req, user)
}
