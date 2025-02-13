package model

import (
	"encoding/json"
	"gorm.io/gorm"
)

const (
	RESERVE uint8 = iota + 1
	CONFIRM
	RELEASE
)

type InventoryJournal struct {
	gorm.Model
	SkuID    uint64          `gorm:"index:idx_sku_order"`
	BucketID string          `gorm:"type:varchar(32)"`
	OrderID  string          `gorm:"type:varchar(64);index:idx_sku_order"` // 联合索引
	OpType   uint8           `gorm:"comment:'1-预占 2-实际扣减 3-释放'"`
	Delta    int32           `gorm:"not null;default:0"`
	Before   json.RawMessage `gorm:"type:json;comment:'变更前状态'"` // 使用JSON类型
	After    json.RawMessage `gorm:"type:json;comment:'变更后状态'"`
}

func (InventoryJournal) TableName() string {
	return "inventory_journal"
}
