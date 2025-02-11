package service

import (
	"context"
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
	// Finish your business logic.

	return
}
