package service

import (
	"context"
	"github.com/Yzc216/gomall/app/product/biz/dal"
	"github.com/Yzc216/gomall/rpc_gen/kitex_gen/product"
	"github.com/joho/godotenv"
	"testing"
)

func TestBatchGetProducts_Run(t *testing.T) {
	godotenv.Load("../../.env")
	dal.Init()
	ctx := context.Background()
	s := NewBatchGetProductsService(ctx)
	// init req and assert value

	req := &product.BatchGetProductsReq{
		Ids: []uint64{555937596203663358, 555940418534047742, 2, 2, 3},
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

}
