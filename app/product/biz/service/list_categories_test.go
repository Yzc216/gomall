package service

import (
	"context"
	product "github.com/Yzc216/gomall/rpc_gen/kitex_gen/product"
	"testing"
)

func TestListCategories_Run(t *testing.T) {
	ctx := context.Background()
	s := NewListCategoriesService(ctx)
	// init req and assert value

	req := &product.ListCategoriesReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
