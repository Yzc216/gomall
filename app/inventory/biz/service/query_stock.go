package service

import (
	"context"
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
	// Finish your business logic.

	return
}
