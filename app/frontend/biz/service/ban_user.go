package service

import (
	"context"

	common "github.com/Yzc216/gomall/app/frontend/hertz_gen/frontend/common"
	user "github.com/Yzc216/gomall/app/frontend/hertz_gen/frontend/user"
	"github.com/cloudwego/hertz/pkg/app"
)

type BanUserService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewBanUserService(Context context.Context, RequestContext *app.RequestContext) *BanUserService {
	return &BanUserService{RequestContext: RequestContext, Context: Context}
}

func (h *BanUserService) Run(req *user.BanUserReq) (resp *common.Empty, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
