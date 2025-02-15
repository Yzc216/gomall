package inventory

import (
	"context"
	inventory "github.com/Yzc216/gomall/rpc_gen/kitex_gen/inventory"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
)

func QueryStock(ctx context.Context, req *inventory.QueryStockReq, callOptions ...callopt.Option) (resp *inventory.QueryStockResp, err error) {
	resp, err = defaultClient.QueryStock(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "QueryStock call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func ReserveStock(ctx context.Context, req *inventory.InventoryReq, callOptions ...callopt.Option) (resp *inventory.InventoryResp, err error) {
	resp, err = defaultClient.ReserveStock(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "ReserveStock call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func ConfirmStock(ctx context.Context, req *inventory.InventoryReq, callOptions ...callopt.Option) (resp *inventory.InventoryResp, err error) {
	resp, err = defaultClient.ConfirmStock(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "ConfirmStock call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func ReleaseStock(ctx context.Context, req *inventory.InventoryReq, callOptions ...callopt.Option) (resp *inventory.InventoryResp, err error) {
	resp, err = defaultClient.ReleaseStock(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "ReleaseStock call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
