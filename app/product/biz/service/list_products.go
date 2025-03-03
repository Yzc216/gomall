package service

import (
	"context"
	"github.com/Yzc216/gomall/app/product/biz/dal/mysql"
	"github.com/Yzc216/gomall/app/product/biz/model"
	product "github.com/Yzc216/gomall/rpc_gen/kitex_gen/product"
)

type ListProductsService struct {
	ctx  context.Context
	repo *model.SPURepo
} // NewListProductsService new ListProductsService
func NewListProductsService(ctx context.Context) *ListProductsService {
	return &ListProductsService{ctx: ctx, repo: model.NewSPURepo(mysql.DB)}
}

// Run create note info
func (s *ListProductsService) Run(req *product.ListProductsReq) (resp *product.ListProductsResp, err error) {
	var filter = &model.SPUFilter{}
	var page = &model.Pagination{}
	if req.Filter != nil {
		filter = &model.SPUFilter{
			Brand:      req.Filter.Brand,
			CategoryID: req.Filter.CategoryId,
			Status:     int8(req.Filter.Status),
			MinPrice:   req.Filter.MinPrice,
			MaxPrice:   req.Filter.MaxPrice,
			Keyword:    req.Filter.Keywords,
		}
		page = &model.Pagination{
			Page:     int(req.Filter.Pagination.Page),
			PageSize: int(req.Filter.Pagination.PageSize),
		}
	}

	products, _, err := s.repo.List(s.ctx, filter, page)
	if err != nil {
		return nil, err
	}

	var SPUs = make([]*product.SPU, len(products))
	for i, v := range products {
		spu, err := convertToProtoSPU(v)
		if err != nil {
			return nil, err
		}
		SPUs[i] = spu
	}
	return &product.ListProductsResp{Products: SPUs}, nil
}
