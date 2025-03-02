package clients

import (
	"context"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/discovery"
	"github.com/cloudwego/kitex/pkg/loadbalance"
	"github.com/cloudwego/kitex/pkg/rpcinfo"

	// 导入各个微服务的客户端
	"github.com/strings77wzq/gomall/rpc_gen/kitex_gen/cart/cartservice"
	"github.com/strings77wzq/gomall/rpc_gen/kitex_gen/order/orderservice"
	"github.com/strings77wzq/gomall/rpc_gen/kitex_gen/product/productservice"
	"github.com/strings77wzq/gomall/rpc_gen/kitex_gen/user/userservice"
)

// RPCClients 保存所有RPC客户端实例
type RPCClients struct {
	ProductClient productservice.Client
	CartClient    cartservice.Client
	OrderClient   orderservice.Client
	UserClient    userservice.Client
}

// NewRPCClients 创建所有微服务的RPC客户端
func NewRPCClients(ctx context.Context, registry discovery.Resolver) (*RPCClients, error) {
	productClient, err := newProductClient(registry)
	if err != nil {
		return nil, err
	}

	cartClient, err := newCartClient(registry)
	if err != nil {
		return nil, err
	}

	orderClient, err := newOrderClient(registry)
	if err != nil {
		return nil, err
	}

	userClient, err := newUserClient(registry)
	if err != nil {
		return nil, err
	}

	return &RPCClients{
		ProductClient: productClient,
		CartClient:    cartClient,
		OrderClient:   orderClient,
		UserClient:    userClient,
	}, nil
}

// 创建商品服务客户端
func newProductClient(registry discovery.Resolver) (productservice.Client, error) {
	return productservice.NewClient(
		"product",
		client.WithResolver(registry),
		client.WithLoadBalancer(loadbalance.NewWeightedBalancer()),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: "mall_agent",
		}),
	)
}

// 创建购物车服务客户端
func newCartClient(registry discovery.Resolver) (cartservice.Client, error) {
	return cartservice.NewClient(
		"cart",
		client.WithResolver(registry),
		client.WithLoadBalancer(loadbalance.NewWeightedBalancer()),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: "mall_agent",
		}),
	)
}

// 创建订单服务客户端
func newOrderClient(registry discovery.Resolver) (orderservice.Client, error) {
	return orderservice.NewClient(
		"order",
		client.WithResolver(registry),
		client.WithLoadBalancer(loadbalance.NewWeightedBalancer()),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: "mall_agent",
		}),
	)
}

// 创建用户服务客户端
func newUserClient(registry discovery.Resolver) (userservice.Client, error) {
	return userservice.NewClient(
		"user",
		client.WithResolver(registry),
		client.WithLoadBalancer(loadbalance.NewWeightedBalancer()),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: "mall_agent",
		}),
	)
}
