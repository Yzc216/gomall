package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/Yzc216/gomall/app/checkout/infra/mq"
	"github.com/Yzc216/gomall/app/checkout/infra/rpc"
	"github.com/Yzc216/gomall/rpc_gen/kitex_gen/cart"
	checkout "github.com/Yzc216/gomall/rpc_gen/kitex_gen/checkout"
	"github.com/Yzc216/gomall/rpc_gen/kitex_gen/email"
	"github.com/Yzc216/gomall/rpc_gen/kitex_gen/order"
	"github.com/Yzc216/gomall/rpc_gen/kitex_gen/payment"
	"github.com/Yzc216/gomall/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/nats-io/nats.go"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"google.golang.org/protobuf/proto"
)

type CheckoutService struct {
	ctx context.Context
} // NewCheckoutService new CheckoutService
func NewCheckoutService(ctx context.Context) *CheckoutService {
	return &CheckoutService{ctx: ctx}
}

/*
Run
1. get cart
2. calculate cart
3. create order
4. empty cart
5. pay
6. change order result
7. finish
*/
func (s *CheckoutService) Run(req *checkout.CheckoutReq) (resp *checkout.CheckoutResp, err error) {
	// get cart
	cartResult, err := rpc.CartClient.GetCart(s.ctx, &cart.GetCartReq{UserId: req.UserId})
	if err != nil {
		klog.Error(err)
		err = fmt.Errorf("GetCart.err:%v", err)
		return
	}
	if cartResult == nil || cartResult.Cart == nil || len(cartResult.Cart.Items) == 0 {
		err = errors.New("cart is empty")
		return
	}

	// calculate cart
	var (
		oi    []*order.OrderItem
		total float64
	)
	spuIDs := make([]uint64, 0, len(cartResult.Cart.Items))
	for _, item := range cartResult.Cart.Items {
		spuIDs = append(spuIDs, item.SpuId)
	}
	productRes, err := rpc.ProductClient.BatchGetProducts(s.ctx, &product.BatchGetProductsReq{Ids: spuIDs})
	if err != nil {
		return nil, fmt.Errorf("批量获取商品失败: %v", err)
	}
	spuMap := productRes.Products

	for _, cartItem := range cartResult.Cart.Items {
		spu, exists := spuMap[cartItem.SpuId]
		if !exists {
			return nil, fmt.Errorf("商品信息未找到: spu_id=%d", cartItem.SpuId)
		}

		targetSku := findSkuByID(spu.Skus, cartItem.SkuId)
		if targetSku == nil {
			return nil, fmt.Errorf("规格信息未找到: spu_id=%d, sku_id=%d", cartItem.SpuId, cartItem.SkuId)
		}

		cost := targetSku.Price * float64(cartItem.Quantity)
		total += cost
		oi = append(oi, &order.OrderItem{
			Item: &cart.CartItem{SpuId: cartItem.SpuId, SkuId: cartItem.SkuId, Quantity: cartItem.Quantity},
			Cost: cost,
		})
	}

	// create order
	orderReq := &order.PlaceOrderReq{
		UserId:       req.UserId,
		UserCurrency: "USD",
		Items:        oi,
		Email:        req.Email,
	}
	if req.Address != nil {
		addr := req.Address
		orderReq.Address = &order.Address{
			StreetAddress: addr.StreetAddress,
			City:          addr.City,
			Country:       addr.Country,
			State:         addr.State,
			ZipCode:       addr.ZipCode,
		}
	}
	orderResult, err := rpc.OrderClient.PlaceOrder(s.ctx, orderReq)
	if err != nil {
		err = fmt.Errorf("PlaceOrder.err:%v", err)
		return
	}
	klog.Info("orderResult: ", orderResult)

	// empty cart
	emptyResult, err := rpc.CartClient.EmptyCart(s.ctx, &cart.EmptyCartReq{UserId: req.UserId})
	if err != nil {
		err = fmt.Errorf("EmptyCart.err:%v", err)
		return
	}
	klog.Info(emptyResult)

	// charge
	var orderId uint64
	if orderResult != nil || orderResult.Order != nil {
		orderId = orderResult.Order.OrderId
	}
	payReq := &payment.ChargeReq{
		UserId:  req.UserId,
		OrderId: orderId,
		Amount:  total,
		CreditCard: &payment.CreditCardInfo{
			CreditCardNumber:          req.CreditCard.CreditCardNumber,
			CreditCardExpirationYear:  req.CreditCard.CreditCardExpirationYear,
			CreditCardExpirationMonth: req.CreditCard.CreditCardExpirationMonth,
			CreditCardCvv:             req.CreditCard.CreditCardCvv,
		},
	}
	paymentResult, err := rpc.PaymentClient.Charge(s.ctx, payReq)
	if err != nil {
		err = fmt.Errorf("Charge.err:%v", err)
		return
	}
	data, _ := proto.Marshal(&email.EmailReq{
		From:        "from@example.com",
		To:          req.Email,
		ContentType: "text/plain",
		Subject:     "You just created an order in CloudWeGo shop",
		Content:     "You just created an order in CloudWeGo shop",
	})
	msg := &nats.Msg{Subject: "email", Data: data, Header: make(nats.Header)}

	// otel inject
	otel.GetTextMapPropagator().Inject(s.ctx, propagation.HeaderCarrier(msg.Header))

	_ = mq.Nc.PublishMsg(msg)

	klog.Info(paymentResult)

	// change order state
	klog.Info(orderResult)
	res, err := rpc.OrderClient.UpdateOrderState(s.ctx, &order.UpdateOrderStateReq{
		UserId:  req.UserId,
		OrderId: orderId,
		State:   order.OrderState_OrderStatePaid,
	})
	if err != nil {
		klog.Error(err)
		return
	}
	if !res.Success {
		klog.Info(res.Error)
	}

	resp = &checkout.CheckoutResp{
		OrderId:       orderId,
		TransactionId: paymentResult.TransactionId,
	}
	return
}

func findSkuByID(skus []*product.SKU, skuID uint64) *product.SKU {
	for _, sku := range skus {
		if sku.Id == skuID {
			return sku
		}
	}
	return nil
}
