package repo

import (
	"context"
	"fmt"
	"github.com/Yzc216/gomall/app/product/biz/model"
	"gorm.io/gorm"
)

// import (
//
//	"context"
//	"errors"
//	"fmt"
//	ordermodel "github.com/Yzc216/gomall/app/order/biz/model"
//	"github.com/Yzc216/gomall/app/product/biz/types"
//	"gorm.io/gorm"
//	"gorm.io/gorm/clause"
//	"maps"
//	"slices"
//	"strings"
//	"sync"
//	"time"
//
// )
var (
// _ SKURepository = (*SKURepo)(nil)
)

// sku_repository.go
type SKURepository interface {
	// 基础CRUD
	Create(ctx context.Context, sku *model.SKU) error
	Update(ctx context.Context, sku *model.SKU) error
	Delete(ctx context.Context, id uint64) error
	GetByID(ctx context.Context, id uint64) (*model.SKU, error)
	ExistsByID(ctx context.Context, id uint64) (bool, error)

	// 复杂查询
	List(ctx context.Context, filter SKUFilter, page Pagination) ([]*model.SKU, int64, error)

	// 批量操作
	BatchGetByIDs(ctx context.Context, ids []uint64) ([]*model.SKU, error)
	BatchCreate(ctx context.Context, spuID uint64, skus []*model.SKU) error
	BatchToggleActive(ctx context.Context, spuID uint64, active bool) error

	//// 库存管理
	//IncreaseStock(ctx context.Context, skuID uint64, quantity uint32) error
	//DecreaseStock(ctx context.Context, skuID uint64, quantity uint32) error
	//FreezeStock(ctx context.Context, skuID uint64, quantity uint32) error
	//GetStockInfo(ctx context.Context, spuID uint64) (totalStock uint32, availableStock uint32, err error)

	// 关联查询
	GetActiveSKUs(ctx context.Context, spuID uint64) ([]*model.SKU, error)
	GetSpecs(ctx context.Context, spuID uint64) (string, error)
	GetSKUsByAttributes(ctx context.Context, attrFilters map[string]string, page Pagination) ([]*model.SKU, error)

	// 销售统计
	UpdateSales(ctx context.Context, skuID uint64, quantity uint32) error
	GetTopSales(ctx context.Context, limit int) ([]*model.SKU, error)
}

type SKURepo struct {
	db *gorm.DB
}

func NewSKURepo(db *gorm.DB) *SKURepo {
	return &SKURepo{db: db}
}

type SKUFilter struct {
	SPUID    uint64
	MinPrice float64
	MaxPrice float64
	InStock  bool  // 是否有库存
	IsActive *bool // 是否上架
	HasSpecs bool  // 是否有规格属性
}

// 基础CRUD
// --------------------------------------------------
func (r *SKURepo) Create(ctx context.Context, sku *model.SKU) error {
	//// 检查同一SPU下标题是否重复
	//var exist int64
	//r.db.WithContext(ctx).Model(&SKU{}).
	//	Where("spu_id = ? AND title = ?", sku.SpuID, sku.Title).
	//	Count(&exist)
	//if exist > 0 {
	//	return types.ErrDuplicateSKUTitle
	//}

	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 获取SPU标题
		//var spu SPU
		//if err := tx.Select("title").First(&spu, sku.SpuID).Error; err != nil {
		//	return fmt.Errorf("get SPU failed: %w", err)
		//}
		//sku.SpuTitle = spu.Title

		if err := tx.Create(sku).Error; err != nil {
			return fmt.Errorf("create SKU failed: %w", err)
		}

		//// 处理规格属性
		//if len(sku.Specs) > 0 {
		//	if err := tx.Model(sku).Association("Specs").Append(sku.Specs); err != nil {
		//		return fmt.Errorf("add specs failed: %w", err)
		//	}
		//}
		return nil
	})
}

//func (r *SKURepo) Update(ctx context.Context, sku *SKU) error {
//	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
//		// 获取当前数据并加锁
//		var current SKU
//		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
//			First(&current, sku.ID).Error; err != nil {
//			return types.ErrSKUNotFound
//		}
//
//		// 校验标题唯一性
//		if current.Title != sku.Title {
//			var count int64
//			tx.Model(&SKU{}).
//				Where("spu_id = ? AND title = ? AND id != ?", current.SpuID, sku.Title, sku.ID).
//				Count(&count)
//			if count > 0 {
//				return types.ErrDuplicateSKUTitle
//			}
//		}
//
//		// 保留重要字段
//		updates := map[string]interface{}{
//			"title":     sku.Title,
//			"price":     sku.Price,
//			"is_active": sku.IsActive,
//		}
//		return tx.Model(sku).Updates(updates).Error
//	})
//}
//
//func (r *SKURepo) Delete(ctx context.Context, id uint64) error {
//	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
//		// 检查是否存在关联订单
//		//TODO
//		var orderCount int64
//		if err := tx.Model(&ordermodel.OrderItem{}).Where("sku_id = ?", id).Count(&orderCount).Error; err != nil {
//			return err
//		}
//		if orderCount > 0 {
//			return types.ErrSKUInUse
//		}
//
//		// 删除规格属性
//		if err := tx.Model(&SKU{ID: id}).Association("Specs").Clear(); err != nil {
//			return err
//		}
//
//		// 删除主记录
//		return tx.Delete(&SKU{}, id).Error
//	})
//}
//
//func (r *SKURepo) GetByID(ctx context.Context, id uint64) (*SKU, error) {
//	var sku SKU
//	err := r.db.WithContext(ctx).
//		Preload("Specs").
//		First(&sku, id).Error
//	if errors.Is(err, gorm.ErrRecordNotFound) {
//		return nil, types.ErrSKUNotFound
//	}
//	return &sku, err
//}
//
//func (r *SKURepo) ExistsByID(ctx context.Context, id uint64) (bool, error) {
//	var count int64
//	err := r.db.WithContext(ctx).
//		Model(&SKU{}).
//		Where("id = ?", id).
//		Count(&count).Error
//	return count > 0, err
//}
//
//// 复杂查询
//// --------------------------------------------------
//func (r *SKURepo) List(ctx context.Context, filter SKUFilter, page Pagination) ([]*SKU, int64, error) {
//	query := r.buildListQuery(filter)
//
//	var total int64
//	if err := query.Model(&SKU{}).Count(&total).Error; err != nil {
//		return nil, 0, err
//	}
//
//	var skus []*SKU
//	err := query.Scopes(Paginate(page)).
//		Preload("Specs").
//		Find(&skus).Error
//
//	return skus, total, err
//}
//
//// 批量操作
//// --------------------------------------------------
//func (r *SKURepo) BatchGetByIDs(ctx context.Context, ids []uint64) ([]*SKU, error) {
//	var skus []*SKU
//	err := r.db.WithContext(ctx).
//		Where("id IN ?", ids).
//		Find(&skus).Error
//	return skus, err
//}

func (r *SKURepo) BatchCreate(ctx context.Context, spuID uint64, skus []*model.SKU) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 批量创建
		if err := tx.CreateInBatches(skus, 100).Error; err != nil {
			return fmt.Errorf("batch create SKUs failed: %w", err)
		}
		return nil
	})
}

//func (r *SKURepo) BatchToggleActive(ctx context.Context, spuID uint64, active bool) error {
//	return r.db.WithContext(ctx).
//		Model(&SKU{}).
//		Where("spu_id = ?", spuID).
//		Update("is_active", active).Error
//}
//
//// 库存管理（建议添加到接口）
//// --------------------------------------------------
//func (r *SKURepo) IncreaseStock(ctx context.Context, skuID uint64, quantity uint32) error {
//	return r.db.WithContext(ctx).
//		Model(&SKU{}).
//		Where("id = ?", skuID).
//		Update("stock", gorm.Expr("stock + ?", quantity)).Error
//}
//
//func (r *SKURepo) DecreaseStock(ctx context.Context, skuID uint64, quantity uint32) error {
//	result := r.db.WithContext(ctx).
//		Model(&SKU{}).
//		Where("id = ? AND stock >= ?", skuID, quantity).
//		Update("stock", gorm.Expr("stock - ?", quantity))
//
//	if result.Error != nil {
//		return result.Error
//	}
//	if result.RowsAffected == 0 {
//		return types.ErrStockInsufficient
//	}
//	return nil
//}
//
//// 关联查询
//// --------------------------------------------------
//func (r *SKURepo) GetActiveSKUs(ctx context.Context, spuID uint64) ([]*SKU, error) {
//	var skus []*SKU
//	err := r.db.WithContext(ctx).
//		Where("spu_id = ? AND is_active = ?", spuID, true).
//		Find(&skus).Error
//	return skus, err
//}
//
//func (r *SKURepo) GetSpecs(ctx context.Context, skuID uint64) ([]*AttributeValue, error) {
//	var specs []*AttributeValue
//	err := r.db.WithContext(ctx).
//		Model(&SKU{ID: skuID}).
//		Association("Specs").
//		Find(&specs)
//	return specs, err
//}
//
//func (r *SKURepo) GetSKUsByAttributes(ctx context.Context, attrFilters map[string]string, page Pagination) ([]*SKU, error) {
//	// 1. 获取属性KeyID
//	var keyIDs []uint64
//	err := r.db.WithContext(ctx).Model(&AttributeKey{}).
//		Where("name IN (?)", maps.Keys(attrFilters)).
//		Pluck("key_id", &keyIDs).Error
//	if err != nil {
//		return nil, err
//	}
//
//	// 2. 构建子查询：获取满足所有属性的SkuID
//	subQuery := r.db.WithContext(ctx).Model(&AttributeValue{}).
//		Select("sku_id").
//		Where("(key_id, value) IN ?", getKeyValuePairs(keyIDs, attrFilters)).
//		Group("sku_id").
//		Having("COUNT(DISTINCT key_id) = ?", len(keyIDs))
//
//	// 3. 主查询
//	var skus []*SKU
//	err = r.db.WithContext(ctx).Preload("AttributeValues", func(db *gorm.DB) *gorm.DB {
//		return db.Joins("JOIN attribute_keys ON attribute_values.key_id = attribute_keys.key_id").
//			Order("attribute_keys.order ASC")
//	}).
//		Joins("JOIN (?) AS matched ON skus.id = matched.sku_id", subQuery).
//		Order("sales DESC").
//		Scopes(Paginate(page)).
//		Find(&skus).Error
//
//	return skus, err
//}
//
//// 销售统计
//// --------------------------------------------------
//func (r *SKURepo) UpdateSales(ctx context.Context, skuID uint64, quantity uint32) error {
//	return r.db.WithContext(ctx).
//		Model(&SKU{}).
//		Where("id = ?", skuID).
//		Update("sales", gorm.Expr("sales + ?", quantity)).Error
//}
//
//func (r *SKURepo) GetTopSales(ctx context.Context, limit int) ([]*SKU, error) {
//	var skus []*SKU
//	err := r.db.WithContext(ctx).
//		Order("sales DESC").
//		Limit(limit).
//		Find(&skus).Error
//	return skus, err
//}
//
//// 私有辅助方法
//// --------------------------------------------------
//func (r *SKURepo) buildListQuery(filter SKUFilter) *gorm.DB {
//	query := r.db.Model(&SKU{})
//
//	if filter.SPUID != 0 {
//		query = query.Where("spu_id = ?", filter.SPUID)
//	}
//	if filter.MinPrice > 0 {
//		query = query.Where("price >= ?", filter.MinPrice)
//	}
//	if filter.MaxPrice > 0 {
//		query = query.Where("price <= ?", filter.MaxPrice)
//	}
//	if filter.InStock {
//		query = query.Where("stock > 0")
//	}
//	if filter.IsActive != nil {
//		query = query.Where("is_active = ?", *filter.IsActive)
//	}
//	if filter.HasSpecs {
//		query = query.Joins("JOIN attribute_values ON skus.id = attribute_values.sku_id").
//			Group("skus.id")
//	}
//	if !filter.CreatedAfter.IsZero() {
//		query = query.Where("created_at >= ?", filter.CreatedAfter)
//	}
//
//	return query
//}
//
//// 保存规格属性（内部方法）
//func (r *SKURepo) saveSKUAttributes(tx *gorm.DB, skuID uint64, specs []AttributeValue) error {
//	if len(specs) == 0 {
//		return nil
//	}
//
//	cleanSpecs := make([]AttributeValue, len(specs))
//	for i := range specs {
//		cleanSpecs[i] = AttributeValue{
//			SkuID: skuID,
//			KeyID: specs[i].KeyID,
//			Value: specs[i].Value,
//		}
//	}
//
//	return tx.CreateInBatches(cleanSpecs, 500).Error
//}
//
//// 将属性名转换为(key_id, value)对 (内部方法)
//func getKeyValuePairs(keyIDs []uint64, filters map[string]string) [][]interface{} {
//	var pairs [][]interface{}
//	for _, value := range filters {
//		for _, keyID := range keyIDs {
//			// 实际业务中需要建立Name到KeyID的映射缓存优化
//			pairs = append(pairs, []interface{}{keyID, value})
//		}
//	}
//	return pairs
//}
//
//// UpdateSKUAttributes 全量更新SKU属性（原子操作）
//func (r *SKURepo) UpdateSKUAttributes(ctx context.Context, skuID uint64, newSpecs []AttributeValue) error {
//	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
//		// 1. 校验属性键有效性
//		keyIDs := make([]uint64, len(newSpecs))
//		for i, spec := range newSpecs {
//			keyIDs[i] = spec.KeyID
//		}
//
//		var validKeys int64
//		if err := tx.Model(&AttributeKey{}).
//			Where("key_id IN ?", keyIDs).
//			Count(&validKeys).Error; err != nil {
//			return fmt.Errorf("属性键校验失败: %w", err)
//		}
//		if validKeys != int64(len(keyIDs)) {
//			return fmt.Errorf("存在%d个无效属性键", len(keyIDs)-int(validKeys))
//		}
//
//		// 2. 删除旧属性（带行锁）
//		if err := tx.Set("gorm:query_option", "FOR UPDATE").
//			Where("sku_id = ?", skuID).
//			Delete(&AttributeValue{}).Error; err != nil {
//			return fmt.Errorf("删除旧属性失败: %w", err)
//		}
//
//		// 3. 插入新属性
//		for i := range newSpecs {
//			newSpecs[i].SkuID = skuID // 统一设置SKU ID
//		}
//
//		if err := tx.CreateInBatches(newSpecs, 500).Error; err != nil {
//			return fmt.Errorf("插入新属性失败: %w", err)
//		}
//
//		// 4. 更新SKU更新时间
//		return tx.Model(&SKU{}).
//			Where("id = ?", skuID).
//			Update("updated_at", gorm.Expr("NOW()")).Error
//	})
//}
//
//// PatchSKUAttributes 增量更新SKU属性（仅更新指定属性）
//func (r *SKURepo) PatchSKUAttributes(ctx context.Context, skuID uint64, updates map[string]string) error {
//	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
//		// 1. 转换属性名称为ID
//		nameToID, invalidNames, err := r.resolveAttributeNamesConcurrently(tx, updates)
//		if err != nil {
//			return err
//		}
//		if len(invalidNames) > 0 {
//			return fmt.Errorf("无效的属性名称: %v", strings.Join(invalidNames, ", "))
//		}
//
//		// 2. 构建ID到值的映射
//		idUpdates := make(map[uint64]string)
//		for name, value := range updates {
//			idUpdates[nameToID[name]] = value
//		}
//
//		// 3. 获取现有属性键
//		var existingKeys []uint64
//		if err := tx.Model(&AttributeValue{}).
//			Where("sku_id = ?", skuID).
//			Pluck("key_id", &existingKeys).Error; err != nil {
//			return err
//		}
//
//		// 4. 分离更新操作
//		var toCreate []AttributeValue
//		var toUpdate []struct {
//			KeyID uint64
//			Value string
//		}
//
//		for keyID, value := range idUpdates {
//			if slices.Contains(existingKeys, keyID) {
//				toUpdate = append(toUpdate, struct {
//					KeyID uint64
//					Value string
//				}{KeyID: keyID, Value: value})
//			} else {
//				toCreate = append(toCreate, AttributeValue{
//					SkuID: skuID,
//					KeyID: keyID,
//					Value: value,
//				})
//			}
//		}
//
//		// 5. 批量更新现有属性
//		if len(toUpdate) > 0 {
//			updateQuery := "UPDATE attribute_values SET value = CASE key_id "
//			var params []interface{}
//
//			for _, item := range toUpdate {
//				updateQuery += "WHEN ? THEN ? "
//				params = append(params, item.KeyID, item.Value)
//			}
//
//			updateQuery += "END WHERE sku_id = ?"
//			params = append(params, skuID)
//
//			if err := tx.Exec(updateQuery, params...).Error; err != nil {
//				return fmt.Errorf("属性更新失败: %w", err)
//			}
//		}
//
//		// 6. 批量创建新属性
//		if len(toCreate) > 0 {
//			if err := tx.CreateInBatches(toCreate, 500).Error; err != nil {
//				return fmt.Errorf("属性创建失败: %w", err)
//			}
//		}
//
//		return nil
//	})
//}
//
//// 使用goroutine并行处理,解析属性名称到ID的映射，返回无效名称列表
//func (r *SKURepo) resolveAttributeNamesConcurrently(tx *gorm.DB, updates map[string]string) (map[string]uint64, []string, error) {
//	// 收集所有属性名称
//	names := make([]string, 0, len(updates))
//	for name := range updates {
//		names = append(names, name)
//	}
//
//	type result struct {
//		name string
//		id   uint64
//		err  error
//	}
//
//	// 批量查询属性键
//	ch := make(chan result, len(names))
//	sem := make(chan struct{}, 10) // 控制并发数
//	wg := sync.WaitGroup{}
//	for _, name := range names {
//		wg.Add(1)
//		go func(n string) {
//			defer wg.Done()
//			sem <- struct{}{}
//			defer func() { <-sem }()
//
//			var key AttributeKey
//			err := tx.Where("name = ?", n).First(&key).Error
//			ch <- result{n, key.KeyID, err}
//		}(name)
//	}
//	go func() {
//		wg.Wait()
//		close(ch)
//	}()
//
//	// 构建名称到ID的映射
//	nameToID := make(map[string]uint64)
//	for range names {
//		res := <-ch
//		if res.err != nil {
//			continue // 或收集错误
//		}
//		nameToID[res.name] = res.id
//	}
//
//	// 检查无效名称
//	var invalidNames []string
//	for name := range updates {
//		if _, exists := nameToID[name]; !exists {
//			invalidNames = append(invalidNames, name)
//		}
//	}
//
//	return nameToID, invalidNames, nil
//}
