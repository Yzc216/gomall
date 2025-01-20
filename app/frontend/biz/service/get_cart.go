package service

import (
	"context"
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
		UserId: uint32(frontendUtils.GetUserIdFromCtx(h.Context)),
	})
	if err != nil {
		return nil, err
	}

	var items []map[string]any
	var total float64
	for _, item := range cartResp.Items {
		productResp, err := rpc.ProductClient.GetProduct(h.Context, &product.GetProductReq{Id: item.ProductId})
		if err != nil {
			return nil, err
		}
		p := productResp.Product
		items = append(items, map[string]any{
			"Name":        p.Name,
			"Price":       p.Price,
			"Description": p.Description,
			"Picture":     p.Picture,
			"Qty":         item.Quantity,
		})
		total += float64(p.Price) * float64(item.Quantity)
	}

	return utils.H{
		"title": "Cart",
		"items": items,
		"total": strconv.FormatFloat(total, 'f', 2, 64),
	}, nil
}
