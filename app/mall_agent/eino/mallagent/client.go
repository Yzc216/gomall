package mallagent

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/discovery"
	cart "github.com/strings77wzq/gomall/rpc_gen/kitex_gen/cart/cartservice"
	order "github.com/strings77wzq/gomall/rpc_gen/kitex_gen/order/orderservice"
	product "github.com/strings77wzq/gomall/rpc_gen/kitex_gen/product/productservice"
	user "github.com/strings77wzq/gomall/rpc_gen/kitex_gen/user/userservice"
)

type ServiceClients struct {
	ProductClient product.Client
	CartClient    cart.Client
	OrderClient   order.Client
	UserClient    user.Client
	once          sync.Once
}

var defaultClients *ServiceClients

func initProductClient(ctx context.Context, registry discovery.Registry) (product.Client, error) {
	return product.NewClient("product-service",
		client.WithRegistry(registry),
		client.WithRPCTimeout(3*time.Second))
}

func initCartClient(ctx context.Context, registry discovery.Registry) (cart.Client, error) {
	return cart.NewClient("cart-service",
		client.WithRegistry(registry),
		client.WithRPCTimeout(3*time.Second))
}

func initOrderClient(ctx context.Context, registry discovery.Registry) (order.Client, error) {
	return order.NewClient("order-service",
		client.WithRegistry(registry),
		client.WithRPCTimeout(3*time.Second))
}

func initUserClient(ctx context.Context, registry discovery.Registry) (user.Client, error) {
	return user.NewClient("user-service",
		client.WithRegistry(registry),
		client.WithRPCTimeout(3*time.Second))
}

func InitClients(ctx context.Context, registry discovery.Registry) error {
	if defaultClients == nil {
		defaultClients = &ServiceClients{}
	}

	var err error
	defaultClients.once.Do(func() {
		if defaultClients.ProductClient, err = initProductClient(ctx, registry); err != nil {
			return
		}
		if defaultClients.CartClient, err = initCartClient(ctx, registry); err != nil {
			return
		}
		if defaultClients.OrderClient, err = initOrderClient(ctx, registry); err != nil {
			return
		}
		if defaultClients.UserClient, err = initUserClient(ctx, registry); err != nil {
			return
		}
	})
	return err
}

func GetClients() *ServiceClients {
	return defaultClients
}

// HealthCheck 检查所有服务健康状态
func (s *ServiceClients) HealthCheck(ctx context.Context) error {
	// 检查Product服务
	if s.ProductClient != nil {
		if err := s.checkProductHealth(ctx); err != nil {
			return fmt.Errorf("product service unhealthy: %v", err)
		}
	}

	// 检查Cart服务
	if s.CartClient != nil {
		if err := s.checkCartHealth(ctx); err != nil {
			return fmt.Errorf("cart service unhealthy: %v", err)
		}
	}

	// 检查Order服务
	if s.OrderClient != nil {
		if err := s.checkOrderHealth(ctx); err != nil {
			return fmt.Errorf("order service unhealthy: %v", err)
		}
	}

	// 检查User服务
	if s.UserClient != nil {
		if err := s.checkUserHealth(ctx); err != nil {
			return fmt.Errorf("user service unhealthy: %v", err)
		}
	}

	return nil
}

// 各服务健康检查实现
func (s *ServiceClients) checkProductHealth(ctx context.Context) error {
	// 这里应调用商品服务的健康检查接口
	// 为简化示例，使用超时上下文模拟检查
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	// 实际应调用ProductClient的健康检查方法
	return nil
}

func (s *ServiceClients) checkCartHealth(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	// 实际应调用CartClient的健康检查方法
	return nil
}

func (s *ServiceClients) checkOrderHealth(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	// 实际应调用OrderClient的健康检查方法
	return nil
}

// 检查User服务健康状态
func (s *ServiceClients) checkUserHealth(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	// 实际应调用UserClient的健康检查方法
	return nil
}
