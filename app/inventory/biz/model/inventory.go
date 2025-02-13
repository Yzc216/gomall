package model

import (
	"context"
	"fmt"
	"github.com/Yzc216/gomall/app/inventory/types"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

type Inventory struct {
	SkuID     uint64 `gorm:"primaryKey;column:sku_id"`
	Total     uint32 `gorm:"not null"`  // 总库存
	Available int32  `gorm:"not null"`  // 可用库存（可能为负用于超卖控制）
	Locked    uint32 `gorm:"default:0"` // 预占库存
	Version   uint32 `gorm:"default:0"` // 乐观锁版本
	CreatedAt time.Time
	UpdatedAt time.Time `gorm:"index"` // 更新时间加索引
}

func (Inventory) TableName() string {
	return "inventory"
}

func InitStock(ctx context.Context, db *gorm.DB, inventory *Inventory) error {
	return db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 使用原子性操作实现UPSERT
		result := tx.Clauses(clause.OnConflict{
			Columns: []clause.Column{{Name: "sku_id"}},
			DoUpdates: clause.Assignments(map[string]interface{}{
				"total":     inventory.Total,
				"available": inventory.Available,
				"locked":    0,
				"version":   gorm.Expr("version + 1"),
			}),
		}).Create(inventory)

		if result.Error != nil {
			return fmt.Errorf("库存操作失败: %w", result.Error)
		}
		return nil
	})
}

func BatchInitStock(ctx context.Context, db *gorm.DB, inventories []*Inventory) error {
	return db.WithContext(ctx).Clauses(clause.OnConflict{ // 冲突检测列
		Columns: []clause.Column{{Name: "sku_id"}},
		DoUpdates: clause.Assignments(map[string]interface{}{ // 冲突时更新操作
			"total":     gorm.Expr("VALUES(total)"),     // 使用新值覆盖
			"available": gorm.Expr("VALUES(available)"), // MySQL 语法需要 VALUES()
			"locked":    0,                              // 重置锁定库存
			"version":   gorm.Expr("version + 1"),       // 版本号自增
		}),
	}).Create(inventories).Error
}

func ReserveStockWithOptimistic(ctx context.Context, db *gorm.DB, skuID uint64, orderID string, quantity int32) error {
	return db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 1. 获取当前库存信息
		var inv Inventory
		//实现行级锁
		if err := tx.WithContext(ctx).Clauses(clause.Locking{Strength: "UPDATE"}).
			First(&inv, skuID).Error; err != nil {
			return err
		}

		// 2. 校验库存
		if inv.Available < quantity {
			return types.ErrAvailableStockInsufficient
		}

		// 3. 执行预占操作
		result := tx.WithContext(ctx).Model(&Inventory{}).
			Where("sku_id = ? AND version = ?", skuID, inv.Version).
			Updates(map[string]interface{}{
				"available": gorm.Expr("available - ?", quantity),
				"locked":    gorm.Expr("locked + ?", quantity),
				"version":   gorm.Expr("version + 1"),
			})

		if result.Error != nil {
			return fmt.Errorf("库存更新失败: %w", result.Error)
		}

		if result.RowsAffected == 0 {
			return types.ErrConcurrentModification
		}

		// 4. 记录流水 //TODO 消息队列异步记录
		journal := &InventoryJournal{
			SkuID:   skuID,
			OrderID: orderID,
			OpType:  RESERVE,
			Delta:   -quantity,
		}
		if err := tx.WithContext(ctx).Model(&InventoryJournal{}).Create(journal).Error; err != nil {
			return err
		}

		return nil
	})
}

func ConfirmStock(ctx context.Context, db *gorm.DB, skuID uint64, orderID string, quantity int32) error {
	return db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 1. 获取当前库存信息
		var inv Inventory
		if err := tx.WithContext(ctx).Model(&Inventory{}).First(&inv, skuID).Error; err != nil {
			return err
		}

		if inv.Locked < uint32(quantity) {
			return types.ErrLockedStockInsufficient
		}

		// 2. 执行原子化更新（双重校验）
		result := tx.WithContext(ctx).Model(&Inventory{}).
			Where("sku_id = ? AND locked >= ?", skuID, quantity).
			Updates(map[string]interface{}{
				"total":  gorm.Expr("total - ?", quantity),
				"locked": gorm.Expr("locked - ?", quantity),
			})

		if result.Error != nil {
			return fmt.Errorf("库存更新失败: %w", result.Error)
		}

		if result.RowsAffected == 0 {
			return types.ErrConcurrentModification
		}

		// 2. 记录确认流水
		journal := &InventoryJournal{
			SkuID:   skuID,
			OrderID: orderID,
			OpType:  CONFIRM,
			Delta:   -quantity,
		}
		if err := tx.WithContext(ctx).Model(&InventoryJournal{}).Create(journal).Error; err != nil {
			return err
		}

		return nil
	})
}

func ReleaseStock(ctx context.Context, db *gorm.DB, skuID uint64, orderID string, quantity int32, force bool) error {
	return db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var inv Inventory
		// 1. 获取当前库存信息（仅强制释放需要版本控制）
		if err := tx.WithContext(ctx).First(&inv, skuID).Error; err != nil {
			return err
		}
		if inv.Locked < uint32(quantity) {
			return types.ErrLockedStockInsufficient
		}

		// 2. 构建更新条件
		query := tx.WithContext(ctx).Model(&Inventory{}).
			Where("sku_id = ? AND locked >= ?", skuID, quantity)

		if force {
			query.WithContext(ctx).Where("version = ?", inv.Version)
		}

		// 3. 执行释放操作
		result := query.WithContext(ctx).Updates(map[string]interface{}{
			"available": gorm.Expr("available + ?", quantity),
			"locked":    gorm.Expr("locked - ?", quantity),
			"version":   gorm.Expr("version + ?", cond(force, 1, 0)),
		})

		if result.Error != nil {
			return fmt.Errorf("库存更新失败: %w", result.Error)
		}

		if force && result.RowsAffected == 0 {
			return types.ErrConcurrentModification
		}

		// 4. 记录流水
		journal := &InventoryJournal{
			SkuID:   skuID,
			OrderID: orderID,
			OpType:  RELEASE,
			Delta:   quantity,
		}
		if err := tx.WithContext(ctx).Model(&InventoryJournal{}).Create(journal).Error; err != nil {
			return err
		}

		return nil
	})
}

func cond(b bool, t, f int) int {
	if b {
		return t
	}
	return f
}
