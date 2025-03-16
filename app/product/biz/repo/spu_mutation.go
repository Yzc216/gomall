package repo

import (
	"context"
	"errors"
	"fmt"
	"github.com/Yzc216/gomall/app/product/biz/model"
	"github.com/Yzc216/gomall/app/product/biz/types"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type SPUMutationRepository interface {
	Create(ctx context.Context, spu *model.SPU) error
	Update(ctx context.Context, spu *model.SPU) error
	Delete(ctx context.Context, id uint64) error

	// 关联数据操作
	AddSKUs(ctx context.Context, spuID uint64, skuIDs []uint64) error
	RemoveSKUs(ctx context.Context, spuID uint64, skuIDs []uint64) error
	AddCategory(ctx context.Context, spuID, categoryID uint64) error
	RemoveCategories(ctx context.Context, spuID uint64, categoryIDs []uint64) error
	ReplaceCategories(ctx context.Context, spuID uint64, categoryIDs []uint64) error

	// 多媒体管理
	UpdateMedia(ctx context.Context, spuID uint64, media Media) error

	// 状态管理
	UpdateStatus(ctx context.Context, spuID uint64, status int8) error

	//BatchCreateSPUs(ctx context.Context, spus []*model.SPU) error
	//BatchUpdateStatus(ctx context.Context, spuIDs []uint64, status int8) error
}

type SPUMutation struct {
	db *gorm.DB
}

func NewSPUMutation(db *gorm.DB) *SPUMutation {
	return &SPUMutation{db: db}
}

func (m *SPUMutation) Create(ctx context.Context, spu *model.SPU) error {
	return m.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 1. 创建SPU基础信息
		if err := tx.Create(spu).Error; err != nil {
			// 捕获唯一键冲突错误
			if isDuplicateKeyError(err) {
				return types.ErrSPUTitleExists
			}
			return fmt.Errorf("创建SPU失败: %w", err)
		}

		// 2. 处理分类关联（多对多）
		if len(spu.Categories) > 0 {
			// 使用关联模式创建关系
			if err := tx.Model(&spu).Association("Categories").Append(spu.Categories); err != nil {
				return fmt.Errorf("创建分类关联失败: %w", err)
			}
		}

		return nil
	})
}

func (m *SPUMutation) Update(ctx context.Context, spu *model.SPU) error {
	return m.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 1. 检查记录是否存在并加锁（防止并发修改）
		var current model.SPU
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
			First(&current, spu.ID).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return types.ErrSPUNotFound
			}
			return fmt.Errorf("查询SPU失败: %w", err)
		}

		// 2. 直接更新记录（依赖唯一索引拦截标题重复）
		updateData := map[string]interface{}{
			"title":       spu.Title,
			"sub_title":   spu.SubTitle,
			"brand":       spu.Brand,
			"description": spu.Description,
			// 其他需要更新的字段...
		}
		result := tx.Model(spu).Where("id = ? ", spu.ID).Updates(updateData)
		if result.Error != nil {
			// 捕获唯一键冲突错误
			if isDuplicateKeyError(result.Error) {
				return types.ErrSPUTitleExists
			}
			return fmt.Errorf("更新SPU失败: %w", result.Error)
		}

		return nil
	})
}

func (m *SPUMutation) Delete(ctx context.Context, spuID uint64) error {
	return m.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 1. 清除 SPU 与分类的关联关系（不删除分类本身）
		if err := tx.Exec(
			"DELETE FROM spu_categories WHERE spu_id = ?", spuID,
		).Error; err != nil {
			return fmt.Errorf("failed to clear categories association: %w", err)
		}

		// 先删除所有关联 SKU（物理删除）
		if err := tx.Where("spu_id = ?", spuID).
			Delete(&model.SKU{}).Error; err != nil {
			return fmt.Errorf("failed to delete SKUs: %w", err)
		}

		// 3. 删除 SPU（软删除或物理删除根据业务需求）
		// 此处示例为软删除（使用 gorm.DeletedAt）
		if err := tx.Delete(&model.SPU{ID: spuID}).Error; err != nil {
			return fmt.Errorf("failed to delete SPU: %w", err)
		}

		return nil
	})
}

func (m *SPUMutation) AddSKUs(ctx context.Context, spuID uint64, skuIDs []uint64) error {
	// 验证SKU存在性
	var count int64
	if err := m.db.WithContext(ctx).
		Model(&model.SKU{}).
		Where("id IN ?", skuIDs).
		Count(&count).Error; err != nil {
		return err
	}
	if int(count) != len(skuIDs) {
		return errors.New("invalid SKU IDs")
	}

	return m.db.WithContext(ctx).Model(&model.SPU{ID: spuID}).
		Association("SKUs").
		Append(generateSKUReferences(skuIDs))
}

func (m *SPUMutation) RemoveSKUs(ctx context.Context, spuID uint64, skuIDs []uint64) error {
	return m.db.WithContext(ctx).Model(&model.SPU{ID: spuID}).
		Association("SKUs").
		Delete(generateSKUReferences(skuIDs))
}

func (m *SPUMutation) AddCategory(ctx context.Context, spuID, categoryID uint64) error {
	return m.db.WithContext(ctx).Model(&model.SPU{ID: spuID}).
		Association("Categories").
		Append(&model.Category{ID: categoryID})
}

func (m *SPUMutation) RemoveCategories(ctx context.Context, spuID uint64, categoryIDs []uint64) error {
	refs := make([]*model.Category, len(categoryIDs))
	for i, id := range categoryIDs {
		refs[i] = &model.Category{ID: id}
	}
	return m.db.WithContext(ctx).Model(&model.SPU{ID: spuID}).
		Association("Categories").
		Delete(refs)
}

func (m *SPUMutation) ReplaceCategories(ctx context.Context, spuID uint64, categoryIDs []uint64) error {
	categories := make([]*model.Category, len(categoryIDs))
	for i, id := range categoryIDs {
		categories[i] = &model.Category{ID: id}
	}
	return m.db.WithContext(ctx).Model(&model.SPU{ID: spuID}).
		Association("Categories").
		Replace(categories)
}

// 多媒体管理
func (m *SPUMutation) UpdateMedia(ctx context.Context, spuID uint64, media Media) error {
	return m.db.WithContext(ctx).Model(&model.SPU{}).
		Where("id = ?", spuID).
		Updates(map[string]interface{}{
			"main_images": media.MainImages,
			"video":       media.Video,
		}).Error
}

// 状态管理
func (m *SPUMutation) UpdateStatus(ctx context.Context, spuID uint64, status int8) error {
	return m.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var current model.SPU
		if err := tx.First(&current, spuID).Error; err != nil {
			return types.ErrSPUNotFound
		}

		// 校验状态合法性
		if !isValidStatusTransition(current.Status, status) {
			return types.ErrInvalidStatus
		}

		return tx.Model(&current).Update("status", status).Error
	})
}

func generateSKUReferences(ids []uint64) []*model.SKU {
	skus := make([]*model.SKU, len(ids))
	for i, id := range ids {
		skus[i] = &model.SKU{ID: id}
	}
	return skus
}

func isValidStatusTransition(oldStatus, newStatus int8) bool {
	// 状态机校验逻辑
	transitions := map[int8][]int8{
		0: {1},    // 草稿 -> 待审核
		1: {2, 3}, // 待审核 -> 审核通过/拒绝
		2: {4, 5}, // 审核通过 -> 上架/下架
		3: {0},    // 审核拒绝 -> 草稿
		4: {5},    // 上架 -> 下架
		5: {4, 6}, // 下架 -> 上架/删除
	}
	allowed, ok := transitions[oldStatus]
	if !ok {
		return false
	}
	for _, s := range allowed {
		if s == newStatus {
			return true
		}
	}
	return false
}

func isDuplicateKeyError(err error) bool {
	var dbErr *mysql.MySQLError
	if errors.As(errors.Unwrap(err), &dbErr) {
		return dbErr.Number == 1062
	}
	return false
}
