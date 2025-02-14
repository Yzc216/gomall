package service

import (
	"context"
	"github.com/Yzc216/gomall/app/inventory/biz/dal/mysql"
	"github.com/Yzc216/gomall/app/inventory/biz/model"
	"github.com/Yzc216/gomall/app/inventory/types"
	inventory "github.com/Yzc216/gomall/rpc_gen/kitex_gen/inventory"
)

type QueryStockService struct {
	ctx context.Context
} // NewQueryStockService new QueryStockService
func NewQueryStockService(ctx context.Context) *QueryStockService {
	return &QueryStockService{ctx: ctx}
}

// Run create note info
func (s *QueryStockService) Run(req *inventory.QueryStockReq) (resp *inventory.QueryStockResp, err error) {
	if len(req.SkuId) == 0 {
		return nil, types.ErrInvalidSKU
	}

	invs, err := model.GetStock(s.ctx, mysql.DB, req.SkuId)
	if err != nil {
		return nil, err
	}

	Stocks := make(map[uint64]uint32, len(invs))
	for _, inv := range invs {
		Stocks[inv.SkuID] = inv.Total
	}

	return &inventory.QueryStockResp{
		CurrentStock: Stocks,
	}, nil
}
