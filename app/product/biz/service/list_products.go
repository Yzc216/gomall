package service

import (
	"context"
	"github.com/Yzc216/gomall/app/product/biz/dal/mysql"
	"github.com/Yzc216/gomall/app/product/biz/model"
	product "github.com/Yzc216/gomall/rpc_gen/kitex_gen/product"
)

type ListProductsService struct {
	ctx context.Context
} // NewListProductsService new ListProductsService
func NewListProductsService(ctx context.Context) *ListProductsService {
	return &ListProductsService{ctx: ctx}
}

// Run create note info
func (s *ListProductsService) Run(req *product.ListProductsReq) (resp *product.ListProductsResp, err error) {
	categoryQuery := model.NewCategoryQuery(s.ctx, mysql.DB)

	c, err := categoryQuery.GetProductsByCategoryName(req.CategoryName)
	if err != nil {
		return nil, err
	}

	resp = &product.ListProductsResp{}
	for _, v := range c {
		for _, v1 := range v.Products {
			resp.Products = append(resp.Products, &product.Product{
				Id:          uint32(v1.ID),
				Name:        v1.Name,
				Price:       v1.Price,
				Picture:     v1.Picture,
				Description: v1.Description,
			})
		}
	}

	return
}
