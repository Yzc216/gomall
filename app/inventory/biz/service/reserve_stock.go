package service

import (
	"context"
	inventory "github.com/Yzc216/gomall/rpc_gen/kitex_gen/inventory"
)

type ReserveStockService struct {
	ctx context.Context
} // NewReserveStockService new ReserveStockService
func NewReserveStockService(ctx context.Context) *ReserveStockService {
	return &ReserveStockService{ctx: ctx}
}

// Run create note info
func (s *ReserveStockService) Run(req *inventory.InventoryReq) (resp *inventory.InventoryResp, err error) {
	// Finish your business logic.

	return
}
