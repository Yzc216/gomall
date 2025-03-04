package category

import (
	"context"

	"github.com/Yzc216/gomall/app/frontend/biz/service"
	"github.com/Yzc216/gomall/app/frontend/biz/utils"
	category "github.com/Yzc216/gomall/app/frontend/hertz_gen/frontend/category"
	common "github.com/Yzc216/gomall/app/frontend/hertz_gen/frontend/common"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// ListCategory .
// @router /categories [GET]
func ListCategory(ctx context.Context, c *app.RequestContext) {
	var err error
	var req common.Empty
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewListCategoryService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	c.HTML(consts.StatusOK, "categories", utils.WarpResponse(ctx, c, resp))
}

// GetProductByCategoryID .
// @router /product/category/:category_id/:category_name [GET]
func GetProductByCategoryID(ctx context.Context, c *app.RequestContext) {
	var err error
	var req category.GetProductByCategoryIDReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewGetProductByCategoryIDService(ctx, c).Run(&req)

	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	c.HTML(consts.StatusOK, "category", utils.WarpResponse(ctx, c, resp))
}

// CreateCategory .
// @router /category [POST]
func CreateCategory(ctx context.Context, c *app.RequestContext) {
	var err error
	var req category.CreateCategoryReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &category.CreateCategoryResp{}
	resp, err = service.NewCreateCategoryService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// CategoryManagement .
// @router /admin/categories [GET]
func CategoryManagement(ctx context.Context, c *app.RequestContext) {
	var err error
	var req common.Empty
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewCategoryManagementService(ctx, c).Run(&req)

	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	c.HTML(consts.StatusOK, "categories-admin", utils.WarpResponse(ctx, c, resp))
}

// UpdateCategory .
// @router /admin/categories [PUT]
func UpdateCategory(ctx context.Context, c *app.RequestContext) {
	var err error
	var req category.UpdateCategoryReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	_, err = service.NewUpdateCategoryService(ctx, c).Run(&req)

	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	c.Redirect(consts.StatusFound, []byte("/admin/categories"))
}

// DeleteCategory .
// @router /admin/categories [DELETE]
func DeleteCategory(ctx context.Context, c *app.RequestContext) {
	var err error
	var req category.DeleteCategoryReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	_, err = service.NewDeleteCategoryService(ctx, c).Run(&req)

	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	c.Redirect(consts.StatusFound, []byte("/admin/categories"))
}
