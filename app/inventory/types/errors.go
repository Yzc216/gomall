package types

import "github.com/cloudwego/kitex/pkg/kerrors"

var (
	ErrConcurrentModification     = kerrors.NewGRPCBizStatusError(5008001, "Concurrent modification detected")
	ErrInvalidOrderId             = kerrors.NewGRPCBizStatusError(5008002, "Invalid order id")
	ErrInvalidSKU                 = kerrors.NewGRPCBizStatusError(5008003, "Invalid sku")
	ErrReserveStockFailed         = kerrors.NewGRPCBizStatusError(5008004, "Reserve stock failed")
	ErrAvailableStockInsufficient = kerrors.NewGRPCBizStatusError(5008005, "Available stock insufficient")
	ErrLockedStockInsufficient    = kerrors.NewGRPCBizStatusError(5008006, "Locked stock insufficient")
)
