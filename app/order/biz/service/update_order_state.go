package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/Yzc216/gomall/app/order/biz/dal/mysql"
	"github.com/Yzc216/gomall/app/order/biz/model"
	"github.com/Yzc216/gomall/app/order/infra/rpc"
	"github.com/Yzc216/gomall/rpc_gen/kitex_gen/inventory"
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
	//参数校验
	if req.UserId == 0 || req.OrderId == 0 {
		err = fmt.Errorf("user_id or order_id can not be empty")
		return
	}
	if req.GetState().Enum() == nil {
		err = fmt.Errorf("state can not be empty")
		return
	}

	//获取订单状态
	orderRes, err := model.GetOrder(s.ctx, mysql.DB, req.UserId, req.OrderId)
	if err != nil {
		klog.Errorf("model.GetOrder.err:%v", err)
		return nil, err
	}
	fmt.Println(orderRes)
	if orderRes.OrderState == model.OrderStateCancelled {
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
		state = model.OrderStateCancelled
	}

	//更新库存
	var invItems []*inventory.InventoryReq_Item
	items := orderRes.OrderItems
	for _, item := range items {
		invItems = append(invItems, &inventory.InventoryReq_Item{
			SkuId:    item.SkuId,
			Quantity: item.Quantity,
		})
	}
	fmt.Println(invItems)
	if state == model.OrderStatePaid {
		stock, err := rpc.InventoryClient.ConfirmStock(s.ctx, &inventory.InventoryReq{
			OrderId: req.OrderId,
			Items:   invItems,
			Force:   false,
		})
		fmt.Println(stock)
		if err != nil || !stock.Success {
			return nil, errors.New("confirm stock fail")
		}
	}

	if state == model.OrderStateCancelled {
		stock, err := rpc.InventoryClient.ReleaseStock(s.ctx, &inventory.InventoryReq{
			OrderId: req.OrderId,
			Items:   invItems,
			Force:   false,
		})
		if err != nil || !stock.Success {
			return nil, errors.New("release stock fail")
		}
	}

	//更新订单状态
	err = model.UpdateOrderState(s.ctx, mysql.DB, req.UserId, req.OrderId, state)
	if err != nil {
		klog.Errorf("model.ListOrder.err:%v", err)
		return nil, err
	}

	return &order.UpdateOrderStateResp{
		Success: true,
	}, nil
}
