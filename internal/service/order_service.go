/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package service

import (
	"github.com/dubbogo/gost/log/logger"
	"github.com/flyu518/dubbo-test-sdk/order/api"
	userApi "github.com/flyu518/dubbo-test-sdk/user/api"
)

var OrderService *orderService

// orderService 实现订单服务
type orderService struct {
}

// GetOrder 实现获取订单服务
func (s *orderService) GetOrder(req *api.GetOrderRequest, userService *userApi.GetUserResponse) (*api.GetOrderResponse, error) {

	logger.Infof("收到获取订单请求: %v", req.OrderId)

	logger.Infof("查询用户id: %v", userService.User.Username)

	// 实际应用中，这里应该有真正的业务逻辑实现
	// 这里简单返回成功
	return &api.GetOrderResponse{
		Order: &api.Order{
			OrderId:    req.OrderId,
			OrderName:  "订单名称",
			OrderPrice: "100",
		},
	}, nil
}
