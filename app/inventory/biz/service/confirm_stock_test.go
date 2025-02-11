package service

import (
	"context"
	inventory "github.com/Yzc216/gomall/rpc_gen/kitex_gen/inventory"
	"testing"
)

func TestConfirmStock_Run(t *testing.T) {
	ctx := context.Background()
	s := NewConfirmStockService(ctx)
	// init req and assert value

	req := &inventory.InventoryReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
