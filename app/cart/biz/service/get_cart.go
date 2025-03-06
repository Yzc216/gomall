package service

import (
	"context"
	"github.com/Yzc216/gomall/app/cart/biz/dal/mysql"
	"github.com/Yzc216/gomall/app/cart/biz/model"
	cart "github.com/Yzc216/gomall/rpc_gen/kitex_gen/cart"
	"github.com/cloudwego/kitex/pkg/kerrors"
)

type GetCartService struct {
	ctx context.Context
} // NewGetCartService new GetCartService
func NewGetCartService(ctx context.Context) *GetCartService {
	return &GetCartService{ctx: ctx}
}

// Run create note info
func (s *GetCartService) Run(req *cart.GetCartReq) (resp *cart.GetCartResp, err error) {
	// Finish your business logic.
	cartList, err := model.GetCartByUserId(s.ctx, mysql.DB, req.UserId)
	if err != nil {
		return nil, kerrors.NewBizStatusError(5000, err.Error())
	}

	var cartItems []*cart.CartItem
	for _, v := range cartList {
		cartItems = append(cartItems, &cart.CartItem{
			ProductId: v.ProductId,
			Quantity:  v.Qty,
		})
	}
	return &cart.GetCartResp{Cart: &cart.Cart{UserId: req.GetUserId(), Items: cartItems}}, nil
}
