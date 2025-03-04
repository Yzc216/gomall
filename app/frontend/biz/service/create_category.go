package service

import (
	"context"
	"fmt"
	"github.com/Yzc216/gomall/app/frontend/infra/rpc"
	"github.com/Yzc216/gomall/rpc_gen/kitex_gen/product"

	category "github.com/Yzc216/gomall/app/frontend/hertz_gen/frontend/category"
	"github.com/cloudwego/hertz/pkg/app"
)

type CreateCategoryService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCreateCategoryService(Context context.Context, RequestContext *app.RequestContext) *CreateCategoryService {
	return &CreateCategoryService{RequestContext: RequestContext, Context: Context}
}

func (h *CreateCategoryService) Run(req *category.CreateCategoryReq) (resp *category.CreateCategoryResp, err error) {
	if req.Name == "" {
		return nil, fmt.Errorf("category name is required")
	}
	if req.Description == "" {
		return nil, fmt.Errorf("category description is required")
	}
	if req.ParentId < 0 {
		return nil, fmt.Errorf("category parent id is invalid")
	}

	c := &product.CreateCategoryReq{
		Name:        req.Name,
		Description: req.Description,
		ParentId:    req.ParentId,
		ImageUrl:    req.ImageUrl,
	}
	Category, err := rpc.ProductClient.CreateCategory(h.Context, c)
	if err != nil {
		return nil, err
	}
	if Category == nil {
		return nil, fmt.Errorf("create category fail")
	}
	return
}
