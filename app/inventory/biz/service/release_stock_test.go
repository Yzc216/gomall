package service

import (
	"context"
	"github.com/Yzc216/gomall/app/inventory/biz/dal/mysql"
	inventory "github.com/Yzc216/gomall/rpc_gen/kitex_gen/inventory"
	"github.com/joho/godotenv"
	"testing"
)

func TestReleaseStock_Run(t *testing.T) {
	godotenv.Load("../../.env")
	mysql.Init()

	ctx := context.Background()
	s := NewReleaseStockService(ctx)
	// init req and assert value

	var items []*inventory.InventoryReq_Item
	items = append(items, &inventory.InventoryReq_Item{
		SkuId:    553509862131171326,
		Quantity: 10,
	})
	req := &inventory.InventoryReq{
		OrderId: "123",
		Items:   items,
		Force:   false,
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
