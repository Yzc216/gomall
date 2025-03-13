package service

import (
	"context"
	"github.com/Yzc216/gomall/app/order/biz/dal/mysql"
	"github.com/Yzc216/gomall/app/order/infra/rpc"
	"github.com/Yzc216/gomall/rpc_gen/kitex_gen/cart"
	order "github.com/Yzc216/gomall/rpc_gen/kitex_gen/order"
	"github.com/joho/godotenv"
	"testing"
)

func TestPlaceOrder_Run(t *testing.T) {
	godotenv.Load("../../.env")
	mysql.Init()
	rpc.InitClient()

	ctx := context.Background()
	s := NewPlaceOrderService(ctx)
	// init req and assert value

	items := []*order.OrderItem{
		{
			Item: &cart.CartItem{
				SpuId:    555937596203663358,
				SkuId:    2,
				Quantity: 10,
			},
			Cost: 18.99,
		},
	}
	req := &order.PlaceOrderReq{
		UserId:       551017802534813694,
		UserCurrency: "",
		Address:      nil,
		Email:        "yangzc216@163.com",
		Items:        items,
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
