package model

import "time"

type InventoryBucket struct {
	BucketID  string `gorm:"primaryKey;column:bucket_id;type:varchar(32)"` // 格式：sku_id:seq
	SkuID     uint64 `gorm:"index"`
	Total     uint32 `gorm:"not null"`
	Available int    `gorm:"not null"`
	Locked    uint32 `gorm:"default:0"`
	Version   uint32 `gorm:"default:0"`
	CreatedAt time.Time
}

func (InventoryBucket) TableName() string {
	return "inventory_buckets"
}
