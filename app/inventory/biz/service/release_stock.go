package service

import (
	"context"
	"github.com/Yzc216/gomall/app/inventory/biz/dal/mysql"
	"github.com/Yzc216/gomall/app/inventory/biz/model"
	"github.com/Yzc216/gomall/app/inventory/types"
	inventory "github.com/Yzc216/gomall/rpc_gen/kitex_gen/inventory"
)

type ReleaseStockService struct {
	ctx context.Context
} // NewReleaseStockService new ReleaseStockService
func NewReleaseStockService(ctx context.Context) *ReleaseStockService {
	return &ReleaseStockService{ctx: ctx}
}

// Run create note info
func (s *ReleaseStockService) Run(req *inventory.InventoryReq) (resp *inventory.InventoryResp, err error) {
	if req.OrderId == "" {
		return nil, types.ErrInvalidOrderId
	}
	if len(req.Items) == 0 {
		return nil, types.ErrInvalidSKU
	}

	for _, item := range req.Items {
		err = model.ReleaseStockWithLock(s.ctx, mysql.DB, item.SkuId, req.OrderId, item.Quantity, req.Force)
		if err != nil {
			return nil, err
		}
	}
	return &inventory.InventoryResp{
		Success: true,
	}, nil
}
