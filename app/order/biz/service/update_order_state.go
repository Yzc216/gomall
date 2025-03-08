package service

import (
	"context"
	"fmt"
	"github.com/Yzc216/gomall/app/order/biz/dal/mysql"
	"github.com/Yzc216/gomall/app/order/biz/model"
	order "github.com/Yzc216/gomall/rpc_gen/kitex_gen/order"
	"github.com/cloudwego/kitex/pkg/klog"
)

type UpdateOrderStateService struct {
	ctx context.Context
} // NewUpdateOrderStateService new UpdateOrderStateService
func NewUpdateOrderStateService(ctx context.Context) *UpdateOrderStateService {
	return &UpdateOrderStateService{ctx: ctx}
}

// Run create note info
func (s *UpdateOrderStateService) Run(req *order.UpdateOrderStateReq) (resp *order.UpdateOrderStateResp, err error) {
	if req.UserId == 0 || req.OrderId == 0 {
		err = fmt.Errorf("user_id or order_id can not be empty")
		return
	}
	if req.GetState().Enum() == nil {
		err = fmt.Errorf("state can not be empty")
		return
	}

	orderRes, err := model.GetOrder(s.ctx, mysql.DB, req.UserId, req.OrderId)
	if err != nil {
		klog.Errorf("model.GetOrder.err:%v", err)
		return nil, err
	}
	if orderRes.OrderState == model.OrderStateCanceled {
		return &order.UpdateOrderStateResp{
			Success: false,
			Error:   "订单已取消，无法修改状态",
		}, nil
	}
	var state model.OrderState
	switch req.State {
	case order.OrderState_OrderStatePaid:
		state = model.OrderStatePaid
	case order.OrderState_OrderStatePlaced:
		state = model.OrderStatePlaced
	case order.OrderState_OrderStateCanceled:
		state = model.OrderStateCanceled
	}
	err = model.UpdateOrderState(s.ctx, mysql.DB, req.UserId, req.OrderId, state)
	if err != nil {
		klog.Errorf("model.ListOrder.err:%v", err)
		return nil, err
	}

	return &order.UpdateOrderStateResp{
		Success: true,
	}, nil
}
