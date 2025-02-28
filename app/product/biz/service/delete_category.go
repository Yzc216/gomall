package service

import (
	"context"
	product "github.com/Yzc216/gomall/rpc_gen/kitex_gen/product"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type DeleteCategoryService struct {
	ctx context.Context
} // NewDeleteCategoryService new DeleteCategoryService
func NewDeleteCategoryService(ctx context.Context) *DeleteCategoryService {
	return &DeleteCategoryService{ctx: ctx}
}

// Run create note info
func (s *DeleteCategoryService) Run(req *product.DeleteCategoryReq) (resp *emptypb.Empty, err error) {
	// Finish your business logic.

	return
}
