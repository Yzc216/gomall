package model

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/Yzc216/gomall/app/inventory/types"
	"github.com/Yzc216/gomall/app/inventory/util"
	"github.com/cloudwego/kitex/tool/internal_pkg/log"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"strings"
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

	result := db.WithContext(ctx).Model(&Inventory{}).Create(&inventory)

	if result.Error != nil {
		// 直接判断是否为唯一冲突错误
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return types.ErrInvalidSKU
		}
		return fmt.Errorf("数据库错误: %v", result.Error)
	}

	journal := &InventoryJournal{
		SkuID:  inventory.SkuID,
		OpType: INIT,
		Delta:  int32(inventory.Total),
	}
	if err := db.WithContext(ctx).Model(&InventoryJournal{}).Create(journal).Error; err != nil {
		return err
	}

	return nil
}

func BatchInitInventory(db *sql.DB, inventories []*Inventory) error {
	if len(inventories) == 0 {
		return nil
	}

	// 开启事务
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("开启事务失败: %v", err)
	}
	defer tx.Rollback() // 确保失败时回滚

	// 1. 检查所有SKU是否已存在
	skuIDs := make([]interface{}, len(inventories))
	for i, inv := range inventories {
		skuIDs[i] = inv.SkuID
	}

	// 构建动态IN查询（解决变长参数问题）
	query := "SELECT sku_id FROM inventory WHERE sku_id IN (?" + strings.Repeat(",?", len(skuIDs)-1) + ")"
	rows, err := tx.Query(query, skuIDs...)
	if err != nil {
		return fmt.Errorf("查询存在的SKU失败: %v", err)
	}
	defer rows.Close()

	// 收集已存在的SKU
	existingSKUs := make(map[string]bool)
	for rows.Next() {
		var sku string
		if err := rows.Scan(&sku); err != nil {
			return fmt.Errorf("读取SKU失败: %v", err)
		}
		existingSKUs[sku] = true
	}
	if err := rows.Err(); err != nil {
		return fmt.Errorf("查询遍历失败: %v", err)
	}

	// 若存在重复SKU，返回错误
	if len(existingSKUs) > 0 {
		duplicates := make([]string, 0, len(existingSKUs))
		for sku := range existingSKUs {
			duplicates = append(duplicates, sku)
		}
		return fmt.Errorf("SKU已存在: %v", duplicates)
	}

	// 2. 批量插入新SKU
	var placeholders []string
	var values []interface{}
	for _, inv := range inventories {
		placeholders = append(placeholders, "(?, ?)") // 每个SKU占位符
		values = append(values, inv.SkuID, inv.Total)
	}

	// 构建批量插入SQL（如：INSERT INTO ... VALUES (?, ?), (?, ?)...）
	stmt := fmt.Sprintf(
		"INSERT INTO inventory (sku_id, stock) VALUES %s",
		strings.Join(placeholders, ","),
	)
	_, err = tx.Exec(stmt, values...)
	if err != nil {
		// 处理唯一约束冲突（防止并发插入导致冲突）
		if mysqlErr, ok := err.(*mysql.MySQLError); ok && mysqlErr.Number == 1062 {
			return fmt.Errorf("并发冲突: SKU已存在")
		}
		return fmt.Errorf("插入库存失败: %v", err)
	}

	// 提交事务
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("提交事务失败: %v", err)
	}

	return nil
}

// IncreaseStock 增加库存（如补货）
func IncreaseStock(ctx context.Context, db *gorm.DB, skuID uint64, delta uint32) error {
	if delta == 0 {
		return nil // 无需操作
	}

	return db.Transaction(func(tx *gorm.DB) error {
		// 1. 查询当前库存及版本号
		var inv Inventory
		if err := tx.WithContext(ctx).Select("version", "available", "total").
			Where("sku_id = ?", skuID).First(&inv).Error; err != nil {
			return fmt.Errorf("查询库存失败: %w", err)
		}

		// 2. 构建更新条件（原子性更新）
		result := tx.WithContext(ctx).Model(&Inventory{}).
			Where("sku_id = ? AND version = ?", skuID, inv.Version).
			Updates(map[string]interface{}{
				"total":     gorm.Expr("total + ?", delta),
				"available": gorm.Expr("available + ?", delta),
				"version":   inv.Version + 1,
			})

		if result.Error != nil {
			return fmt.Errorf("更新库存失败: %w", result.Error)
		}

		if result.RowsAffected == 0 {
			return types.ErrConcurrentModification
		}

		return nil
	})
}

// DecreaseStock 减少库存（如人为调减库存，需业务层确保合理性）
// allowOversell: 是否允许超卖（Available可减至负数）
func DecreaseStock(ctx context.Context, db *gorm.DB, skuID uint64, delta uint32, allowOversell bool) error {
	if delta == 0 {
		return nil
	}

	return db.Transaction(func(tx *gorm.DB) error {
		// 1. 查询当前库存及版本号
		var inv Inventory
		if err := tx.WithContext(ctx).Select("version", "available", "total").
			Where("sku_id = ?", skuID).First(&inv).Error; err != nil {
			return fmt.Errorf("查询库存失败: %w", err)
		}

		// 2. 检查可用库存是否足够（若不允超卖）
		if !allowOversell && inv.Available < int32(delta) {
			return fmt.Errorf("可用库存不足，当前可用: %d", inv.Available)
		}

		// 3. 原子性更新
		updates := map[string]interface{}{
			"total":     gorm.Expr("total - ?", delta),
			"available": gorm.Expr("available - ?", delta),
			"version":   inv.Version + 1,
		}

		result := tx.WithContext(ctx).Model(&Inventory{}).
			Where("sku_id = ? AND version = ?", skuID, inv.Version).
			Updates(updates)

		if result.Error != nil {
			return fmt.Errorf("更新库存失败: %w", result.Error)
		}

		if result.RowsAffected == 0 {
			return fmt.Errorf("库存版本冲突，请重试")
		}

		return nil
	})
}

// 查询库存
func GetStock(ctx context.Context, db *gorm.DB, skuIDs []uint64) ([]*Inventory, error) {
	var inventories []*Inventory

	// 执行批量查询，使用 IN 条件
	err := db.WithContext(ctx).
		Model(&Inventory{}).          // 指定模型
		Where("sku_id IN ?", skuIDs). // IN 查询条件
		Find(&inventories).           // 结果写入 inventories
		Error

	if err != nil {
		return nil, fmt.Errorf("failed to batch get stock: %w", err)
	}

	return inventories, nil
}

// 预占库存
func ReserveStockWithOptimistic(ctx context.Context, db *gorm.DB, skuID uint64, orderID string, quantity int32, allowOversell bool) error {
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
		if !allowOversell && inv.Available < quantity {
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

func ReserveStockWithLock(ctx context.Context, db *gorm.DB, skuID uint64, orderID string, quantity int32, allowOversell bool) error {
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

		if !allowOversell && inv.Available < quantity {
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
		if err = tx.WithContext(ctx).Model(&InventoryJournal{}).Create(journal).Error; err != nil {
			return err
		}

		return nil
	})
}

// 确认库存
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

		// 2. 检查库存是否足够（超卖需在此处理）
		if inv.Total < uint32(quantity) {
			// 触发超卖补偿逻辑（如异步通知运营）
			log.Warnf("库存不足，需补货: SKU=%d, 需要=%d, 实际=%d", skuID, quantity, inv.Total)
		}
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

		// 2. 检查库存是否足够（超卖需在此处理）
		if inv.Total < uint32(quantity) {
			// 触发超卖补偿逻辑（如异步通知运营）
			log.Warnf("库存不足，需补货: SKU=%d, 需要=%d, 实际=%d", skuID, quantity, inv.Total)
		}
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
		if err = tx.WithContext(ctx).Model(&InventoryJournal{}).Create(journal).Error; err != nil {
			return err
		}

		return nil
	})
}

// 释放库存
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

func cond(b bool, t, f int) int {
	if b {
		return t
	}
	return f
}
