package service

import (
	"context"
	product "github.com/Yzc216/gomall/rpc_gen/kitex_gen/product"
	"testing"
)

func TestUpdateCategory_Run(t *testing.T) {
	ctx := context.Background()
	s := NewUpdateCategoryService(ctx)
	// init req and assert value

	req := &product.UpdateCategoryReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
