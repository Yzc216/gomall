package service

import (
	"context"
	"github.com/Yzc216/gomall/app/product/biz/dal"
	product "github.com/Yzc216/gomall/rpc_gen/kitex_gen/product"
	"github.com/joho/godotenv"
	"testing"
)

func TestGetProduct_Run(t *testing.T) {
	godotenv.Load("../../.env")
	dal.Init()
	ctx := context.Background()
	s := NewGetProductService(ctx)
	// init req and assert value

	req := &product.GetProductReq{
		Id: 555939261627564030,
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
