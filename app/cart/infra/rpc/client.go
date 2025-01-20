package rpc

import (
	cartUtils "github.com/Yzc216/gomall/app/cart/utils"
	"github.com/Yzc216/gomall/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
	"sync"

	"github.com/Yzc216/gomall/app/cart/conf"
)

var (
	ProductClient productcatalogservice.Client
	Once          sync.Once
)

func Init() {
	Once.Do(func() {

		initProductClient()
	})
}

func initProductClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	cartUtils.MustHandleError(err)
	opts = append(opts, client.WithResolver(r))

	ProductClient, err = productcatalogservice.NewClient("product", opts...)
	cartUtils.MustHandleError(err)

}
