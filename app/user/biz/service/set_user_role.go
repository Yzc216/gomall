package service

import (
	"context"
	"github.com/Yzc216/gomall/app/user/biz/dal/mysql"
	"github.com/Yzc216/gomall/app/user/biz/model"
	user "github.com/Yzc216/gomall/app/user/kitex_gen/user"
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
	if req.SetType == AddAuthority { //0:添加
		err = model.AddAuthority(s.ctx, mysql.DB, req.UserId, req.NewRole)
		if err != nil {
			return nil, err
		}
	} else if req.SetType == UpdateAuthority { //1:修改
		err = model.UpdateAuthority(s.ctx, mysql.DB, req.UserId, req.NewRole)
		if err != nil {
			return nil, err
		}
	}
	return
}
