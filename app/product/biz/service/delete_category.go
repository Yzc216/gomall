package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/Yzc216/gomall/app/product/biz/dal/mysql"
	"github.com/Yzc216/gomall/app/product/biz/model"
	"github.com/Yzc216/gomall/app/product/biz/types"
	"github.com/Yzc216/gomall/rpc_gen/kitex_gen/common"
	product "github.com/Yzc216/gomall/rpc_gen/kitex_gen/product"
)

type DeleteCategoryService struct {
	ctx  context.Context
	repo *model.CategoryRepo
} // NewDeleteCategoryService new DeleteCategoryService
func NewDeleteCategoryService(ctx context.Context) *DeleteCategoryService {
	return &DeleteCategoryService{
		ctx:  ctx,
		repo: model.NewCategoryRepo(mysql.DB),
	}
}

// Run create note info
func (s *DeleteCategoryService) Run(req *product.DeleteCategoryReq) (resp *common.Empty, err error) {
	// 1. 参数校验
	if req.Id == 0 {
		return nil, errors.New("id is required")
	}

	// 2. 获取分类信息
	category, err := s.repo.GetChildren(s.ctx, req.Id)
	fmt.Println(len(category))
	if errors.Is(err, types.ErrRecordNotFound) {
		return nil, types.ErrCategoryNotFound
	}
	if err != nil {
		return nil, errors.New("get category failed")
	}

	// 3. 检查子分类
	if len(category) > 0 {
		return nil, types.ErrHasChildren
	}

	// 4. 执行删除
	if err = s.repo.DeleteCascade(s.ctx, req.Id, req.Force); err != nil {
		return nil, err
	}

	return &common.Empty{}, nil
}
