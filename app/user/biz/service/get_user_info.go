package service

import (
	"context"
	"errors"
	"github.com/Yzc216/gomall/app/user/biz/dal/mysql"
	"github.com/Yzc216/gomall/app/user/biz/model"
	user "github.com/Yzc216/gomall/app/user/kitex_gen/user"
)

type GetUserInfoService struct {
	ctx context.Context
} // NewGetUserInfoService new GetUserInfoService
func NewGetUserInfoService(ctx context.Context) *GetUserInfoService {
	return &GetUserInfoService{ctx: ctx}
}

// Run create note info
func (s *GetUserInfoService) Run(req *user.GetUserInfoReq) (resp *user.GetUserInfoResp, err error) {
	if req.UserId == 0 {
		return nil, errors.New("user id is empty")
	}

	u, err := model.GetById(s.ctx, mysql.DB, req.UserId)
	if err != nil {
		return nil, err
	}
	if u.Enable != 1 {
		return nil, errors.New("用户已封禁")
	}

	return &user.GetUserInfoResp{UserInfo: &user.User{
		Username: u.Username,
		Avatar:   u.Avatar,
		Email:    u.Email,
		Phone:    u.Phone,
		Role:     u.GetUint32Auth(),
	}}, nil
}
