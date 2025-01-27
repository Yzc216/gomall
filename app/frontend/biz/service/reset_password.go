package service

import (
	"context"
	"errors"
	user "github.com/Yzc216/gomall/app/frontend/hertz_gen/frontend/user"
	"github.com/Yzc216/gomall/app/frontend/infra/rpc"
	frontendUtils "github.com/Yzc216/gomall/app/frontend/utils"
	rpcuser "github.com/Yzc216/gomall/rpc_gen/kitex_gen/user"
	"github.com/cloudwego/hertz/pkg/app"
)

type ResetPasswordService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewResetPasswordService(Context context.Context, RequestContext *app.RequestContext) *ResetPasswordService {
	return &ResetPasswordService{RequestContext: RequestContext, Context: Context}
}

func (h *ResetPasswordService) Run(req *user.ResetPasswordReq) (resp bool, err error) {

	if req.ConfirmPassword != req.NewPassword {
		return false, errors.New("两次输入密码不一致")
	}
	res, err := rpc.UserClient.ResetPassword(h.Context, &rpcuser.ResetPasswordReq{
		UserId:      frontendUtils.GetUserIdFromCtx(h.Context),
		Password:    req.CurrentPassword,
		NewPassword: req.NewPassword,
	})
	if err != nil {
		return false, err
	}
	return res.IsReset, nil
}
