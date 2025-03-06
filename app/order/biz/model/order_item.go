package model

import "gorm.io/gorm"

type OrderItem struct {
	gorm.Model
	ProductId    uint64  `gorm:"type:bigint(11)"`
	OrderIdRefer uint64  `gorm:"type:bigint(11);index"`
	Quantity     uint32  `gorm:"type:int(11)"`
	Cost         float64 `gorm:"type:decimal(10,2)"`
}

func (OrderItem) TableName() string {
	return "order_item"
}
