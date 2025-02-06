package service

import (
	"context"
	"errors"
	"github.com/Yzc216/gomall/app/user/biz/dal/mysql"
	"github.com/Yzc216/gomall/app/user/biz/model"
	user "github.com/Yzc216/gomall/rpc_gen/kitex_gen/user"
)

const (
	AddAuthority    = uint32(0)
	UpdateAuthority = uint32(1)
)

type SetUserRoleService struct {
	ctx context.Context
} // NewSetUserRoleService new SetUserRoleService
func NewSetUserRoleService(ctx context.Context) *SetUserRoleService {
	return &SetUserRoleService{ctx: ctx}
}

// Run create note info
func (s *SetUserRoleService) Run(req *user.SetUserRoleReq) (resp *user.SetUserRoleResp, err error) {
	if req.UserId == 0 {
		return nil, errors.New("user id is required")
	}
	err = model.UpdateAuthority(s.ctx, mysql.DB, req.UserId, req.NewRole)
	if err != nil {
		return nil, err
	}

	return &user.SetUserRoleResp{IsSet: true}, nil
}
