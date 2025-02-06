package service

import (
	"context"
	"github.com/Yzc216/gomall/app/frontend/infra/rpc"
	"github.com/Yzc216/gomall/rpc_gen/kitex_gen/user"
	"github.com/cloudwego/hertz/pkg/common/utils"

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

func (h *AdminService) Run(req *common.Empty) (resp map[string]any, err error) {
	res, err := rpc.UserClient.GetUserInfoList(h.Context, &user.GetUserInfoListReq{})
	if err != nil {
		return nil, err
	}

	var users []map[string]any
	for _, item := range res.UserInfos {
		users = append(users, map[string]any{
			"Username": item.Username,
			"Email":    item.Email,
			"Phone":    item.Phone,
			"Role":     item.Role,
		})
	}
	return utils.H{
		"Users": users,
	}, nil
}
