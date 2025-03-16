package service

import (
	"context"
	"fmt"
	"github.com/Yzc216/gomall/app/product/biz/dal/mysql"
	"github.com/Yzc216/gomall/app/product/biz/model"
	"github.com/Yzc216/gomall/app/product/biz/repo"
	product "github.com/Yzc216/gomall/rpc_gen/kitex_gen/product"
	"gorm.io/gorm"
)

type CreateCategoryService struct {
	ctx  context.Context
	repo *repo.CategoryRepo
}

// NewCreateCategoryService new CreateCategoryService
func NewCreateCategoryService(ctx context.Context) *CreateCategoryService {
	return &CreateCategoryService{
		ctx:  ctx,
		repo: repo.NewCategoryRepo(mysql.DB),
	}
}

// Run create note info
func (s *CreateCategoryService) Run(req *product.CreateCategoryReq) (resp *product.Category, err error) {
	// 参数校验
	if err = s.Validate(req); err != nil {
		return nil, fmt.Errorf("validation error: %w", err)
	}

	// 名称冲突检查
	exists, err := s.repo.ExistByName(s.ctx, req.ParentId, req.Name, 0)
	if err != nil || exists {
		return nil, fmt.Errorf("category name must be unique under parent")
	}

	var c *model.Category
	err = s.repo.DB.Transaction(func(tx *gorm.DB) error {
		txRepo := repo.NewCategoryRepo(tx)

		category := &model.Category{
			Name:        req.Name,
			Description: req.Description,
			ParentID:    req.ParentId,
		}
		if req.ImageUrl != nil {
			category.Image = *req.ImageUrl
		}

		// 创建分类
		if c, err = txRepo.Create(s.ctx, category); err != nil {
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

	return &product.Category{
		Id:          c.ID,
		Name:        c.Name,
		Description: c.Description,
		ParentId:    c.ParentID,
		Level:       int32(c.Level),
		IsLeaf:      c.IsLeaf,
		Sort:        int32(c.Sort),
		ImageUrl:    c.Image,
	}, nil

}

func (s *CreateCategoryService) Validate(req *product.CreateCategoryReq) error {
	if req.Name == "" {
		return fmt.Errorf("category name is required")
	}
	if req.Description == "" {
		return fmt.Errorf("category description is required")
	}
	if req.ParentId < 0 {
		return fmt.Errorf("category parent id is invalid")
	}
	return nil
}
