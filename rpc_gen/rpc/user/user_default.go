package user

import (
	"context"
	user "github.com/Yzc216/gomall/rpc_gen/kitex_gen/user"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
)

func Register(ctx context.Context, req *user.RegisterReq, callOptions ...callopt.Option) (resp *user.RegisterResp, err error) {
	resp, err = defaultClient.Register(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "Register call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func Login(ctx context.Context, req *user.LoginReq, callOptions ...callopt.Option) (resp *user.LoginResp, err error) {
	resp, err = defaultClient.Login(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "Login call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func ResetPassword(ctx context.Context, req *user.ResetPasswordReq, callOptions ...callopt.Option) (resp *user.ResetPasswordResp, err error) {
	resp, err = defaultClient.ResetPassword(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "ResetPassword call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func UpdateUserInfo(ctx context.Context, req *user.UpdateUserInfoReq, callOptions ...callopt.Option) (resp *user.UpdateUserInfoResp, err error) {
	resp, err = defaultClient.UpdateUserInfo(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "UpdateUserInfo call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func SetUserRole(ctx context.Context, req *user.SetUserRoleReq, callOptions ...callopt.Option) (resp *user.SetUserRoleResp, err error) {
	resp, err = defaultClient.SetUserRole(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "SetUserRole call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func GetUserInfo(ctx context.Context, req *user.GetUserInfoReq, callOptions ...callopt.Option) (resp *user.GetUserInfoResp, err error) {
	resp, err = defaultClient.GetUserInfo(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "GetUserInfo call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func GetUserInfoList(ctx context.Context, req *user.GetUserInfoListReq, callOptions ...callopt.Option) (resp *user.GetUserInfoListResp, err error) {
	resp, err = defaultClient.GetUserInfoList(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "GetUserInfoList call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func BanUser(ctx context.Context, req *user.BanUserReq, callOptions ...callopt.Option) (resp *user.BanUserResp, err error) {
	resp, err = defaultClient.BanUser(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "BanUser call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
