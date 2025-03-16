package service

import (
	"context"
	"github.com/Yzc216/gomall/app/product/biz/dal/mysql"
	"github.com/Yzc216/gomall/app/product/biz/repo"
	product "github.com/Yzc216/gomall/rpc_gen/kitex_gen/product"
)

type SearchProductsService struct {
	ctx  context.Context
	repo *repo.SPUQuery
} // NewSearchProductsService new SearchProductsService
func NewSearchProductsService(ctx context.Context) *SearchProductsService {
	return &SearchProductsService{ctx: ctx, repo: repo.NewSPUQuery(mysql.DB)}
}

// Run create note info
func (s *SearchProductsService) Run(req *product.SearchProductsReq) (resp *product.SearchProductsResp, err error) {
	filter := &repo.SPUFilter{
		Keyword: req.Query,
	}
	var page = &repo.Pagination{}
	// TODO 分页待proto补充
	//if req.page != nil {
	//	page = &model.Pagination{
	//		Page:     int(req.Filter.Pagination.Page),
	//		PageSize: int(req.Filter.Pagination.PageSize),
	//	}
	//}

	products, _, err := s.repo.List(s.ctx, filter, page)
	if err != nil {
		return nil, err
	}
	var SPUs []*product.SPU
	for _, v := range products {
		spu, err := convertToProtoSPU(v)
		if err != nil {
			return nil, err
		}
		SPUs = append(SPUs, spu)
	}
	return &product.SearchProductsResp{Results: SPUs}, nil
}
