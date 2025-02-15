package service

import (
	"context"
	"errors"
	"github.com/Yzc216/gomall/app/inventory/biz/dal/mysql"
	"github.com/Yzc216/gomall/app/inventory/biz/model"
	"github.com/Yzc216/gomall/app/inventory/types"
	inventory "github.com/Yzc216/gomall/rpc_gen/kitex_gen/inventory"
	"gorm.io/gorm"
	"time"
)

const maxRetries = 10

type ReserveStockService struct {
	ctx context.Context
} // NewReserveStockService new ReserveStockService
func NewReserveStockService(ctx context.Context) *ReserveStockService {
	return &ReserveStockService{ctx: ctx}
}

// Run create note info
func (s *ReserveStockService) Run(req *inventory.InventoryReq) (resp *inventory.InventoryResp, err error) {
	if req.OrderId == "" {
		return nil, types.ErrInvalidOrderId
	}
	if len(req.Items) == 0 {
		return nil, types.ErrInvalidSKU
	}

	//for _, item := range req.Items {
	//	err = ReserveStockWithOptimistic(s.ctx, mysql.DB, req.OrderId, item)
	//	if err != nil {
	//		return nil, err
	//	}
	//}
	for _, item := range req.Items {
		err = model.ReserveStockWithLock(s.ctx, mysql.DB, item.SkuId, req.OrderId, item.Quantity, false)
		if err != nil {
			return nil, err
		}
	}
	return &inventory.InventoryResp{
		Success: true,
	}, nil
}

func ReserveStockWithOptimistic(ctx context.Context, db *gorm.DB, orderId string, item *inventory.InventoryReq_Item) error {
	for i := 0; i < maxRetries; i++ {
		err := model.ReserveStockWithOptimistic(ctx, db, item.SkuId, orderId, item.Quantity, false)
		if err == nil {
			return nil
		}
		if errors.Is(err, types.ErrConcurrentModification) {
			time.Sleep(time.Duration(i) * 100 * time.Millisecond)
			continue
		}
		// 处理其他错误
		if err != nil {
			return err
		}
	}
	return types.ErrReserveStockFailed
}
