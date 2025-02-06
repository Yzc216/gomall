package service

import (
	"context"
	"github.com/Yzc216/gomall/app/user/biz/dal/mysql"
	"github.com/Yzc216/gomall/app/user/biz/model"
	user "github.com/Yzc216/gomall/rpc_gen/kitex_gen/user"
)

type BanUserService struct {
	ctx context.Context
} // NewBanUserService new BanUserService
func NewBanUserService(ctx context.Context) *BanUserService {
	return &BanUserService{ctx: ctx}
}

// Run create note info
func (s *BanUserService) Run(req *user.BanUserReq) (resp *user.BanUserResp, err error) {
	err = model.BanUser(s.ctx, mysql.DB, req.UserId)
	if err != nil {
		return nil, err
	}

	return &user.BanUserResp{IsBan: true}, nil
}
