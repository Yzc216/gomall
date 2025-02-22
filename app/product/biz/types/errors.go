package types

import "github.com/cloudwego/kitex/pkg/kerrors"

var (
	//category
	ErrDuplicateName        = kerrors.NewGRPCBizStatusError(5002101, "分类名称已存在")
	ErrInvalidParent        = kerrors.NewGRPCBizStatusError(5002102, "无效的父分类")
	ErrHasChildren          = kerrors.NewGRPCBizStatusError(5002103, "存在子分类不可删除")
	ErrCategoryNotFound     = kerrors.NewGRPCBizStatusError(5002104, "未找到分类")
	ErrInvalidCategoryChain = kerrors.NewGRPCBizStatusError(5002105, "分类链不连续")

	//brand
	ErrBrandNotFound = kerrors.NewGRPCBizStatusError(5002201, "brand not found")
	ErrBrandExists   = kerrors.NewGRPCBizStatusError(5002202, "brand name already exists")

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
