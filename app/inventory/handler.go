package main

import (
	"context"
	"github.com/Yzc216/gomall/app/inventory/biz/service"
	inventory "github.com/Yzc216/gomall/rpc_gen/kitex_gen/inventory"
)

// InventoryServiceImpl implements the last service interface defined in the IDL.
type InventoryServiceImpl struct{}

// QueryStock implements the InventoryServiceImpl interface.
func (s *InventoryServiceImpl) QueryStock(ctx context.Context, req *inventory.QueryStockReq) (resp *inventory.QueryStockResp, err error) {
	resp, err = service.NewQueryStockService(ctx).Run(req)

	return resp, err
}

// ReserveStock implements the InventoryServiceImpl interface.
func (s *InventoryServiceImpl) ReserveStock(ctx context.Context, req *inventory.InventoryReq) (resp *inventory.InventoryResp, err error) {
	resp, err = service.NewReserveStockService(ctx).Run(req)

	return resp, err
}

// ConfirmStock implements the InventoryServiceImpl interface.
func (s *InventoryServiceImpl) ConfirmStock(ctx context.Context, req *inventory.InventoryReq) (resp *inventory.InventoryResp, err error) {
	resp, err = service.NewConfirmStockService(ctx).Run(req)

	return resp, err
}

// ReleaseStock implements the InventoryServiceImpl interface.
func (s *InventoryServiceImpl) ReleaseStock(ctx context.Context, req *inventory.InventoryReq) (resp *inventory.InventoryResp, err error) {
	resp, err = service.NewReleaseStockService(ctx).Run(req)

	return resp, err
}
