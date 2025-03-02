package service

import (
	"context"
	"encoding/json"
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

	SPUs := []*product.SPU{}
	for _, v := range products {
		spu, err := convertToProtoSPU(v)
		if err != nil {
			return nil, err
		}
		SPUs = append(SPUs, spu)
	}
	return &product.ListProductsResp{Products: SPUs}, nil
}

func convertToProtoSPU(v *model.SPU) (res *product.SPU, err error) {
	var skus []*product.SKU
	for _, v1 := range v.SKUs {
		var specs map[string]string
		if v1.Specs != nil {
			err := json.Unmarshal(v1.Specs, &specs)
			if err != nil {
				return nil, err
			}
		}
		skus = append(skus, &product.SKU{
			Id:       v1.ID,
			Title:    v1.Title,
			Price:    v1.Price,
			SpuId:    v1.SpuID,
			IsActive: v1.IsActive,
			Specs:    specs,
			Stock:    v1.Stock,
			Sales:    v1.Sales,
		})
	}

	res = &product.SPU{
		Id: v.ID,
		BasicInfo: &product.SPUBasicInfo{
			Title:       v.Title,
			SubTitle:    v.SubTitle,
			Description: v.Description,
			ShopId:      v.ShopID,
			Brand:       v.Brand,
			Status:      product.SPUStatus(v.Status),
		},
		Media: &product.SPUMedia{
			MainImages: v.MainImages,
			VideoUrl:   v.Video,
		},
		CategoryRelation: nil,
		Skus:             skus,
	}

	return res, nil
}
