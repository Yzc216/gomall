package service

import (
	"context"
	"github.com/Yzc216/gomall/app/cart/biz/dal/mysql"
	"github.com/Yzc216/gomall/app/cart/biz/model"
	"github.com/Yzc216/gomall/app/cart/infra/rpc"
	cart "github.com/Yzc216/gomall/app/cart/kitex_gen/cart"
	"github.com/Yzc216/gomall/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/kerrors"
)

type AddItemService struct {
	ctx context.Context
} // NewAddItemService new AddItemService
func NewAddItemService(ctx context.Context) *AddItemService {
	return &AddItemService{ctx: ctx}
}

// Run create note info
func (s *AddItemService) Run(req *cart.AddItemReq) (resp *cart.AddItemResp, err error) {
	getProduct, err := rpc.ProductClient.GetProduct(s.ctx, &product.GetProductReq{Id: req.Item.ProductId})
	if err != nil {
		return nil, err
	}

	if getProduct == nil || getProduct.Product.Id == 0 {
		return nil, kerrors.NewBizStatusError(400, "product not found")
	}

	cartItem := &model.Cart{
		UserId:    req.UserId,
		ProductId: req.Item.ProductId,
		Qty:       req.Item.Quantity,
	}
	err = model.AddItem(s.ctx, mysql.DB, cartItem)
	if err != nil {
		return nil, kerrors.NewBizStatusError(500, err.Error())
	}

	return &cart.AddItemResp{}, nil
}
