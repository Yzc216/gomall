package service

import (
	"context"
	"fmt"
	"github.com/Yzc216/gomall/app/frontend/hertz_gen/frontend/user"
	"github.com/Yzc216/gomall/app/frontend/infra/rpc"
	frontendUtils "github.com/Yzc216/gomall/app/frontend/utils"
	rpcuser "github.com/Yzc216/gomall/rpc_gen/kitex_gen/user"
	"github.com/cloudwego/hertz/pkg/app"
)

type UpdateUserService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewUpdateUserService(Context context.Context, RequestContext *app.RequestContext) *UpdateUserService {
	return &UpdateUserService{RequestContext: RequestContext, Context: Context}
}

func (h *UpdateUserService) Run(req *user.UpdateUserReq) (resp bool, err error) {

	fmt.Println(req.Avatar)
	res, err := rpc.UserClient.UpdateUserInfo(h.Context, &rpcuser.UpdateUserInfoReq{
		UserId: frontendUtils.GetUserIdFromCtx(h.Context),
		UserInfo: &rpcuser.User{
			Username: req.Username,
			Avatar:   req.Avatar,
			Phone:    req.Phone,
			Email:    req.Email,
		},
	})
	if err != nil {
		return false, err
	}

	return res.IsUpdated, nil
}
