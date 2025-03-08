package service

import (
	"context"
	"fmt"
	"github.com/Yzc216/gomall/app/frontend/infra/rpc"
	frontendUtils "github.com/Yzc216/gomall/app/frontend/utils"
	rpccart "github.com/Yzc216/gomall/rpc_gen/kitex_gen/cart"
	"github.com/Yzc216/gomall/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"strconv"

	common "github.com/Yzc216/gomall/app/frontend/hertz_gen/frontend/common"
	"github.com/cloudwego/hertz/pkg/app"
)

type GetCartService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetCartService(Context context.Context, RequestContext *app.RequestContext) *GetCartService {
	return &GetCartService{RequestContext: RequestContext, Context: Context}
}

func (h *GetCartService) Run(req *common.Empty) (resp map[string]any, err error) {
	cartResp, err := rpc.CartClient.GetCart(h.Context, &rpccart.GetCartReq{
		UserId: frontendUtils.GetUserIdFromCtx(h.Context),
	})
	if err != nil {
		return nil, err
	}

	spuIDs := make([]uint64, 0, len(cartResp.Cart.Items))
	for _, item := range cartResp.Cart.Items {
		spuIDs = append(spuIDs, item.SpuId)
	}
	productRes, err := rpc.ProductClient.BatchGetProducts(h.Context, &product.BatchGetProductsReq{Ids: spuIDs})
	if err != nil {
		return nil, fmt.Errorf("批量获取商品失败: %v", err)
	}
	spuMap := productRes.Products

	var items []map[string]any
	var total float64
	for _, item := range cartResp.Cart.Items {
		spu, exists := spuMap[item.SpuId]
		if !exists {
			return nil, fmt.Errorf("商品信息未找到: spu_id=%d", item.SpuId)
		}

		targetSku := findSkuByID(spu.Skus, item.SkuId)
		if targetSku == nil {
			return nil, fmt.Errorf("规格信息未找到: spu_id=%d, sku_id=%d", item.SpuId, item.SkuId)
		}

		items = append(items, map[string]any{
			"Name":        spu.BasicInfo.Title,
			"Price":       strconv.FormatFloat(targetSku.Price, 'f', 2, 64),
			"Specs":       targetSku.Specs, // 使用目标SKU的规格
			"Description": spu.BasicInfo.Description,
			"Picture":     getMainImage(spu.Media.MainImages),
			"Quantity":    item.Quantity,
		})
		total += targetSku.Price * float64(item.Quantity)
	}

	return utils.H{
		"title": "Cart",
		"items": items,
		"total": strconv.FormatFloat(total, 'f', 2, 64),
	}, nil
}

func findSkuByID(skus []*product.SKU, skuID uint64) *product.SKU {
	for _, sku := range skus {
		if sku.Id == skuID {
			return sku
		}
	}
	return nil
}

// 安全获取主图
func getMainImage(images []string) string {
	if len(images) > 0 {
		return images[0]
	}
	return ""
}
