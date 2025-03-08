package model

import (
	"context"
	"gorm.io/gorm"
)

type Consignee struct {
	Email         string
	StreetAddress string
	City          string
	State         string
	Country       string
	ZipCode       int32
}

type OrderState string

const (
	OrderStatePlaced   OrderState = "placed"
	OrderStatePaid     OrderState = "paid"
	OrderStateCanceled OrderState = "canceled"
)

type Order struct {
	gorm.Model
	OrderId      uint64 `gorm:"type:bigint(10);uniqueIndex;"`
	UserId       uint64 `gorm:"type:bigint(10);index;"`
	UserCurrency string

	Consignee  Consignee   `gorm:"embedded"`
	OrderItems []OrderItem `gorm:"foreignKey:OrderIdRefer;references:OrderId"`
	OrderState OrderState
}

func (Order) TableName() string {
	return "order"
}

func ListOrder(ctx context.Context, db *gorm.DB, userId uint64) (orders []Order, err error) {
	err = db.WithContext(ctx).Model(&Order{}).Where(&Order{UserId: userId}).Preload("OrderItems").Find(&orders).Error
	return
}

func GetOrder(ctx context.Context, db *gorm.DB, userId uint64, orderId uint64) (order Order, err error) {
	err = db.WithContext(ctx).Where(&Order{UserId: userId, OrderId: orderId}).First(&order).Error
	return
}

func UpdateOrderState(ctx context.Context, db *gorm.DB, userId uint64, orderId uint64, state OrderState) error {
	return db.WithContext(ctx).Model(&Order{}).Where(&Order{UserId: userId, OrderId: orderId}).Update("order_state", state).Error
}
