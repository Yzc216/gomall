package model

import (
	"context"
	"errors"
	"fmt"
	"github.com/Yzc216/gomall/app/inventory/types"
	"github.com/Yzc216/gomall/app/inventory/util"
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
		if err := tx.WithContext(ctx).Model(&inv).First(&inv, skuID).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return types.ErrInvalidSKU
			}
			return err
		}

		// 2. 校验库存
		if inv.Available < quantity {
			return types.ErrAvailableStockInsufficient
		}

		// 3. 执行预占操作
		result := tx.WithContext(ctx).Model(&inv).
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

func ReserveStockWithLock(ctx context.Context, db *gorm.DB, skuID uint64, orderID string, quantity int32) error {
	// 1. 获取分布式锁
	mutex := util.GetLock(skuID)
	err := mutex.Lock()
	if err != nil {
		return err
	}
	defer mutex.Unlock()

	// 2. 执行库存操作
	return db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var inv Inventory
		if err = tx.WithContext(ctx).Model(&Inventory{}).First(&inv, skuID).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return types.ErrInvalidSKU
			}
			return err
		}

		if inv.Available < quantity {
			return types.ErrAvailableStockInsufficient
		}

		// 3. 执行预占操作
		result := tx.WithContext(ctx).Model(&inv).
			Where("sku_id = ?", skuID).
			Updates(map[string]interface{}{
				"available": gorm.Expr("available - ?", quantity),
				"locked":    gorm.Expr("locked + ?", quantity),
			})

		if result.Error != nil {
			return fmt.Errorf("库存更新失败: %w", result.Error)
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

func ConfirmStockWithOptimistic(ctx context.Context, db *gorm.DB, skuID uint64, orderID string, quantity int32) error {
	return db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 1. 获取当前库存信息
		var inv Inventory
		if err := tx.WithContext(ctx).Model(&Inventory{}).First(&inv, skuID).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return types.ErrInvalidSKU
			}
			return err
		}

		//2. 校验锁定库存
		if inv.Locked < uint32(quantity) {
			return types.ErrLockedStockInsufficient
		}

		// 3. 执行原子化更新
		result := tx.WithContext(ctx).Model(&Inventory{}).
			Where("sku_id = ? AND version = ?", skuID, inv.Version).
			Updates(map[string]interface{}{
				"total":   gorm.Expr("total - ?", quantity),
				"locked":  gorm.Expr("locked - ?", quantity),
				"version": gorm.Expr("version + 1"),
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

func ConfirmStockWithLock(ctx context.Context, db *gorm.DB, skuID uint64, orderID string, quantity int32) error {
	// 获取分布式锁
	mutex := util.GetLock(skuID)
	err := mutex.Lock()
	if err != nil {
		return err
	}
	defer mutex.Unlock()

	return db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 1. 获取当前库存信息
		var inv Inventory
		if err = tx.WithContext(ctx).Model(&Inventory{}).First(&inv, skuID).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return types.ErrInvalidSKU
			}
			return err
		}

		//2. 校验锁定库存
		if inv.Locked < uint32(quantity) {
			return types.ErrLockedStockInsufficient
		}

		// 3. 执行原子化更新
		result := tx.WithContext(ctx).Model(&Inventory{}).
			Where("sku_id = ?", skuID).
			Updates(map[string]interface{}{
				"total":  gorm.Expr("total - ?", quantity),
				"locked": gorm.Expr("locked - ?", quantity),
			})

		if result.Error != nil {
			return fmt.Errorf("库存更新失败: %w", result.Error)
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

func ReleaseStockWithOptimistic(ctx context.Context, db *gorm.DB, skuID uint64, orderID string, quantity int32, force bool) error {
	return db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var inv Inventory
		var total int32

		//当前订单预占库存数
		err := db.Model(&InventoryJournal{}).
			Select("SUM(delta)").
			Where("sku_id = ? AND order_id = ? AND op_type = ?", skuID, orderID, RESERVE).
			Scan(&total).Error
		if err != nil {
			return err
		}
		if total < quantity {

		}

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
			query.Where("version = ?", inv.Version)
		}

		// 3. 执行释放操作
		result := query.Updates(map[string]interface{}{
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

func ReleaseStockWithLock(ctx context.Context, db *gorm.DB, skuID uint64, orderID string, quantity int32, force bool) error {
	// 1. 获取分布式锁
	mutex := util.GetLock(skuID)
	err := mutex.Lock()
	if err != nil {
		return err
	}
	defer mutex.Unlock()

	return db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var inv Inventory

		////当前订单预占库存数
		//var total int32
		//err = db.Model(&InventoryJournal{}).
		//	Select("SUM(delta)").
		//	Where("sku_id = ? AND order_id = ? AND op_type = ?", skuID, orderID, RESERVE).
		//	Scan(&total).Error
		//if err != nil {
		//	return err
		//}
		//if total < quantity {
		//	fmt.Println("订单")
		//	return types.ErrLockedStockInsufficient
		//}

		// 2. 获取当前库存信息
		if err = tx.WithContext(ctx).Model(&Inventory{}).First(&inv, skuID).Error; err != nil {
			return err
		}
		if inv.Locked < uint32(quantity) {
			return types.ErrLockedStockInsufficient
		}

		// 3. 执行释放操作
		result := tx.WithContext(ctx).Model(&Inventory{}).
			Where("sku_id = ? AND locked >= ?", skuID, quantity).
			Updates(map[string]interface{}{
				"available": gorm.Expr("available + ?", quantity),
				"locked":    gorm.Expr("locked - ?", quantity),
			})

		if result.Error != nil {
			return fmt.Errorf("库存更新失败: %w", result.Error)
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
