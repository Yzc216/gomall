package service

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/Yzc216/gomall/app/product/biz/dal/mysql"
	"github.com/Yzc216/gomall/app/product/biz/dal/redis"
	"github.com/Yzc216/gomall/app/product/biz/model"
	"github.com/Yzc216/gomall/app/product/biz/repo"
	"github.com/Yzc216/gomall/app/product/infra/rpc"
	"github.com/Yzc216/gomall/rpc_gen/kitex_gen/inventory"
	product "github.com/Yzc216/gomall/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/kerrors"
)

type GetProductService struct {
	ctx  context.Context
	repo *repo.CachedProductQuery
} // NewGetProductService new GetProductService
func NewGetProductService(ctx context.Context) *GetProductService {
	return &GetProductService{
		ctx:  ctx,
		repo: repo.NewCachedProductQuery(mysql.DB, redis.RedisClient)}
}

// Run create note info
func (s *GetProductService) Run(req *product.GetProductReq) (resp *product.GetProductResp, err error) {
	if req.Id == 0 {
		return nil, kerrors.NewBizStatusError(40000, "product id is required")
	}
	spu, err := s.repo.GetByID(s.ctx, req.Id)
	if err != nil {
		return nil, err
	}

	//查询库存
	skus := spu.SKUs
	var skuIds = make([]uint64, 0, len(skus))
	for _, v := range skus {
		skuIds = append(skuIds, v.ID)
	}
	stockMap, err := rpc.InventoryClient.QueryStock(s.ctx, &inventory.QueryStockReq{SkuId: skuIds})
	if err != nil {
		return nil, err
	}
	if stockMap.CurrentStock == nil {
		return nil, errors.New("获取库存失败")
	}
	for i := range spu.SKUs {
		sku := &spu.SKUs[i] // 获取元素的指针
		if newStock, exists := stockMap.CurrentStock[sku.ID]; exists {
			sku.Stock = newStock // 直接更新指针指向的 SKU 的库存
		}
	}

	protoSPU, err := convertToProtoSPU(spu)
	if err != nil {
		return nil, err
	}

	return &product.GetProductResp{Product: protoSPU}, nil
}

func convertToProtoSPU(v *model.SPU) (res *product.SPU, err error) {
	var skus = make([]*product.SKU, len(v.SKUs))
	for i, v1 := range v.SKUs {
		var specs map[string]string
		if v1.Specs != nil {
			err := json.Unmarshal(v1.Specs, &specs)
			if err != nil {
				return nil, err
			}
		}
		skus[i] = &product.SKU{
			Id:       v1.ID,
			Title:    v1.Title,
			Price:    v1.Price,
			SpuId:    v1.SpuID,
			IsActive: v1.IsActive,
			Specs:    specs,
			Stock:    v1.Stock,
			Sales:    v1.Sales,
		}
	}

	var categoryIDs = make([]uint64, len(v.Categories))
	for i, v := range v.Categories {
		categoryIDs[i] = v.ID
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
		CategoryRelation: &product.CategoryRelation{
			CategoryIds: categoryIDs,
		},
		Skus: skus,
	}

	return res, nil
}
