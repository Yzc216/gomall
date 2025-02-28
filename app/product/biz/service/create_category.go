package service

import (
	"context"
	"fmt"
	"github.com/Yzc216/gomall/app/product/biz/dal/mysql"
	"github.com/Yzc216/gomall/app/product/biz/model"
	product "github.com/Yzc216/gomall/rpc_gen/kitex_gen/product"
	"gorm.io/gorm"
)

type CreateCategoryService struct {
	ctx  context.Context
	repo *model.CategoryRepo
}

// NewCreateCategoryService new CreateCategoryService
func NewCreateCategoryService(ctx context.Context) *CreateCategoryService {
	return &CreateCategoryService{
		ctx:  ctx,
		repo: model.NewCategoryRepo(mysql.DB),
	}
}

// Run create note info
func (s *CreateCategoryService) Run(req *product.CreateCategoryReq) (resp *product.Category, err error) {
	// 参数校验
	if err = s.Validate(req); err != nil {
		return nil, fmt.Errorf("validation error: %w", err)
	}

	// 名称冲突检查
	exists, err := s.repo.CheckNameExists(s.ctx, req.ParentId, req.Name)
	if err != nil || exists {
		return nil, fmt.Errorf("category name must be unique under parent")
	}

	err = s.repo.DB.Transaction(func(tx *gorm.DB) error {
		txRepo := model.NewCategoryRepo(tx)

		category := &model.Category{
			Name:        req.Name,
			Description: req.Description,
			ParentID:    req.ParentId,
			Image:       *req.ImageUrl,
		}

		// 创建分类
		if err = txRepo.Create(s.ctx, category); err != nil {
			return err
		}
		// 更新父节点状态
		if req.ParentId != 0 {
			return txRepo.UpdateParentLeafStatus(s.ctx, req.ParentId, false)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return &product.Category{}, nil

}

func (s *CreateCategoryService) Validate(req *product.CreateCategoryReq) error {
	if req.Name == "" {
		return fmt.Errorf("category name is required")
	}
	if req.Description == "" {
		return fmt.Errorf("category description is required")
	}
	if req.ParentId < 0 || req.ParentId > 5 {
		return fmt.Errorf("category parent id must be between 0 and 5")
	}
	return nil
}
