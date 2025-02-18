package model

import (
	"context"
	"gorm.io/gorm"
	"time"
)

type SPU struct {
	ID          uint64   `gorm:"primaryKey"`                 // SPU ID（雪花算法生成）
	Title       string   `gorm:"type:varchar(200);not null"` // 商品标题
	SubTitle    string   `gorm:"type:varchar(200)"`          // 副标题
	ShopID      uint64   `gorm:"type:bigint unsigned"`       // 店铺id
	BrandID     uint64   `gorm:"type:bigint unsigned"`       // 关联品牌，多对一
	MainImages  []string `gorm:"type:varchar(200)"`          // 主图URL列表
	Video       string   `gorm:"type:varchar(200)"`          // 商品视频
	Description string   `gorm:"type:text"`                  // 商品详情
	Status      int      `gorm:"default:0"`                  // 状态（0-下架 1-上架 2-待审核）
	IsDeleted   bool     `gorm:"default:false"`              // 软删除标记
	CreatedAt   time.Time
	UpdatedAt   time.Time

	SKUs       []SKU      `gorm:"foreignKey:SpuID;references:ID"`            // 关联SKU，一对多
	Categories []Category `json:"categories" gorm:"many2many:SPU_category;"` // 关联分类，多对多
}

func (SPU) TableName() string {
	return "spu"
}

type SPUQuery struct {
	ctx context.Context
	db  *gorm.DB
}

func NewSPUQuery(ctx context.Context, db *gorm.DB) *SPUQuery {
	return &SPUQuery{
		ctx: ctx,
		db:  db,
	}
}

func (q SPUQuery) GetById(spuId int) (spu SPU, err error) {
	err = q.db.WithContext(q.ctx).Model(&SPU{}).First(&SPU{}, spuId).Error
	return
}

func (q SPUQuery) SearchProducts(query string) (spus []*SPU, err error) {
	err = q.db.WithContext(q.ctx).Model(&SPU{}).Find(&spus, "title like ? or description like ?", "%"+query+"%", "%"+query+"%").Error
	return
}
