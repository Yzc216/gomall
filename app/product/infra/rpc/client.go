package rpc

import (
	"github.com/Yzc216/gomall/app/checkout/conf"
	"github.com/Yzc216/gomall/common/clientsuite"
	"github.com/Yzc216/gomall/rpc_gen/kitex_gen/inventory/inventoryservice"
	"github.com/cloudwego/kitex/client"
	"sync"
)

var (
	InventoryClient inventoryservice.Client
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
	})
}

func initInventoryClient() {
	InventoryClient, err = inventoryservice.NewClient("inventory", commonSuite)
	if err != nil {
		panic(err)
	}
}
