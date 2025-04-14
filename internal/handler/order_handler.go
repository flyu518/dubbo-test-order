package handler

import (
	"context"

	"order/internal/service"

	"github.com/dubbogo/gost/log/logger"
	"github.com/flyu518/dubbo-test-sdk/order/api"
	userApi "github.com/flyu518/dubbo-test-sdk/user/api"
)

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
