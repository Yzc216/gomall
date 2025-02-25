package user

import (
	"context"
	"errors"

	"github.com/Yzc216/gomall/app/frontend/biz/service"
	"github.com/Yzc216/gomall/app/frontend/biz/utils"
	common "github.com/Yzc216/gomall/app/frontend/hertz_gen/frontend/common"
	user "github.com/Yzc216/gomall/app/frontend/hertz_gen/frontend/user"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// User .
// @router /user/profile [GET]
func User(ctx context.Context, c *app.RequestContext) {
	var err error
	var req common.Empty
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewUserService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	c.HTML(consts.StatusOK, "user", utils.WarpResponse(ctx, c, resp))
}

// UpdateUser .
// @router /user/profile [POST]
func UpdateUser(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.UpdateUserReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	//fmt.Println(req)
	resp, err := service.NewUpdateUserService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	if !resp {
		utils.SendErrResponse(ctx, c, consts.StatusInternalServerError, errors.New("用户信息保存失败"))
	}
	c.Redirect(consts.StatusFound, []byte("/user/profile?success=true&message=用户信息保存成功"))
}

// ResetPassword .
// @router /user/password [POST]
func ResetPassword(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.ResetPasswordReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewResetPasswordService(ctx, c).Run(&req)

	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	if !resp {
		utils.SendErrResponse(ctx, c, consts.StatusInternalServerError, errors.New("reset password failed"))
	}

	c.Redirect(consts.StatusFound, []byte("/user/profile?success=true&message=密码修改成功"))

}

// Admin .
// @router /admin/users [GET]
func Admin(ctx context.Context, c *app.RequestContext) {
	var err error
	var req common.Empty
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewAdminService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	c.HTML(consts.StatusOK, "admin", utils.WarpResponse(ctx, c, resp))
}

// BanUser .
// @router /admin/users/ban [POST]
func BanUser(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.BanUserReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewBanUserService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// SetRole .
// @router /admin/users/role [POST]
func SetRole(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.SetRoleReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewSetRoleService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}
