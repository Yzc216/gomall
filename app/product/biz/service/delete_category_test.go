package service

import (
	"context"
	product "github.com/Yzc216/gomall/rpc_gen/kitex_gen/product"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

func TestDeleteCategory_Run(t *testing.T) {
	ctx := context.Background()
	s := NewDeleteCategoryService(ctx)
	// init req and assert value

	req := &product.DeleteCategoryReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
