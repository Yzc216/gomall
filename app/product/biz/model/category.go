package model

import (
	"context"
	"gorm.io/gorm"
	"time"
)

type Category struct {
	ID          uint64    `gorm:"primaryKey;comment:分类ID"`
	Name        string    `gorm:"type:varchar(100);uniqueIndex:idx_name;comment:分类名称"`
	Description string    `gorm:"type:text;comment:分类描述"`
	ParentID    uint64    `gorm:"index:idx_parent_sort;comment:父分类ID（0表示根节点）"`
	Level       int8      `gorm:"index;type:tinyint;default:1;comment:层级（1-一级分类）"`
	IsLeaf      bool      `gorm:"default:true;comment:是否为叶子分类"`
	Sort        int       `gorm:"index:idx_parent_sort;default:1;comment:排序权重"`
	Image       string    `gorm:"type:varchar(500);comment:分类图标URL"`
	CreatedAt   time.Time `gorm:"comment:创建时间"`
	UpdatedAt   time.Time `gorm:"comment:更新时间"`

	SPUs []SPU `gorm:"many2many:spu_categories;joinForeignKey:category_id;joinReferences:spu_id;comment:关联SPU"`
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

func (c CategoryQuery) GetSubCategories(parentID uint64) ([]Category, error) {
	var categories []Category
	err := c.db.Where("parent_id = ?", parentID).Find(&categories).Error
	return categories, err
}
