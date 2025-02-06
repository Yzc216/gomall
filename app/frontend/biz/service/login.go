package service

import (
	"context"
	auth "github.com/Yzc216/gomall/app/frontend/hertz_gen/frontend/auth"
	"github.com/Yzc216/gomall/app/frontend/infra/rpc"
	"github.com/Yzc216/gomall/rpc_gen/kitex_gen/user"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/sessions"
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

	//session存UserId
	session := sessions.Default(h.RequestContext)
	session.Set("user_id", uint64(resp.UserId))
	err = session.Save()
	if err != nil {
		return "", err
	}

	//// 生成JWT Token
	//token, err := utils.GenerateJWT(resp.UserId, resp.Role) // 假设RPC返回Role
	//if err != nil {
	//	return "", err
	//}
	//
	//// 在生成JWT后设置Cookie
	//h.RequestContext.SetCookie(
	//	"token",                        // name
	//	token,                          // value
	//	3600,                           // maxAge (秒)
	//	"/",                            // path
	//	"",                             // domain
	//	protocol.CookieSameSiteLaxMode, // sameSite
	//	false,                          // secure (根据实际HTTPS配置调整)
	//	true,                           // httpOnly
	//)

	redirect = "/"
	if req.Next != "" {
		redirect = req.Next
	}
	return
}
