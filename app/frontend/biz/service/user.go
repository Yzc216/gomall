package service

import (
	"context"
	"github.com/Yzc216/gomall/app/frontend/infra/rpc"
	frontendUtils "github.com/Yzc216/gomall/app/frontend/utils"
	"github.com/Yzc216/gomall/rpc_gen/kitex_gen/user"
	"github.com/cloudwego/hertz/pkg/common/utils"

	common "github.com/Yzc216/gomall/app/frontend/hertz_gen/frontend/common"
	"github.com/cloudwego/hertz/pkg/app"
)

type UserService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewUserService(Context context.Context, RequestContext *app.RequestContext) *UserService {
	return &UserService{RequestContext: RequestContext, Context: Context}
}

func (h *UserService) Run(req *common.Empty) (resp map[string]any, err error) {
	u, err := rpc.UserClient.GetUserInfo(h.Context, &user.GetUserInfoReq{
		UserId: frontendUtils.GetUserIdFromCtx(h.Context),
	})
	if err != nil {
		return nil, err
	}
	return utils.H{
		"item": u.UserInfo,
	}, nil
}
