package service

import (
	"context"

	auth "github.com/Yzc216/gomall/app/frontend/hertz_gen/frontend/auth"
	"github.com/Yzc216/gomall/app/frontend/infra/rpc"
	"github.com/Yzc216/gomall/rpc_gen/kitex_gen/user"
	"github.com/cloudwego/hertz/pkg/app"
)

type LoginService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewLoginService(Context context.Context, RequestContext *app.RequestContext) *LoginService {
	return &LoginService{RequestContext: RequestContext, Context: Context}
}

func (h *LoginService) Run(req *auth.LoginReq) (redirect string, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()

	// user服务
	resp, err := rpc.UserClient.Login(h.Context, &user.LoginReq{
		LoginInfo: req.LoginInfo,
		Password:  req.Password,
	})
	if err != nil {
		return "", err
	}
	h.RequestContext.Set("role", resp.Role)

	redirect = "/"
	if req.Next != "" {
		redirect = req.Next
	}
	return
}
