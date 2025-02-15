package inventory

import (
	"context"
	inventory "github.com/Yzc216/gomall/rpc_gen/kitex_gen/inventory"

	"github.com/Yzc216/gomall/rpc_gen/kitex_gen/inventory/inventoryservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

type RPCClient interface {
	KitexClient() inventoryservice.Client
	Service() string
	QueryStock(ctx context.Context, Req *inventory.QueryStockReq, callOptions ...callopt.Option) (r *inventory.QueryStockResp, err error)
	ReserveStock(ctx context.Context, Req *inventory.InventoryReq, callOptions ...callopt.Option) (r *inventory.InventoryResp, err error)
	ConfirmStock(ctx context.Context, Req *inventory.InventoryReq, callOptions ...callopt.Option) (r *inventory.InventoryResp, err error)
	ReleaseStock(ctx context.Context, Req *inventory.InventoryReq, callOptions ...callopt.Option) (r *inventory.InventoryResp, err error)
}

func NewRPCClient(dstService string, opts ...client.Option) (RPCClient, error) {
	kitexClient, err := inventoryservice.NewClient(dstService, opts...)
	if err != nil {
		return nil, err
	}
	cli := &clientImpl{
		service:     dstService,
		kitexClient: kitexClient,
	}

	return cli, nil
}

type clientImpl struct {
	service     string
	kitexClient inventoryservice.Client
}

func (c *clientImpl) Service() string {
	return c.service
}

func (c *clientImpl) KitexClient() inventoryservice.Client {
	return c.kitexClient
}

func (c *clientImpl) QueryStock(ctx context.Context, Req *inventory.QueryStockReq, callOptions ...callopt.Option) (r *inventory.QueryStockResp, err error) {
	return c.kitexClient.QueryStock(ctx, Req, callOptions...)
}

func (c *clientImpl) ReserveStock(ctx context.Context, Req *inventory.InventoryReq, callOptions ...callopt.Option) (r *inventory.InventoryResp, err error) {
	return c.kitexClient.ReserveStock(ctx, Req, callOptions...)
}

func (c *clientImpl) ConfirmStock(ctx context.Context, Req *inventory.InventoryReq, callOptions ...callopt.Option) (r *inventory.InventoryResp, err error) {
	return c.kitexClient.ConfirmStock(ctx, Req, callOptions...)
}

func (c *clientImpl) ReleaseStock(ctx context.Context, Req *inventory.InventoryReq, callOptions ...callopt.Option) (r *inventory.InventoryResp, err error) {
	return c.kitexClient.ReleaseStock(ctx, Req, callOptions...)
}
