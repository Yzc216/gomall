package service

import (
	"context"

	common "github.com/Yzc216/gomall/app/frontend/hertz_gen/frontend/common"
	"github.com/cloudwego/hertz/pkg/app"
)

type AdminService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewAdminService(Context context.Context, RequestContext *app.RequestContext) *AdminService {
	return &AdminService{RequestContext: RequestContext, Context: Context}
}

func (h *AdminService) Run(req *common.Empty) (resp *common.Empty, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
