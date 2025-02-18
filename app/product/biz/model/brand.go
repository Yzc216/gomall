package model

type Brand struct {
	ID          uint64 `gorm:"primaryKey"`                   // 品牌id
	Name        string `gorm:"type:varchar(50);uniqueIndex"` // 品牌名
	Logo        []byte `gorm:"type:blob"`                    // logo图片
	Description string `gorm:"type:text"`                    // 品牌描述
	IsOfficial  bool   `gorm:"default:false"`                // 是否官方品牌

	SPUs []SPU `gorm:"foreignKey:BrandID;references:ID"`
}

func (Brand) TableName() string {
	return "brand"
}
