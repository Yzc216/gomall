package service

import (
	"context"
	"github.com/Yzc216/gomall/app/order/biz/dal"
	"github.com/Yzc216/gomall/app/order/infra/rpc"
	"github.com/Yzc216/gomall/rpc_gen/kitex_gen/order"
	"github.com/joho/godotenv"
	"testing"
)

func TestUpdateOrderState_Run(t *testing.T) {
	godotenv.Load("../../.env")
	dal.Init()
	rpc.InitClient()

	ctx := context.Background()
	s := NewUpdateOrderStateService(ctx)
	// init req and assert value

	req := &order.UpdateOrderStateReq{
		UserId:  551017802534813694,
		OrderId: 557585680365060094,
		State:   order.OrderState_OrderStateCanceled,
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
