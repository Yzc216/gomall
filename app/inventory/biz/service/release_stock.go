package service

import (
	"context"
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
	// Finish your business logic.

	return
}
