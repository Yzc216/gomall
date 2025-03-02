package service

import (
	"context"
	"github.com/Yzc216/gomall/app/product/biz/dal"
	product "github.com/Yzc216/gomall/rpc_gen/kitex_gen/product"
	"github.com/joho/godotenv"
	"testing"
)

func TestListProducts_Run(t *testing.T) {
	godotenv.Load("../../.env")
	dal.Init()
	ctx := context.Background()
	s := NewListProductsService(ctx)
	// init req and assert value

	req := &product.ListProductsReq{
		Filter: &product.ProductFilter{
			CategoryId: 6,
			Pagination: &product.Pagination{
				Page:     1,
				PageSize: 20,
			},
		},
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
