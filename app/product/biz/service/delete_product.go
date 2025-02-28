package service

import (
	"context"
	product "github.com/Yzc216/gomall/rpc_gen/kitex_gen/product"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type DeleteProductService struct {
	ctx context.Context
} // NewDeleteProductService new DeleteProductService
func NewDeleteProductService(ctx context.Context) *DeleteProductService {
	return &DeleteProductService{ctx: ctx}
}

// Run create note info
func (s *DeleteProductService) Run(req *product.DeleteProductReq) (resp *emptypb.Empty, err error) {
	// Finish your business logic.

	return
}
