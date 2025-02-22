package model

// 属性名表
type AttributeKey struct {
	KeyID    uint64 `gorm:"primaryKey;AUTO_INCREMENT;comment:属性名ID"`
	Name     string `gorm:"type:varchar(50);uniqueIndex;comment:属性名"`
	Unit     string `gorm:"type:varchar(20);comment:单位"`
	Order    int    `gorm:"index;comment:展示顺序"`
	DataType string `gorm:"type:varchar(20);comment:数据类型"`
	IsFilter bool   `gorm:"default:false;comment:是否参与筛选"`
}

// 属性值表
type AttributeValue struct {
	ID    int    `gorm:"primaryKey;column:id;AUTO_INCREMENT"`
	KeyID uint64 `gorm:"index:idx_key_sku;comment:属性名ID"`
	Value string `gorm:"type:varchar(100);comment:属性值"`
	SkuID uint64 `gorm:"index:idx_key_sku;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;comment:SKU ID"`
}

func (AttributeKey) TableName() string {
	return "attribute_key"
}

func (AttributeValue) TableName() string {
	return "attribute_value"
}
