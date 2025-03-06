package service

import (
	"context"
	"github.com/Yzc216/gomall/rpc_gen/kitex_gen/order"
	"testing"
)

func TestUpdateOrderState_Run(t *testing.T) {
	ctx := context.Background()
	s := NewUpdateOrderStateService(ctx)
	// init req and assert value

	req := &order.UpdateOrderStateReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
