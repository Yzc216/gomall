package repo

import (
	"context"
	"errors"
	"fmt"
	"github.com/Yzc216/gomall/app/product/biz/model"
	"github.com/Yzc216/gomall/app/product/biz/types"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

var (
	_ SPURepository = (*SPURepo)(nil)
)

type SPURepository interface {
	// 基础单条操作
	Create(ctx context.Context, spu *model.SPU) error
	Update(ctx context.Context, spu *model.SPU) error
	Delete(ctx context.Context, id uint64) error
	GetByID(ctx context.Context, id uint64) (*model.SPU, error)
	GetByName(ctx context.Context, name string) (*model.SPU, error)
	ExistsByID(ctx context.Context, id uint64) (bool, error)
	ExistByTitle(ctx context.Context, title string) (bool, error)

	// 复杂查询
	List(ctx context.Context, filter SPUFilter, page Pagination) ([]*model.SPU, int64, error)
	FullTextSearch(ctx context.Context, keyword string, page Pagination) ([]*model.SPU, int64, error)
	//GetInventoryAlert(ctx context.Context,spuID uint64) (int, error) // 库存预警数量

	// 关联数据操作
	GetSKUs(ctx context.Context, spuID uint64) ([]*model.SKU, error)
	GetSKUCount(ctx context.Context, spuID uint64) (int64, error)
	AddSKUs(ctx context.Context, spuID uint64, skuIDs []uint64) error
	RemoveSKUs(ctx context.Context, spuID uint64, skuIDs []uint64) error

	GetCategories(ctx context.Context, spuID uint64) ([]*model.Category, error)
	AddCategory(ctx context.Context, spuID, categoryID uint64) error
	RemoveCategories(ctx context.Context, spuID uint64, categoryIDs []uint64) error
	ReplaceCategories(ctx context.Context, spuID uint64, categoryIDs []uint64) error

	// 多媒体管理
	UpdateMedia(ctx context.Context, spuID uint64, media Media) error

	// 状态管理
	UpdateStatus(ctx context.Context, spuID uint64, status int8) error

	// 缓存管理
	//RefreshCache(ctx context.Context,spuID uint64) error
	//BatchRefreshCache(ctx context.Context,spuIDs []uint64) error

	// 定时任务
	//AutoOfflineExpiredSPUs(ctx context.Context,) error // 自动下架过期商品
	//UpdateSalesCount(ctx context.Context, spuID uint64) error

	//批量操作
	BatchGetByIDs(ctx context.Context, ids []uint64) ([]*model.SPU, error)
	//BatchCreateSPUs(ctx context.Context, spus []*model.SPU) error
	//BatchUpdateStatus(ctx context.Context, spuIDs []uint64, status int8) error
}

type SPUFilter struct {
	ShopID      []uint64
	BrandID     []uint64
	CategoryID  []uint64
	Status      []int8
	MinPrice    float64 // 基于SKU最低价
	MaxPrice    float64
	CreateStart time.Time
	CreateEnd   time.Time
	Keyword     string
	OrderBy     string // 排序字段示例："sales_desc"/"price_asc"
}

type Media struct {
	MainImages []string
	Video      string
}
type SPURepo struct {
	ctx context.Context
	db  *gorm.DB
}

func NewSPURepo(ctx context.Context, db *gorm.DB) *SPURepo {
	return &SPURepo{
		ctx: ctx,
		db:  db,
	}
}

// 基础单条操作
func (r *SPURepo) Create(ctx context.Context, spu *model.SPU) error {
	exist, err := r.ExistByTitle(ctx, spu.Title)
	if err != nil {
		return fmt.Errorf("check title existence failed: %w", err)
	}
	if exist {
		return types.ErrSPUTitleExists
	}

	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(spu).Error; err != nil {
			if errors.Is(err, gorm.ErrDuplicatedKey) {
				return types.ErrSPUTitleExists
			}
			return fmt.Errorf("create SPU failed: %w", err)
		}

		// 处理分类关联
		if len(spu.Categories) > 0 {
			if err := tx.Model(spu).Association("Categories").Append(spu.Categories); err != nil {
				return fmt.Errorf("add categories failed: %w", err)
			}
		}
		return nil
	})
}

func (r *SPURepo) Update(ctx context.Context, spu *model.SPU) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 获取当前数据并加锁
		var current model.SPU
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
			First(&current, spu.ID).Error; err != nil {
			return types.ErrSPUNotFound
		}

		// 校验标题唯一性
		if current.Title != spu.Title {
			exist, err := r.ExistByTitle(ctx, spu.Title)
			if err != nil || exist {
				return types.ErrSPUTitleExists
			}
		}

		return tx.Select("*").Updates(spu).Error
	})
}

func (r *SPURepo) Delete(ctx context.Context, id uint64) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 检查关联SKU
		var count int64
		if err := tx.Model(&model.SKU{}).Where("spu_id = ?", id).Count(&count).Error; err != nil {
			return err
		}
		if count > 0 {
			return types.ErrHasAssociatedSKUs
		}

		// 删除分类关联
		if err := tx.Model(&model.SPU{ID: id}).Association("Categories").Clear(); err != nil {
			return err
		}

		// 删除主记录
		return tx.Delete(&model.SPU{}, id).Error
	})
}

func (r *SPURepo) GetByID(ctx context.Context, id uint64) (*model.SPU, error) {
	var spu model.SPU
	err := r.db.WithContext(ctx).
		Preload("SKUs").
		Preload("Categories").
		First(&spu, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, types.ErrSPUNotFound
	}
	return &spu, err
}

func (r *SPURepo) GetByName(ctx context.Context, name string) (*model.SPU, error) {
	var spu model.SPU
	err := r.db.WithContext(ctx).
		Where("title = ?", name).
		First(&spu).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, types.ErrSPUNotFound
	}
	return &spu, err
}

func (r *SPURepo) ExistsByID(ctx context.Context, id uint64) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).
		Model(&model.SPU{}).
		Where("id = ?", id).
		Count(&count).Error
	return count > 0, err
}

func (r *SPURepo) ExistByTitle(ctx context.Context, title string) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).
		Model(&model.SPU{}).
		Where("title = ?", title).
		Count(&count).Error
	return count > 0, err
}

// 复杂查询
// --------------------------------------------------
func (r *SPURepo) List(ctx context.Context, filter SPUFilter, page Pagination) ([]*model.SPU, int64, error) {
	query := r.buildListQuery(filter)

	var total int64
	if err := query.Model(&model.SPU{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	var spus []*model.SPU
	err := query.Scopes(r.paginate(page)).
		Preload("SKUs", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, spu_id, price, stock").Where("is_active = ?", true)
		}).
		Find(&spus).Error

	return spus, total, err
}

func (r *SPURepo) FullTextSearch(ctx context.Context, keyword string, page Pagination) ([]*model.SPU, int64, error) {
	query := r.db.WithContext(ctx).
		Model(&model.SPU{}).
		Where("MATCH(title, description) AGAINST(? IN BOOLEAN MODE)", keyword+"*")

	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	var spus []*model.SPU
	err := query.Scopes(r.paginate(page)).
		Find(&spus).Error

	return spus, total, err
}

func (r *SPURepo) BatchGetByIDs(ctx context.Context, ids []uint64) ([]*model.SPU, error) {
	var spus []*model.SPU
	err := r.db.WithContext(ctx).
		Where("id IN ?", ids).
		Find(&spus).Error
	return spus, err
}

// 关联数据操作
// --------------------------------------------------
func (r *SPURepo) GetSKUs(ctx context.Context, spuID uint64) ([]*model.SKU, error) {
	var skus []*model.SKU
	err := r.db.WithContext(ctx).
		Where("spu_id = ?", spuID).
		Find(&skus).Error
	return skus, err
}

func (r *SPURepo) GetSKUCount(ctx context.Context, spuID uint64) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).
		Model(&model.SKU{}).
		Where("spu_id = ?", spuID).
		Count(&count).Error
	return count, err
}

func (r *SPURepo) AddSKUs(ctx context.Context, spuID uint64, skuIDs []uint64) error {
	// 验证SKU存在性
	var count int64
	if err := r.db.WithContext(ctx).
		Model(&model.SKU{}).
		Where("id IN ?", skuIDs).
		Count(&count).Error; err != nil {
		return err
	}
	if int(count) != len(skuIDs) {
		return errors.New("invalid SKU IDs")
	}

	return r.db.WithContext(ctx).Model(&model.SPU{ID: spuID}).
		Association("SKUs").
		Append(generateSKUReferences(skuIDs))
}

func (r *SPURepo) RemoveSKUs(ctx context.Context, spuID uint64, skuIDs []uint64) error {
	return r.db.WithContext(ctx).Model(&model.SPU{ID: spuID}).
		Association("SKUs").
		Delete(generateSKUReferences(skuIDs))
}

// 分类关联操作
// --------------------------------------------------
func (r *SPURepo) GetCategories(ctx context.Context, spuID uint64) ([]*model.Category, error) {
	var categories []*model.Category
	err := r.db.WithContext(ctx).Model(&model.SPU{ID: spuID}).
		Association("Categories").
		Find(&categories)
	return categories, err
}

func (r *SPURepo) AddCategory(ctx context.Context, spuID, categoryID uint64) error {
	return r.db.WithContext(ctx).Model(&model.SPU{ID: spuID}).
		Association("Categories").
		Append(&model.Category{ID: categoryID})
}

func (r *SPURepo) RemoveCategories(ctx context.Context, spuID uint64, categoryIDs []uint64) error {
	refs := make([]*model.Category, len(categoryIDs))
	for i, id := range categoryIDs {
		refs[i] = &model.Category{ID: id}
	}
	return r.db.WithContext(ctx).Model(&model.SPU{ID: spuID}).
		Association("Categories").
		Delete(refs)
}

func (r *SPURepo) ReplaceCategories(ctx context.Context, spuID uint64, categoryIDs []uint64) error {
	categories := make([]*model.Category, len(categoryIDs))
	for i, id := range categoryIDs {
		categories[i] = &model.Category{ID: id}
	}
	return r.db.WithContext(ctx).Model(&model.SPU{ID: spuID}).
		Association("Categories").
		Replace(categories)
}

// 多媒体管理
// --------------------------------------------------
func (r *SPURepo) UpdateMedia(ctx context.Context, spuID uint64, media Media) error {
	return r.db.WithContext(ctx).Model(&model.SPU{}).
		Where("id = ?", spuID).
		Updates(map[string]interface{}{
			"main_images": media.MainImages,
			"video":       media.Video,
		}).Error
}

// 状态管理
// --------------------------------------------------
func (r *SPURepo) UpdateStatus(ctx context.Context, spuID uint64, status int8) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
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

// 私有辅助方法
// --------------------------------------------------
func (r *SPURepo) buildListQuery(filter SPUFilter) *gorm.DB {
	query := r.db.Model(&model.SPU{})

	if len(filter.ShopID) > 0 {
		query = query.Where("shop_id IN ?", filter.ShopID)
	}
	if len(filter.BrandID) > 0 {
		query = query.Where("brand_id IN ?", filter.BrandID)
	}
	if len(filter.Status) > 0 {
		query = query.Where("status IN ?", filter.Status)
	}
	if !filter.CreateStart.IsZero() {
		query = query.Where("created_at >= ?", filter.CreateStart)
	}
	if !filter.CreateEnd.IsZero() {
		query = query.Where("created_at <= ?", filter.CreateEnd)
	}
	if filter.Keyword != "" {
		query = query.Where("title LIKE ?", "%"+filter.Keyword+"%")
	}
	if filter.MinPrice > 0 || filter.MaxPrice > 0 {
		query = query.Joins(
			"INNER JOIN (SELECT spu_id, MIN(price) AS min_price FROM skus GROUP BY spu_id) s ON spus.id = s.spu_id").
			Where("s.min_price BETWEEN ? AND ?", filter.MinPrice, filter.MaxPrice)
	}
	if len(filter.CategoryID) > 0 {
		query = query.Joins(
			"JOIN spu_categories ON spus.id = spu_categories.spu_id").
			Where("spu_categories.category_id IN ?", filter.CategoryID)
	}

	// 排序处理
	switch filter.OrderBy {
	case "sales_desc":
		query = query.Order("sales DESC")
	case "price_asc":
		query = query.Order("s.min_price ASC")
	default:
		query = query.Order("created_at DESC")
	}

	return query
}

func (r *SPURepo) paginate(page Pagination) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset := (page.Page - 1) * page.PageSize
		return db.Offset(offset).Limit(page.PageSize)
	}
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
