package types

import "github.com/cloudwego/kitex/pkg/kerrors"

var (
	//category
	ErrNoRowsAffected     = kerrors.NewGRPCBizStatusError(5002101, "no rows affected")
	ErrRecordNotFound     = kerrors.NewGRPCBizStatusError(5002102, "record not found")
	ErrCategoryNotFound   = kerrors.NewGRPCBizStatusError(5002103, "category not found")
	ErrCategoryNameExists = kerrors.NewGRPCBizStatusError(5002104, "category name already exists")
	ErrInvalidUpdate      = kerrors.NewGRPCBizStatusError(5002105, "invalid update")
	ErrHasChildren        = kerrors.NewGRPCBizStatusError(5002106, "category has children")
	ErrAssociatedSPUs     = kerrors.NewGRPCBizStatusError(5002107, "category has associated SPUs")

	//SPU
	ErrHasAssociatedSPUs = kerrors.NewGRPCBizStatusError(5002301, "存在关联SPU不可删除")
	ErrSPUNotFound       = kerrors.NewGRPCBizStatusError(5002303, "部分SPU不存在")
	ErrSPUTitleExists    = kerrors.NewGRPCBizStatusError(5002304, "SPU title already exists")

	//SKU
	ErrHasAssociatedSKUs = kerrors.NewGRPCBizStatusError(5002302, "cannot delete SPU with associated SKUs")
	ErrSKUNotFound       = kerrors.NewGRPCBizStatusError(5002401, "SKU not found")
	ErrStockInsufficient = kerrors.NewGRPCBizStatusError(5002401, "insufficient stock")
	ErrDuplicateSKUTitle = kerrors.NewGRPCBizStatusError(5002401, "duplicate SKU title in SPU")
	ErrSKUInUse          = kerrors.NewGRPCBizStatusError(5002401, "SKU has existing orders")

	//common
	ErrInvalidIDs        = kerrors.NewGRPCBizStatusError(5002305, "包含无效的ID")
	ErrInvalidStatus     = kerrors.NewGRPCBizStatusError(5002010, "invalid status transition")
	ErrCircularReference = kerrors.NewGRPCBizStatusError(5002005, "检测到循环引用")
)
