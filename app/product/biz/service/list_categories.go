package service

import (
	"context"
	"errors"
	"github.com/Yzc216/gomall/app/product/biz/dal/mysql"
	"github.com/Yzc216/gomall/app/product/biz/model"
	"github.com/Yzc216/gomall/app/product/biz/repo"
	"github.com/Yzc216/gomall/app/product/biz/types"
	product "github.com/Yzc216/gomall/rpc_gen/kitex_gen/product"
)

type ListCategoriesService struct {
	ctx  context.Context
	repo *repo.CategoryRepo
} // NewListCategoriesService new ListCategoriesService
func NewListCategoriesService(ctx context.Context) *ListCategoriesService {
	return &ListCategoriesService{
		ctx:  ctx,
		repo: repo.NewCategoryRepo(mysql.DB),
	}
}

// Run create note info
func (s *ListCategoriesService) Run(req *product.ListCategoriesReq) (resp *product.CategoryNode, err error) {
	// 参数校验
	var c *model.Category
	if req.Id > 0 {
		if c, err = s.repo.GetByID(s.ctx, req.Id); err != nil || c == nil {
			return nil, types.ErrCategoryNotFound
		}
	} else {
		return nil, types.ErrInvalidIDs
	}

	// 获取分类树
	categories, err := s.repo.GetCategoryTreeByID(s.ctx, req.Id, req.WithSpus)
	if err != nil {
		return nil, err
	}

	// 获取SPU统计
	var spuCounts map[uint64]uint32
	if req.WithSpus {
		ids := collectCategoryIDs(categories)
		spuCounts, err = s.repo.GetSPUCountsByCategoryIDs(s.ctx, ids)
		if err != nil {
			return nil, errors.New("failed to get spu counts")
		}
	}

	tree := buildCategoryTree(categories, spuCounts)

	r := &product.CategoryNode{
		Category: tree[0].Category,
	}
	if req.WithChildren {
		r.Children = tree[0].Children
	}
	if req.WithSpus {
		r.SpuCount = spuCounts[req.Id]
	}
	return r, nil
}
