package repo

import (
	"context"
	"fmt"
	"github.com/Yzc216/gomall/app/product/biz/model"
	"gorm.io/gorm"
)

type AttributeKeyRepository interface {
}

type AttributeKeyRepo struct {
	ctx context.Context
	db  *gorm.DB
}

func NewAttributeKeyRepo(ctx context.Context, db *gorm.DB) *AttributeKeyRepo {
	return &AttributeKeyRepo{
		ctx: ctx,
		db:  db,
	}
}

type AttributeKeyFilter struct {
	Name     string
	DataType string
	IsFilter *bool
}

// List 分页查询属性列表
func (r *AttributeKeyRepo) List(filter AttributeKeyFilter, p Pagination) ([]*model.AttributeKey, error) {
	var attrs []*model.AttributeKey
	query := r.db.WithContext(r.ctx).Model(&model.AttributeKey{})

	if filter.Name != "" {
		query = query.Where("name LIKE ?", "%"+filter.Name+"%")
	}
	if filter.DataType != "" {
		query = query.Where("data_type = ?", filter.DataType)
	}
	if filter.IsFilter != nil {
		query = query.Where("is_filter = ?", *filter.IsFilter)
	}

	err := query.Order("`order` ASC").
		Offset(p.Offset()).
		Limit(p.PageSize).
		Find(&attrs).Error

	return attrs, err
}

// Count 统计属性数量
func (r *AttributeKeyRepo) Count(filter AttributeKeyFilter) (int64, error) {
	var count int64
	query := r.db.WithContext(r.ctx).Model(&model.AttributeKey{})

	if filter.Name != "" {
		query = query.Where("name LIKE ?", "%"+filter.Name+"%")
	}
	if filter.DataType != "" {
		query = query.Where("data_type = ?", filter.DataType)
	}
	if filter.IsFilter != nil {
		query = query.Where("is_filter = ?", *filter.IsFilter)
	}

	err := query.Count(&count).Error
	return count, err
}

// GetByID 根据ID获取属性
func (r *AttributeKeyRepo) GetByID(id uint64) (*model.AttributeKey, error) {
	var attr model.AttributeKey
	err := r.db.WithContext(r.ctx).
		First(&attr, id).Error
	return &attr, err
}

// Create 创建属性
func (r *AttributeKeyRepo) Create(attr *model.AttributeKey) error {
	return r.db.WithContext(r.ctx).Transaction(func(tx *gorm.DB) error {
		// 唯一性校验
		var exist model.AttributeKey
		if err := tx.Where("name = ?", attr.Name).First(&exist).Error; err == nil {
			return fmt.Errorf("attribute name already exists: %s", attr.Name)
		}

		return tx.Create(attr).Error
	})
}

// Update 更新属性
func (r *AttributeKeyRepo) Update(attr *model.AttributeKey) error {
	return r.db.WithContext(r.ctx).Transaction(func(tx *gorm.DB) error {
		// 检查名称冲突
		var conflict model.AttributeKey
		if err := tx.Where("name = ? AND key_id != ?", attr.Name, attr.KeyID).
			First(&conflict).Error; err == nil {
			return fmt.Errorf("attribute name conflict: %s", attr.Name)
		}

		return tx.Save(attr).Error
	})
}

// Delete 删除属性
func (r *AttributeKeyRepo) Delete(id uint64) error {
	return r.db.WithContext(r.ctx).Transaction(func(tx *gorm.DB) error {
		// 检查关联属性值
		var count int64
		if err := tx.Model(&model.AttributeValue{}).
			Where("key_id = ?", id).
			Count(&count).Error; err != nil {
			return err
		}

		if count > 0 {
			return fmt.Errorf("cannot delete attribute with %d related values", count)
		}

		return tx.Delete(&model.AttributeKey{}, id).Error
	})
}
