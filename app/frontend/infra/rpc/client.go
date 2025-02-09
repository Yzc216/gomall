package rpc

import (
	"github.com/Yzc216/gomall/app/frontend/conf"
	frontendUtils "github.com/Yzc216/gomall/app/frontend/utils"
	"github.com/Yzc216/gomall/rpc_gen/kitex_gen/cart/cartservice"
	"github.com/Yzc216/gomall/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/Yzc216/gomall/rpc_gen/kitex_gen/user/userservice"
	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
	"sync"
)

var (
	UserClient    userservice.Client
	ProductClient productcatalogservice.Client
	CartClient    cartservice.Client
	Once          sync.Once
)

func Init() {
	Once.Do(func() {
		initUserClient()
		initProductClient()
		initCartClient()
	})
}

func initUserClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	frontendUtils.MustHandleError(err)
	opts = append(opts, client.WithResolver(r))

	UserClient, err = userservice.NewClient("user", opts...)
	frontendUtils.MustHandleError(err)
}

func initProductClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	frontendUtils.MustHandleError(err)
	opts = append(opts, client.WithResolver(r))

	ProductClient, err = productcatalogservice.NewClient("product", opts...)
	frontendUtils.MustHandleError(err)

}

func initCartClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	frontendUtils.MustHandleError(err)
	opts = append(opts, client.WithResolver(r))

	CartClient, err = cartservice.NewClient("cart", opts...)
	frontendUtils.MustHandleError(err)

}
