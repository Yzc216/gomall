package model

import (
	"context"
	"gorm.io/gorm"
)

type Category struct {
	ID          uint64 `gorm:"primaryKey"`                   // 分类ID
	Name        string `gorm:"type:varchar(50);uniqueIndex"` // 种类名
	Description string `gorm:"type:text"`                    // 种类描述
	ParentID    uint64 `gorm:"index"`                        // 父分类ID（0表示根节点）
	Level       int    `gorm:"default:1"`                    // 层级（1-一级分类，2-二级分类）
	IsLeaf      bool   `gorm:"default:true"`                 // 是否为叶子分类
	Sort        int    `gorm:"default:1"`                    // 排序权重
	Image       []byte `gorm:"type:blob"`                    // 分类图标

	SPUs []SPU `gorm:"many2many:SPU_category;"` // 关联spu，多对多
}

func (Category) TableName() string {
	return "category"
}

type CategoryQuery struct {
	ctx context.Context
	db  *gorm.DB
}

func NewCategoryQuery(ctx context.Context, db *gorm.DB) *CategoryQuery {
	return &CategoryQuery{
		ctx: ctx,
		db:  db,
	}
}

func (c CategoryQuery) GetProductsByCategoryName(name string) (categories []Category, err error) {
	err = c.db.WithContext(c.ctx).Model(&Category{}).Where(&Category{Name: name}).Preload("Products").Find(&categories).Error
	return
}
