package service

import (
	"context"
	"github.com/Yzc216/gomall/app/frontend/infra/rpc"
	frontendUtils "github.com/Yzc216/gomall/app/frontend/utils"

	cart "github.com/Yzc216/gomall/app/frontend/hertz_gen/frontend/cart"
	rpccart "github.com/Yzc216/gomall/rpc_gen/kitex_gen/cart"
	"github.com/cloudwego/hertz/pkg/app"
)

type AddCartItemService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewAddCartItemService(Context context.Context, RequestContext *app.RequestContext) *AddCartItemService {
	return &AddCartItemService{RequestContext: RequestContext, Context: Context}
}

func (h *AddCartItemService) Run(req *cart.AddCartItemReq) (resp map[string]any, err error) {

	_, err = rpc.CartClient.AddItem(h.Context, &rpccart.AddItemReq{
		UserId: uint32(frontendUtils.GetUserIdFromCtx(h.Context)),
		Item: &rpccart.CartItem{
			ProductId: req.ProductId,
			Quantity:  uint32(req.ProductNum),
		},
	})
	if err != nil {
		return nil, err
	}
	return
}
