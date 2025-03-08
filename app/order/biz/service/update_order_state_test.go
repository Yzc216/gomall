package service

import (
	"context"
	"github.com/Yzc216/gomall/app/order/biz/dal"
	"github.com/Yzc216/gomall/rpc_gen/kitex_gen/order"
	"github.com/joho/godotenv"
	"testing"
)

func TestUpdateOrderState_Run(t *testing.T) {
	godotenv.Load("../../.env")
	dal.Init()
	ctx := context.Background()
	s := NewUpdateOrderStateService(ctx)
	// init req and assert value

	req := &order.UpdateOrderStateReq{
		UserId:  551017802534813694,
		OrderId: 556716135110737918,
		State:   order.OrderState_OrderStatePaid,
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
