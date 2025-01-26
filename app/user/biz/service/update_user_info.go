package service

import (
	"context"
	"github.com/Yzc216/gomall/app/user/biz/dal/mysql"
	"github.com/Yzc216/gomall/app/user/biz/model"
	user "github.com/Yzc216/gomall/app/user/kitex_gen/user"
)

type UpdateUserInfoService struct {
	ctx context.Context
}

// NewUpdateUserInfoService new UpdateUserInfoService
func NewUpdateUserInfoService(ctx context.Context) *UpdateUserInfoService {
	return &UpdateUserInfoService{ctx: ctx}
}

// Run create note info
func (s *UpdateUserInfoService) Run(req *user.UpdateUserInfoReq) (resp *user.UpdateUserInfoResp, err error) {

	u := &model.User{
		Avatar:   req.UserInfo.Avatar,
		Phone:    req.UserInfo.Phone,
		Email:    req.UserInfo.Email,
		Username: req.UserInfo.Username,
	}
	err = model.UpdateUser(s.ctx, mysql.DB, u)
	if err != nil {
		return nil, err
	}

	return &user.UpdateUserInfoResp{IsUpdated: true}, nil
}
