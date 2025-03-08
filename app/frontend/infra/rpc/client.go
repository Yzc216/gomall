package rpc

import (
	"context"
	"github.com/Yzc216/gomall/app/frontend/conf"
	"github.com/Yzc216/gomall/app/frontend/infra/mtl"
	frontendUtils "github.com/Yzc216/gomall/app/frontend/utils"
	frontendutils "github.com/Yzc216/gomall/app/frontend/utils"
	"github.com/Yzc216/gomall/common/clientsuite"
	"github.com/Yzc216/gomall/rpc_gen/kitex_gen/cart/cartservice"
	"github.com/Yzc216/gomall/rpc_gen/kitex_gen/checkout/checkoutservice"
	"github.com/Yzc216/gomall/rpc_gen/kitex_gen/order/orderservice"
	"github.com/Yzc216/gomall/rpc_gen/kitex_gen/product"
	"github.com/Yzc216/gomall/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/Yzc216/gomall/rpc_gen/kitex_gen/user/userservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/circuitbreak"
	"github.com/cloudwego/kitex/pkg/fallback"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	prometheus "github.com/kitex-contrib/monitor-prometheus"
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
	var opts []client.Option

	// 1. 配置熔断器
	cbs := circuitbreak.NewCBSuite(func(ri rpcinfo.RPCInfo) string {
		return circuitbreak.RPCInfo2Key(ri)
	})
	cbs.UpdateServiceCBConfig("frontend/product/GetProduct", circuitbreak.CBConfig{Enable: true, ErrRate: 0.5, MinSample: 2})

	// 2. 配置降级策略
	opts = append(opts,
		commonSuite,
		client.WithCircuitBreaker(cbs),
		client.WithFallback(
			fallback.NewFallbackPolicy(
				fallback.UnwrapHelper(func(ctx context.Context, req, resp interface{}, err error) (fbResp interface{}, fbErr error) {
					methodName := rpcinfo.GetRPCInfo(ctx).To().Method()
					if err == nil {
						return resp, err
					}
					if methodName != "ListProducts" {
						return resp, err
					}
					return &product.ListProductsResp{
						Products: []*product.SPU{
							{
								BasicInfo: &product.SPUBasicInfo{
									Title:       "Apple iPhone 14 Pro 5G手机",
									SubTitle:    "Pro芯片 双卡双待",
									Description: "<p>旗舰智能手机，6.1英寸超视网膜XDR显示屏</p>",
									ShopId:      1001,
									Brand:       "Apple",
									Status:      1,
								},
								Media: &product.SPUMedia{
									MainImages: []string{
										"https://cdn.rentio.jp/matome/uploads/2022/09/cdfec7f90c26bf83efffeb1adc507599.png",
										"https://tse4-mm.cn.bing.net/th/id/OIP-C.1JtTljmKfBr-KphF4fi-vgHaHa?rs=1&pid=ImgDetMain",
									},
									VideoUrl: "https://cdn.example.com/iphone15_video.mp4",
								},
								CategoryRelation: &product.CategoryRelation{
									CategoryIds: []uint64{1, 2, 3}, // 手机 > 智能手机 > 高端智能手机
								},
								Skus: []*product.SKU{
									{
										Title: "iPhone 14 Pro 128GB 黑色",
										Price: 7999.00,
										Stock: 100,
										Specs: map[string]string{
											"颜色":   "黑色",
											"存储容量": "128GB",
										},
									},
									{
										Title: "iPhone 14 Pro 256GB 金色",
										Price: 8999.00,
										Stock: 50,
										Specs: map[string]string{
											"颜色":   "金色",
											"存储容量": "256GB",
										},
									},
								},
							},
						},
					}, nil
				},
				))))

	// 3. 配置 Prometheus 监控
	opts = append(opts, client.WithTracer(prometheus.NewClientTracer("", "", prometheus.WithDisableServer(true), prometheus.WithRegistry(mtl.Registry))))

	ProductClient, err = productcatalogservice.NewClient("product", opts...)
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
