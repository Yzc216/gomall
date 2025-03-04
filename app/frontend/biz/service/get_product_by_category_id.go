package service

import (
	"context"
	"github.com/Yzc216/gomall/app/frontend/infra/rpc"
	"github.com/Yzc216/gomall/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/common/utils"

	category "github.com/Yzc216/gomall/app/frontend/hertz_gen/frontend/category"
	"github.com/cloudwego/hertz/pkg/app"
)

type GetProductByCategoryIDService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetProductByCategoryIDService(Context context.Context, RequestContext *app.RequestContext) *GetProductByCategoryIDService {
	return &GetProductByCategoryIDService{RequestContext: RequestContext, Context: Context}
}

func (h *GetProductByCategoryIDService) Run(req *category.GetProductByCategoryIDReq) (resp map[string]any, err error) {
	products, err := rpc.ProductClient.ListProducts(h.Context, &product.ListProductsReq{
		Filter: &product.ProductFilter{
			CategoryId: req.CategoryId,
		},
	})
	if err != nil {
		return nil, err
	}
	return utils.H{
		"title":    "Category",
		"category": req.CategoryName,
		"items":    products.Products,
	}, nil
}
