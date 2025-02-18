package model

type Attr struct {
	AttrID uint32 `gorm:"primaryKey"`
	SkuID  uint64 `gorm:"type:bigint unsigned"`
	Name   string `gorm:"type:varchar(50)"`
	Value  string `gorm:"type:varchar(100)"`
}

func (Attr) TableName() string {
	return "sku_attribute"
}
