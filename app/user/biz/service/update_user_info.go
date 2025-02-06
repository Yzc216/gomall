package service

import (
	"context"
	"errors"
	"github.com/Yzc216/gomall/app/user/biz/dal/mysql"
	"github.com/Yzc216/gomall/app/user/biz/model"
	user "github.com/Yzc216/gomall/rpc_gen/kitex_gen/user"
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
	if req.UserId == 0 {
		return nil, errors.New("user id is required")
	}

	u := &model.User{
		ID: req.UserId,
		//Avatar:   req.UserInfo.Avatar,
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
