package model

import (
	"context"
	"gorm.io/gorm"
	"time"
)

type SKU struct {
	ID        uint64  `gorm:"primaryKey"`                  // SKU ID
	SpuID     uint64  `gorm:"index;not null"`              // 关联SPU ID
	SpuTitle  string  `gorm:"type:varchar(255)"`           // 关联SPU名称 冗余字段
	Title     string  `gorm:"type:varchar(100);not null"`  // SKU标题（如：iPhone 13 红色 128GB）
	Price     float64 `gorm:"type:decimal(10,2);not null"` // 价格（单位：元）
	Stock     uint32  `gorm:"type:int unsigned"`           // 库存 冗余字段
	Sales     uint32  `gorm:"default:0"`                   // 销量
	IsActive  bool    `gorm:"default:true"`                // 是否可售
	CreatedAt time.Time
	UpdatedAt time.Time

	Specs []Attr `gorm:"foreignKey:SkuID;references:ID"` // 规格属性（如 {"color":"红", "storage":"128GB"}）
}

func (SKU) TableName() string {
	return "sku"
}

type SKUQuery struct {
	ctx context.Context
	db  *gorm.DB
}

func NewSKUQuery(ctx context.Context, db *gorm.DB) *SKUQuery {
	return &SKUQuery{
		ctx: ctx,
		db:  db,
	}
}

func (q SKUQuery) GetById(skuId int) (sku SKU, err error) {
	err = q.db.WithContext(q.ctx).Model(&SKU{}).First(&SKU{}, skuId).Error
	return
}

func (q SKUQuery) SearchProducts(query string) (skus []*SKU, err error) {
	err = q.db.WithContext(q.ctx).Model(&SKU{}).Find(&skus, "title like ? or spu_title like ?", "%"+query+"%", "%"+query+"%").Error
	return
}
