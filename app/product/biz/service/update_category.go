package service

import (
	"context"
	product "github.com/Yzc216/gomall/rpc_gen/kitex_gen/product"
)

type UpdateCategoryService struct {
	ctx context.Context
} // NewUpdateCategoryService new UpdateCategoryService
func NewUpdateCategoryService(ctx context.Context) *UpdateCategoryService {
	return &UpdateCategoryService{ctx: ctx}
}

// Run create note info
func (s *UpdateCategoryService) Run(req *product.UpdateCategoryReq) (resp *product.Category, err error) {
	// Finish your business logic.

	return
}
