package service

import (
	"context"
	"github.com/Yzc216/gomall/app/inventory/biz/dal"
	"github.com/Yzc216/gomall/rpc_gen/kitex_gen/inventory"
	"github.com/joho/godotenv"
	"sync"
	"testing"
)

func TestReserveStock_Run(t *testing.T) {
	godotenv.Load("../../.env")
	dal.Init()

	ctx := context.Background()

	//init req and assert value
	//in := &model.Inventory{
	//	SkuID:     553509862131171326,
	//	Total:     100,
	//	Available: 100,
	//	Locked:    0,
	//}
	//model.InitStock(ctx, mysql.DB, in)
	var items []*inventory.InventoryReq_Item
	items = append(items, &inventory.InventoryReq_Item{
		SkuId:    553509862131171326,
		Quantity: 1,
	})
	req := &inventory.InventoryReq{
		OrderId: "123",
		Items:   items,
	}

	wg := sync.WaitGroup{}
	for i := 0; i < 90; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			s := NewReserveStockService(ctx)
			resp, err := s.Run(req)
			t.Logf("err: %v", err)
			t.Logf("resp: %v", resp)
		}()
	}
	wg.Wait()

	// todo: edit your unit test

}
