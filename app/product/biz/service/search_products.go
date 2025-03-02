package service

import (
	"context"
	"fmt"
	"github.com/Yzc216/gomall/app/product/biz/dal/mysql"
	"github.com/Yzc216/gomall/app/product/biz/model"
	product "github.com/Yzc216/gomall/rpc_gen/kitex_gen/product"
)

type SearchProductsService struct {
	ctx  context.Context
	repo *model.SPURepo
} // NewSearchProductsService new SearchProductsService
func NewSearchProductsService(ctx context.Context) *SearchProductsService {
	return &SearchProductsService{ctx: ctx, repo: model.NewSPURepo(mysql.DB)}
}

// Run create note info
func (s *SearchProductsService) Run(req *product.SearchProductsReq) (resp *product.SearchProductsResp, err error) {
	filter := model.SPUFilter{
		Keyword: req.Query,
	}
	page := model.Pagination{
		//Page:     1,
		//PageSize: 20,
	}
	products, _, err := s.repo.List(s.ctx, filter, page)
	if err != nil {
		return nil, err
	}
	fmt.Println(products)
	SPUs := []*product.SPU{}
	for _, v := range products {
		spu, err := convertToProtoSPU(v)
		if err != nil {
			return nil, err
		}
		SPUs = append(SPUs, spu)
	}
	return &product.SearchProductsResp{Results: SPUs}, nil
}
