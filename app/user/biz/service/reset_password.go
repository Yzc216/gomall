package service

import (
	"context"
	"errors"
	"github.com/Yzc216/gomall/app/user/biz/dal/mysql"
	"github.com/Yzc216/gomall/app/user/biz/model"
	user "github.com/Yzc216/gomall/app/user/kitex_gen/user"
)

type ResetPasswordService struct {
	ctx context.Context
} // NewResetPasswordService new ResetPasswordService
func NewResetPasswordService(ctx context.Context) *ResetPasswordService {
	return &ResetPasswordService{ctx: ctx}
}

// Run create note info
func (s *ResetPasswordService) Run(req *user.ResetPasswordReq) (resp *user.ResetPasswordResp, err error) {
	if req.UserId == 0 {
		return nil, errors.New("user id is required")
	}
	u := &model.User{
		ID:       req.UserId,
		Password: req.Password,
	}

	err = model.UpdatePassword(s.ctx, mysql.DB, u, req.NewPassword)
	if err != nil {
		return nil, err
	}

	return &user.ResetPasswordResp{IsReset: true}, nil
}
