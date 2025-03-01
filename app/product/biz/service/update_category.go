package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/Yzc216/gomall/app/product/biz/dal/mysql"
	"github.com/Yzc216/gomall/app/product/biz/model"
	"github.com/Yzc216/gomall/app/product/biz/types"
	product "github.com/Yzc216/gomall/rpc_gen/kitex_gen/product"
	"net/url"
)

type UpdateCategoryService struct {
	ctx  context.Context
	repo *model.CategoryRepo
}

// NewUpdateCategoryService new UpdateCategoryService
func NewUpdateCategoryService(ctx context.Context) *UpdateCategoryService {
	return &UpdateCategoryService{
		ctx:  ctx,
		repo: model.NewCategoryRepo(mysql.DB),
	}
}

// Run create note info
func (s *UpdateCategoryService) Run(req *product.UpdateCategoryReq) (resp *product.Category, err error) {
	// 1. 校验必要参数
	if req.Id == 0 {
		return nil, fmt.Errorf("%w: missing category id", types.ErrInvalidUpdate)
	}

	// 2. 获取现有分类
	existing, err := s.repo.GetByID(s.ctx, req.Id)
	if err != nil {
		if errors.Is(err, types.ErrRecordNotFound) {
			return nil, types.ErrCategoryNotFound
		}
		return nil, err
	}

	// 3. 构建更新字段
	updates := make(map[string]interface{})

	// 处理名称更新
	if req.Name != nil {
		// 同名检查（同一父级下名称唯一）
		exist, err := s.repo.ExistByName(s.ctx, existing.ParentID, *req.Name, req.Id)
		if err != nil {
			return nil, err
		}
		if exist {
			return nil, types.ErrCategoryNameExists
		}
		updates["name"] = *req.Name
	}

	// 处理描述更新
	if req.Description != nil {
		updates["description"] = *req.Description
	}

	// 处理排序更新
	if req.Sort != nil {
		if *req.Sort < 0 {
			return nil, fmt.Errorf("%w: sort value cannot be negative", types.ErrInvalidUpdate)
		}
		updates["sort"] = *req.Sort
	}

	// 处理图片URL更新
	if req.ImageUrl != nil {
		if !isValidURL(*req.ImageUrl) {
			return nil, fmt.Errorf("%w: invalid image url format", types.ErrInvalidUpdate)
		}
		updates["image"] = *req.ImageUrl
	}

	// 4. 执行更新
	if len(updates) == 0 {
		return nil, types.ErrInvalidUpdate
	}

	if err := s.repo.UpdatePartial(s.ctx, req.Id, updates); err != nil {
		return nil, fmt.Errorf("update category failed: %w", err)
	}

	return &product.Category{Id: existing.ID}, nil
}

// 辅助函数：验证URL格式
func isValidURL(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}
