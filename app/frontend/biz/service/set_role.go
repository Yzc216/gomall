package service

import (
	"context"
	"github.com/Yzc216/gomall/app/frontend/hertz_gen/frontend/user"

	common "github.com/Yzc216/gomall/app/frontend/hertz_gen/frontend/common"
	"github.com/cloudwego/hertz/pkg/app"
)

type SetRoleService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewSetRoleService(Context context.Context, RequestContext *app.RequestContext) *SetRoleService {
	return &SetRoleService{RequestContext: RequestContext, Context: Context}
}

func (h *SetRoleService) Run(req *user.SetRoleReq) (resp *common.Empty, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
