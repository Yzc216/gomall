package service

import (
	"context"
	"github.com/Yzc216/gomall/app/frontend/infra/rpc"
	frontendutils "github.com/Yzc216/gomall/app/frontend/utils"
	rpccart "github.com/Yzc216/gomall/rpc_gen/kitex_gen/cart"
	rpcproduct "github.com/Yzc216/gomall/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"strconv"

	common "github.com/Yzc216/gomall/app/frontend/hertz_gen/frontend/common"
	"github.com/cloudwego/hertz/pkg/app"
)

type CheckoutService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCheckoutService(Context context.Context, RequestContext *app.RequestContext) *CheckoutService {
	return &CheckoutService{RequestContext: RequestContext, Context: Context}
}

func (h *CheckoutService) Run(req *common.Empty) (resp map[string]any, err error) {
	var items []map[string]any
	userId := frontendutils.GetUserIdFromCtx(h.Context)

	carts, err := rpc.CartClient.GetCart(h.Context, &rpccart.GetCartReq{UserId: userId})
	if err != nil {
		return nil, err
	}
	var total float64

	for _, v := range carts.Cart.Items {
		productResp, err := rpc.ProductClient.GetProduct(h.Context, &rpcproduct.GetProductReq{
			Id: v.ProductId,
		})
		if err != nil {
			return nil, err
		}
		if productResp.Product == nil {
			continue
		}
		p := productResp.Product
		items = append(items, map[string]any{
			"Name":        p.BasicInfo.Title,
			"Price":       strconv.FormatFloat(p.Skus[0].Price, 'f', 2, 64),
			"Picture":     p.Media.MainImages[0],
			"Quantity":    strconv.Itoa(int(v.Quantity)),
			"Specs":       p.Skus[0].Specs,
			"Description": p.BasicInfo.Description,
		})
		total += float64(v.Quantity) * p.Skus[0].Price
	}

	return utils.H{
		"title": "Checkout",
		"items": items,
		"total": strconv.FormatFloat(total, 'f', 2, 64),
	}, nil
}
