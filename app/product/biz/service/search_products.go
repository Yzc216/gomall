package service

import (
	"context"
	"github.com/Yzc216/gomall/app/product/biz/dal/mysql"
	"github.com/Yzc216/gomall/app/product/biz/model"
	product "github.com/Yzc216/gomall/rpc_gen/kitex_gen/product"
)

type SearchProductsService struct {
	ctx context.Context
} // NewSearchProductsService new SearchProductsService
func NewSearchProductsService(ctx context.Context) *SearchProductsService {
	return &SearchProductsService{ctx: ctx}
}

// Run create note info
func (s *SearchProductsService) Run(req *product.SearchProductsReq) (resp *product.SearchProductsResp, err error) {
	productQuery := model.NewProductQuery(s.ctx, mysql.DB)

	products, _ := productQuery.SearchProducts(req.Query)

	var result []*product.Product
	for _, v1 := range products {
		result = append(result, &product.Product{
			Id:          uint32(v1.ID),
			Name:        v1.Name,
			Price:       v1.Price,
			Picture:     v1.Picture,
			Description: v1.Description,
		})
	}
	return &product.SearchProductsResp{Results: result}, nil
}
