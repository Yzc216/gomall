package service

import (
	"context"
	auth "github.com/Yzc216/gomall/app/frontend/hertz_gen/frontend/auth"
	common "github.com/Yzc216/gomall/app/frontend/hertz_gen/frontend/common"
	"github.com/Yzc216/gomall/app/frontend/infra/rpc"
	"github.com/Yzc216/gomall/rpc_gen/kitex_gen/user"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/sessions"
)

type RegisterService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewRegisterService(Context context.Context, RequestContext *app.RequestContext) *RegisterService {
	return &RegisterService{RequestContext: RequestContext, Context: Context}
}

func (h *RegisterService) Run(req *auth.RegisterReq) (resp *common.Empty, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()

	// user服务
	userResp, err := rpc.UserClient.Register(h.Context, &user.RegisterReq{
		UserInfo: &user.User{
			Username: req.Username,
			Password: req.Password,
			Email:    req.Email,
			Phone:    req.Phone,
			Role:     []uint32{req.Role},
		},
		PasswordConfirm: req.PasswordConfirm,
	})
	if err != nil {
		return nil, err
	}

	session := sessions.Default(h.RequestContext)
	session.Set("user_id", uint64(userResp.UserId))
	err = session.Save()
	if err != nil {
		return nil, err
	}
	return
}
