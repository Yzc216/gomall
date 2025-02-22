package repo

import (
	"context"
	"fmt"
	"github.com/Yzc216/gomall/app/product/biz/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type AttributeValueRepo struct {
	ctx context.Context
	db  *gorm.DB
}

func NewAttributeValueRepo(ctx context.Context, db *gorm.DB) *AttributeValueRepo {
	return &AttributeValueRepo{
		ctx: ctx,
		db:  db,
	}
}

// SaveBatch 批量保存属性值（UPSERT语义）
func (r *AttributeValueRepo) SaveBatch(values []model.AttributeValue) error {
	return r.db.WithContext(r.ctx).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "key_id"}, {Name: "sku_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"value"}),
	}).CreateInBatches(values, 500).Error
}

// GetIDsByKey 根据属性键获取值ID列表
func (r *AttributeValueRepo) GetIDsByKey(keyID uint32) ([]int, error) {
	var ids []int
	err := r.db.WithContext(r.ctx).Model(&model.AttributeValue{}).
		Where("key_id = ?", keyID).
		Pluck("id", &ids).Error
	return ids, err
}

// UpdateBatch 批量更新属性值（根据主键）
func (r *AttributeValueRepo) UpdateBatch(values []model.AttributeValue) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		caseStmt := "CASE id "
		params := make(map[string]interface{})
		ids := make([]int, 0, len(values))

		for i, v := range values {
			param := fmt.Sprintf("when %d then ?", v.ID)
			caseStmt += param
			params[fmt.Sprintf("value%d", i)] = v.Value
			ids = append(ids, v.ID)
		}
		caseStmt += " END"

		return tx.Model(&model.AttributeValue{}).
			Where("id IN ?", ids).
			Update("value", gorm.Expr(caseStmt, params)).
			Error
	})
}

// DeleteBatch 批量删除属性值
func (r *AttributeValueRepo) DeleteBatch(ids []int) error {
	return r.db.WithContext(r.ctx).Transaction(func(tx *gorm.DB) error {
		return tx.Where("id IN ?", ids).
			Delete(&model.AttributeValue{}).Error
	})
}
