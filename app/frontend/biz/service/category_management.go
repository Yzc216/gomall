package service

import (
	"context"
	"github.com/Yzc216/gomall/app/frontend/infra/rpc"
	"github.com/Yzc216/gomall/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/common/utils"

	common "github.com/Yzc216/gomall/app/frontend/hertz_gen/frontend/common"
	"github.com/cloudwego/hertz/pkg/app"
)

type CategoryManagementService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCategoryManagementService(Context context.Context, RequestContext *app.RequestContext) *CategoryManagementService {
	return &CategoryManagementService{RequestContext: RequestContext, Context: Context}
}

func (h *CategoryManagementService) Run(req *common.Empty) (resp map[string]any, err error) {
	tree, err := rpc.ProductClient.GetCategoryTree(h.Context, &product.GetCategoryTreeReq{IncludeSpuCount: true})
	if err != nil {
		return nil, err
	}

	return utils.H{
		"tree": tree.Tree,
	}, nil
}
