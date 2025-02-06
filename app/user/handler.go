package main

import (
	"context"
	"github.com/Yzc216/gomall/app/user/biz/service"
	user "github.com/Yzc216/gomall/rpc_gen/kitex_gen/user"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *user.RegisterReq) (resp *user.RegisterResp, err error) {
	resp, err = service.NewRegisterService(ctx).Run(req)

	return resp, err
}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *user.LoginReq) (resp *user.LoginResp, err error) {
	resp, err = service.NewLoginService(ctx).Run(req)

	return resp, err
}

// ResetPassword implements the UserServiceImpl interface.
func (s *UserServiceImpl) ResetPassword(ctx context.Context, req *user.ResetPasswordReq) (resp *user.ResetPasswordResp, err error) {
	resp, err = service.NewResetPasswordService(ctx).Run(req)

	return resp, err
}

// UpdateUserInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) UpdateUserInfo(ctx context.Context, req *user.UpdateUserInfoReq) (resp *user.UpdateUserInfoResp, err error) {
	resp, err = service.NewUpdateUserInfoService(ctx).Run(req)

	return resp, err
}

// SetUserRole implements the UserServiceImpl interface.
func (s *UserServiceImpl) SetUserRole(ctx context.Context, req *user.SetUserRoleReq) (resp *user.SetUserRoleResp, err error) {
	resp, err = service.NewSetUserRoleService(ctx).Run(req)

	return resp, err
}

// GetUserInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUserInfo(ctx context.Context, req *user.GetUserInfoReq) (resp *user.GetUserInfoResp, err error) {
	resp, err = service.NewGetUserInfoService(ctx).Run(req)

	return resp, err
}

// GetUserInfoList implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUserInfoList(ctx context.Context, req *user.GetUserInfoListReq) (resp *user.GetUserInfoListResp, err error) {
	resp, err = service.NewGetUserInfoListService(ctx).Run(req)

	return resp, err
}

// BanUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) BanUser(ctx context.Context, req *user.BanUserReq) (resp *user.BanUserResp, err error) {
	resp, err = service.NewBanUserService(ctx).Run(req)

	return resp, err
}
