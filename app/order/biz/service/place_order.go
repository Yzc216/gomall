package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/Yzc216/gomall/app/order/biz/dal/mysql"
	"github.com/Yzc216/gomall/app/order/biz/model"
	"github.com/Yzc216/gomall/app/order/infra/rpc"
	"github.com/Yzc216/gomall/common/utils"
	"github.com/Yzc216/gomall/rpc_gen/kitex_gen/inventory"
	order "github.com/Yzc216/gomall/rpc_gen/kitex_gen/order"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"gorm.io/gorm"
)

type PlaceOrderService struct {
	ctx context.Context
} // NewPlaceOrderService new PlaceOrderService
func NewPlaceOrderService(ctx context.Context) *PlaceOrderService {
	return &PlaceOrderService{ctx: ctx}
}

// Run create note info
func (s *PlaceOrderService) Run(req *order.PlaceOrderReq) (resp *order.PlaceOrderResp, err error) {
	if len(req.Items) == 0 {
		err = kerrors.NewBizStatusError(500001, "items is empty")
		return
	}
	var orderId uint64
	var stockItems = make([]*inventory.InventoryReq_Item, 0, len(req.Items))

	err = mysql.DB.Transaction(func(tx *gorm.DB) error {
		orderId = utils.GenID()

		o := &model.Order{
			OrderId:      orderId,
			UserId:       req.UserId,
			UserCurrency: req.UserCurrency,
			Consignee: model.Consignee{
				Email: req.Email,
			},
			OrderState: model.OrderStatePlaced,
		}
		if req.Address != nil {
			a := req.Address
			o.Consignee.StreetAddress = a.StreetAddress
			o.Consignee.City = a.City
			o.Consignee.State = a.State
			o.Consignee.Country = a.Country
		}
		if err := tx.Create(o).Error; err != nil {
			return err
		}

		var items []model.OrderItem

		for _, v := range req.Items {
			items = append(items, model.OrderItem{
				OrderIdRefer: orderId,
				SpuId:        v.Item.SpuId,
				SkuId:        v.Item.SkuId,
				Quantity:     v.Item.Quantity,
				Cost:         v.Cost,
			})
			stockItems = append(stockItems, &inventory.InventoryReq_Item{
				SkuId:    v.Item.SkuId,
				Quantity: v.Item.Quantity,
			})
		}
		if err := tx.Create(items).Error; err != nil {
			return err
		}

		return nil
	})

	fmt.Println(stockItems)
	//预占库存
	stockResp, err := rpc.InventoryClient.ReserveStock(s.ctx, &inventory.InventoryReq{
		OrderId: orderId,
		Items:   stockItems,
		Force:   false,
	})
	fmt.Println(stockResp)
	if err != nil || !stockResp.Success {
		// 补偿逻辑：标记订单为无效或触发回滚
		go s.rollbackOrder(orderId)
		return nil, errors.New("库存预留失败")
	}

	resp = &order.PlaceOrderResp{
		Order: &order.OrderResult{
			OrderId: orderId,
		},
	}

	return
}

func (s *PlaceOrderService) rollbackOrder(orderId uint64) {
	mysql.DB.Model(&model.Order{}).Where("order_id = ?", orderId).Update("order_state", model.OrderStateCancelled)
}
