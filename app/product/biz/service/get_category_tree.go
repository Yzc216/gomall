package service

import (
	"context"
	"errors"
	"github.com/Yzc216/gomall/app/product/biz/dal/mysql"
	"github.com/Yzc216/gomall/app/product/biz/model"
	product "github.com/Yzc216/gomall/rpc_gen/kitex_gen/product"
)

type GetCategoryTreeService struct {
	ctx  context.Context
	repo *model.CategoryRepo
} // NewGetCategoryTreeService new GetCategoryTreeService
func NewGetCategoryTreeService(ctx context.Context) *GetCategoryTreeService {
	return &GetCategoryTreeService{ctx: ctx, repo: model.NewCategoryRepo(mysql.DB)}
}

// Run create note info
func (s *GetCategoryTreeService) Run(req *product.GetCategoryTreeReq) (resp *product.CategoryTreeResp, err error) {
	categories, err := s.repo.GetAll(s.ctx)
	if err != nil {
		return nil, errors.New("获取分类数据失败")
	}

	var spuCounts map[uint64]uint32
	if req.IncludeSpuCount {
		ids := collectCategoryIDs(categories)
		spuCounts, err = s.repo.GetSPUCountsByCategoryIDs(s.ctx, ids)
		if err != nil {
			return nil, errors.New("获取商品数量失败")
		}
	}

	tree := buildCategoryTree(categories, spuCounts)
	return &product.CategoryTreeResp{Tree: tree}, nil
}

// 辅助函数
func collectCategoryIDs(categories []*model.Category) []uint64 {
	ids := make([]uint64, len(categories))
	for i, cat := range categories {
		ids[i] = cat.ID
	}
	return ids
}

func buildCategoryTree(categories []*model.Category, spuCounts map[uint64]uint32) []*product.CategoryNode {
	parentMap := make(map[uint64][]*model.Category)
	for _, cat := range categories {
		parentMap[cat.ParentID] = append(parentMap[cat.ParentID], cat)
	}

	var build func(parentID uint64) []*product.CategoryNode
	build = func(parentID uint64) []*product.CategoryNode {
		children := parentMap[parentID]
		nodes := make([]*product.CategoryNode, 0, len(children))

		for _, cat := range children {
			node := &product.CategoryNode{
				Category: convertToproductCategory(cat),
				Children: build(cat.ID),
			}
			if spuCounts != nil {
				node.SpuCount = spuCounts[cat.ID]
			}
			nodes = append(nodes, node)
		}
		return nodes
	}

	return build(0) // 根节点的ParentID为0
}

func convertToproductCategory(c *model.Category) *product.Category {
	return &product.Category{
		Id:          c.ID,
		Name:        c.Name,
		Description: c.Description,
		ParentId:    c.ParentID,
		Level:       int32(c.Level),
		IsLeaf:      c.IsLeaf,
		Sort:        int32(c.Sort),
		ImageUrl:    c.Image,
	}
}
