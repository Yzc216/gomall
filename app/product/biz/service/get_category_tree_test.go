package service

import (
	"context"
	"fmt"
	"github.com/Yzc216/gomall/app/product/biz/dal"
	product "github.com/Yzc216/gomall/rpc_gen/kitex_gen/product"
	"github.com/joho/godotenv"
	"testing"
)

func TestGetCategoryTree_Run(t *testing.T) {
	godotenv.Load("../../.env")
	dal.Init()
	ctx := context.Background()
	s := NewGetCategoryTreeService(ctx)
	// init req and assert value

	req := &product.GetCategoryTreeReq{IncludeSpuCount: false}
	resp, err := s.Run(req)

	for _, v := range resp.Tree {
		fmt.Println(v)
	}
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
