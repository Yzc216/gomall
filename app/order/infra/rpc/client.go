package rpc

import (
	"github.com/Yzc216/gomall/app/order/conf"
	"github.com/Yzc216/gomall/common/clientsuite"
	"github.com/Yzc216/gomall/rpc_gen/kitex_gen/cart/cartservice"
	"github.com/Yzc216/gomall/rpc_gen/kitex_gen/inventory/inventoryservice"
	"github.com/Yzc216/gomall/rpc_gen/kitex_gen/order/orderservice"
	"github.com/Yzc216/gomall/rpc_gen/kitex_gen/payment/paymentservice"
	"github.com/Yzc216/gomall/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/cloudwego/kitex/client"
	"sync"
)

var (
	InventoryClient inventoryservice.Client
	CartClient      cartservice.Client
	ProductClient   productcatalogservice.Client
	PaymentClient   paymentservice.Client
	OrderClient     orderservice.Client
	once            sync.Once
	err             error
	registryAddr    string
	serviceName     string
	commonSuite     client.Option
)

func InitClient() {
	once.Do(func() {
		registryAddr = conf.GetConf().Registry.RegistryAddress[0]
		serviceName = conf.GetConf().Kitex.Service
		commonSuite = client.WithSuite(clientsuite.CommonGrpcClientSuite{
			CurrentServiceName: serviceName,
			RegistryAddr:       registryAddr,
		})
		initInventoryClient()
		initCartClient()
		initProductClient()
		initPaymentClient()
		initOrderClient()
	})
}

func initInventoryClient() {
	InventoryClient, err = inventoryservice.NewClient("inventory", commonSuite)
	if err != nil {
		panic(err)
	}
}

func initCartClient() {
	CartClient, err = cartservice.NewClient("cart", commonSuite)
	if err != nil {
		panic(err)
	}
}

func initProductClient() {
	ProductClient, err = productcatalogservice.NewClient("product", commonSuite)
	if err != nil {
		panic(err)
	}
}

func initPaymentClient() {
	PaymentClient, err = paymentservice.NewClient("payment", commonSuite)
	if err != nil {
		panic(err)
	}
}

func initOrderClient() {
	OrderClient, err = orderservice.NewClient("order", commonSuite)
	if err != nil {
		panic(err)
	}
}
