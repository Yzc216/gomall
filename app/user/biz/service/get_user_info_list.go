package service

import (
	"context"
	"github.com/Yzc216/gomall/app/user/biz/dal/mysql"
	"github.com/Yzc216/gomall/app/user/biz/model"
	user "github.com/Yzc216/gomall/app/user/kitex_gen/user"
)

type GetUserInfoListService struct {
	ctx context.Context
} // NewGetUserInfoListService new GetUserInfoListService
func NewGetUserInfoListService(ctx context.Context) *GetUserInfoListService {
	return &GetUserInfoListService{ctx: ctx}
}

// Run create note info
func (s *GetUserInfoListService) Run(req *user.GetUserInfoListReq) (resp *user.GetUserInfoListResp, err error) {

	users, err := model.GetBatchById(s.ctx, mysql.DB, int(req.Page), int(req.PageSize), req.UserIds)
	if err != nil {
		return nil, err
	}

	var userInfos []*user.User
	for _, u := range users {
		userInfos = append(userInfos, &user.User{
			Username: u.Username,
			Phone:    u.Phone,
			Email:    u.Email,
			Avatar:   u.Avatar,
			Role:     u.GetUint32Auth(),
		})
	}

	return &user.GetUserInfoListResp{
		UserInfos: userInfos,
		Page:      req.Page,
		PageSize:  req.PageSize,
		Total:     int32(len(userInfos)),
	}, nil
}
