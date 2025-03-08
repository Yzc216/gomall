package service

import (
	"context"
	common "github.com/Yzc216/gomall/app/frontend/hertz_gen/frontend/common"
	"github.com/Yzc216/gomall/app/frontend/infra/rpc"
	"github.com/Yzc216/gomall/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

type ListCategoryService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewListCategoryService(Context context.Context, RequestContext *app.RequestContext) *ListCategoryService {
	return &ListCategoryService{RequestContext: RequestContext, Context: Context}
}

func (h *ListCategoryService) Run(req *common.Empty) (resp map[string]any, err error) {
	tree, err := rpc.ProductClient.GetCategoryTree(h.Context, &product.GetCategoryTreeReq{IncludeSpuCount: true})
	if err != nil {
		return nil, err
	}

	products, err := rpc.ProductClient.ListProducts(h.Context, &product.ListProductsReq{})
	if err != nil {
		return nil, err
	}

	return utils.H{
		"tree":  tree.Tree,
		"items": products.Products,
	}, nil
}
