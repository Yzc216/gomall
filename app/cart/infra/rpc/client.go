package rpc

import (
	cartUtils "github.com/Yzc216/gomall/app/cart/utils"
	"github.com/Yzc216/gomall/common/clientsuite"
	"github.com/Yzc216/gomall/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/cloudwego/kitex/client"
	"sync"

	"github.com/Yzc216/gomall/app/cart/conf"
)

var (
	ProductClient productcatalogservice.Client
	Once          sync.Once
	err           error
	registryAddr  string
	serviceName   string
)

func InitClient() {
	Once.Do(func() {
		registryAddr = conf.GetConf().Registry.RegistryAddress[0]
		serviceName = conf.GetConf().Kitex.Service
		initProductClient()
	})
}

func initProductClient() {
	opts := []client.Option{
		client.WithSuite(clientsuite.CommonGrpcClientSuite{
			RegistryAddr:       registryAddr,
			CurrentServiceName: serviceName,
		}),
	}

	ProductClient, err = productcatalogservice.NewClient("product", opts...)
	cartUtils.MustHandleError(err)

}
