package rpc

import (
	"github.com/Yzc216/gomall/app/frontend/conf"
	frontendUtils "github.com/Yzc216/gomall/app/frontend/utils"
	frontendutils "github.com/Yzc216/gomall/app/frontend/utils"
	"github.com/Yzc216/gomall/common/clientsuite"
	"github.com/Yzc216/gomall/rpc_gen/kitex_gen/cart/cartservice"
	"github.com/Yzc216/gomall/rpc_gen/kitex_gen/checkout/checkoutservice"
	"github.com/Yzc216/gomall/rpc_gen/kitex_gen/order/orderservice"
	"github.com/Yzc216/gomall/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/Yzc216/gomall/rpc_gen/kitex_gen/user/userservice"
	"github.com/cloudwego/kitex/client"
	"sync"
)

var (
	UserClient     userservice.Client
	ProductClient  productcatalogservice.Client
	CartClient     cartservice.Client
	CheckoutClient checkoutservice.Client
	OrderClient    orderservice.Client
	Once           sync.Once
	err            error
	registryAddr   string
	commonSuite    client.Option
)

func InitClient() {
	Once.Do(func() {
		registryAddr = conf.GetConf().Hertz.RegistryAddr
		commonSuite = client.WithSuite(clientsuite.CommonGrpcClientSuite{
			RegistryAddr:       registryAddr,
			CurrentServiceName: frontendutils.ServiceName,
		})
		initUserClient()
		initProductClient()
		initCartClient()
		initCheckoutClient()
		initOrderClient()
	})
}

func initUserClient() {
	UserClient, err = userservice.NewClient("user", commonSuite)
	frontendUtils.MustHandleError(err)
}

func initProductClient() {
	//var opts []client.Option
	//
	//// 1. 配置熔断器
	//cbs := circuitbreak.NewCBSuite(func(ri rpcinfo.RPCInfo) string {
	//	return circuitbreak.RPCInfo2Key(ri)
	//})
	//cbs.UpdateServiceCBConfig("shop-frontend/product/GetProduct", circuitbreak.CBConfig{Enable: true, ErrRate: 0.5, MinSample: 2})
	//
	//// 2. 配置降级策略
	//opts = append(opts, commonSuite, client.WithCircuitBreaker(cbs), client.WithFallback(fallback.NewFallbackPolicy(fallback.UnwrapHelper(func(ctx context.Context, req, resp interface{}, err error) (fbResp interface{}, fbErr error) {
	//	methodName := rpcinfo.GetRPCInfo(ctx).To().Method()
	//	if err == nil {
	//		return resp, err
	//	}
	//	if methodName != "ListProducts" {
	//		return resp, err
	//	}
	//	return &product.ListProductsResp{
	//		Products: []*product.Product{
	//			{
	//				Price:       6.6,
	//				Id:          3,
	//				Picture:     "/static/image/t-shirt.jpeg",
	//				Name:        "T-Shirt",
	//				Description: "CloudWeGo T-Shirt",
	//			},
	//		},
	//	}, nil
	//}))))
	//
	//// 3. 配置 Prometheus 监控
	//opts = append(opts, client.WithTracer(prometheus.NewClientTracer("", "", prometheus.WithDisableServer(true), prometheus.WithRegistry(mtl.Registry))))
	ProductClient, err = productcatalogservice.NewClient("product", commonSuite)
	frontendUtils.MustHandleError(err)

}

func initCartClient() {
	CartClient, err = cartservice.NewClient("cart", commonSuite)
	frontendUtils.MustHandleError(err)
}

func initCheckoutClient() {
	CheckoutClient, err = checkoutservice.NewClient("checkout", commonSuite)
	frontendUtils.MustHandleError(err)
}

func initOrderClient() {
	OrderClient, err = orderservice.NewClient("order", commonSuite)
	frontendUtils.MustHandleError(err)
}
