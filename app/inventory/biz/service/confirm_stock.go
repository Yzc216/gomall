package service

import (
	"context"
	"github.com/Yzc216/gomall/app/inventory/biz/dal/mysql"
	"github.com/Yzc216/gomall/app/inventory/biz/model"
	"github.com/Yzc216/gomall/app/inventory/types"
	inventory "github.com/Yzc216/gomall/rpc_gen/kitex_gen/inventory"
)

type ConfirmStockService struct {
	ctx context.Context
} // NewConfirmStockService new ConfirmStockService
func NewConfirmStockService(ctx context.Context) *ConfirmStockService {
	return &ConfirmStockService{ctx: ctx}
}

// Run create note info
func (s *ConfirmStockService) Run(req *inventory.InventoryReq) (resp *inventory.InventoryResp, err error) {
	if req.OrderId == "" {
		return nil, types.ErrInvalidOrderId
	}
	if len(req.Items) == 0 {
		return nil, types.ErrInvalidSKU
	}

	for _, item := range req.Items {
		err = model.ConfirmStock(s.ctx, mysql.DB, item.SkuId, req.OrderId, item.Quantity)
		if err != nil {
			return nil, err
		}
	}
	return &inventory.InventoryResp{
		Success: true,
	}, nil
}
