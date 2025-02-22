package model

import (
	"time"
)

type SKU struct {
	ID        uint64    `gorm:"primaryKey;comment:SKU ID"`
	SpuID     uint64    `gorm:"index:idx_spu_active,priority:1;not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;comment:关联SPU ID"`
	SpuTitle  string    `gorm:"type:varchar(255);comment:关联SPU名称（冗余）"`
	Title     string    `gorm:"type:varchar(255);not null;comment:SKU标题"`
	Price     float64   `gorm:"index:idx_price;type:decimal(10,2);not null;comment:价格（单位：元）"`
	Stock     uint32    `gorm:"type:int unsigned;comment:库存"`
	Sales     uint32    `gorm:"default:0;index:idx_sales;comment:销量"`
	IsActive  bool      `gorm:"index:idx_spu_active,priority:2;default:true;comment:是否可售"`
	CreatedAt time.Time `gorm:"comment:创建时间"`
	UpdatedAt time.Time `gorm:"comment:更新时间"`

	Specs []AttributeValue `gorm:"foreignKey:SkuID;references:ID;constraint:OnDelete:CASCADE;comment:规格属性"` // 如 {"color":"红", "storage":"128GB"}
}

func (SKU) TableName() string {
	return "sku"
}
