package service

import (
	"context"
	"github.com/Yzc216/gomall/app/inventory/biz/dal"
	inventory "github.com/Yzc216/gomall/rpc_gen/kitex_gen/inventory"
	"github.com/joho/godotenv"
	"sync"
	"testing"
)

func TestReleaseStock_Run(t *testing.T) {
	godotenv.Load("../../.env")
	dal.Init()

	ctx := context.Background()

	// init req and assert value

	var items []*inventory.InventoryReq_Item
	items = append(items, &inventory.InventoryReq_Item{
		SkuId:    553509862131171326,
		Quantity: 1,
	})
	req := &inventory.InventoryReq{
		OrderId: "123",
		Items:   items,
		Force:   false,
	}
	wg := sync.WaitGroup{}
	for i := 0; i < 90; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			s := NewReleaseStockService(ctx)
			resp, err := s.Run(req)
			t.Logf("err: %v", err)
			t.Logf("resp: %v", resp)
		}()
	}
	wg.Wait()

	// todo: edit your unit test

}
