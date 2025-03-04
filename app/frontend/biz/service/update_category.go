package service

import (
	"context"

	category "github.com/Yzc216/gomall/app/frontend/hertz_gen/frontend/category"
	"github.com/cloudwego/hertz/pkg/app"
)

type UpdateCategoryService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewUpdateCategoryService(Context context.Context, RequestContext *app.RequestContext) *UpdateCategoryService {
	return &UpdateCategoryService{RequestContext: RequestContext, Context: Context}
}

func (h *UpdateCategoryService) Run(req *category.UpdateCategoryReq) (resp *category.UpdateCategoryResp, err error) {
	return
}
