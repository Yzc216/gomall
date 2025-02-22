package repo

import (
	"context"
	"errors"
	"fmt"
	"github.com/Yzc216/gomall/app/product/biz/model"
	"github.com/Yzc216/gomall/app/product/biz/types"
	"gorm.io/gorm"
	"time"
)

var (
	_ BrandRepository = (*BrandRepo)(nil)
)

type BrandRepository interface {
	// 基础CRUD
	Create(ctx context.Context, brand *model.Brand) error
	Update(ctx context.Context, brand *model.Brand) error
	Delete(ctx context.Context, id uint64) error
	GetByID(ctx context.Context, id uint64) (*model.Brand, error)
	GetByName(ctx context.Context, name string) (*model.Brand, error)
	ExistsByID(ctx context.Context, id uint64) (bool, error)
	ExistsByName(ctx context.Context, name string) (bool, error)

	// 分页与查询
	List(ctx context.Context, filter BrandFilter, page Pagination) ([]*model.Brand, int64, error)

	// 关联数据操作
	ListSPUs(ctx context.Context, brandID uint64, includeInactive bool) ([]*model.SPU, error)
	GetSPUCount(ctx context.Context, brandID uint64) (int64, error)
	AddSPUs(ctx context.Context, brandID uint64, spuIDs []uint64) error
	RemoveSPUs(ctx context.Context, brandID uint64, spuIDs []uint64) error

	// 批量操作
	BatchCreate(ctx context.Context, brands []*model.Brand) ([]uint64, error)
	BatchUpdateLogo(ctx context.Context, ids []uint64, logo []byte) error
	BatchToggleOfficialStatus(ctx context.Context, ids []uint64, isOfficial bool) error
}

// 查询过滤条件
type BrandFilter struct {
	Name         string
	IsOfficial   *bool // 使用指针允许三态查询
	MinSPUCount  int
	CreatedAfter time.Time
}

type BrandRepo struct {
	db *gorm.DB
}

func NewBrandRepo(db *gorm.DB) *BrandRepo {
	return &BrandRepo{
		db: db,
	}
}

// Create 创建品牌（带唯一性校验）
func (r *BrandRepo) Create(ctx context.Context, brand *model.Brand) error {
	exist, err := r.ExistsByName(ctx, brand.Name)
	if err != nil {
		return fmt.Errorf("check brand existence failed: %w", err)
	}
	if exist {
		return types.ErrBrandExists
	}

	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(brand).Error; err != nil {
			if errors.Is(err, gorm.ErrDuplicatedKey) {
				return types.ErrBrandExists
			}
			return fmt.Errorf("create brand failed: %w", err)
		}
		return nil
	})
}

// Update 更新品牌信息
func (r *BrandRepo) Update(ctx context.Context, brand *model.Brand) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		// 获取当前数据
		var current model.Brand
		if err := tx.First(&current, brand.ID).Error; err != nil {
			return types.ErrBrandNotFound
		}

		// 名称修改时需要校验唯一性
		if current.Name != brand.Name {
			exist, err := r.ExistsByName(ctx, brand.Name)
			if err != nil || exist {
				return types.ErrBrandExists
			}
		}

		// 执行更新
		if err := tx.Select("*").Updates(brand).Error; err != nil {
			return fmt.Errorf("update brand failed: %w", err)
		}
		return nil
	})
}

// Delete 删除品牌（带关联检查）
func (r *BrandRepo) Delete(ctx context.Context, id uint64) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		// 检查存在性
		if exist, _ := r.ExistsByID(ctx, id); !exist {
			return types.ErrBrandNotFound
		}

		// 检查关联SPU
		var count int64
		if err := tx.Model(&model.SPU{}).Where("brand_id = ?", id).Count(&count).Error; err != nil {
			return err
		}
		if count > 0 {
			return types.ErrHasAssociatedSPUs
		}

		// 执行删除
		return tx.Delete(&model.Brand{}, id).Error
	})
}

// GetByID 获取品牌详情
func (r *BrandRepo) GetByID(ctx context.Context, id uint64) (*model.Brand, error) {
	var brand model.Brand
	err := r.db.Preload("SPUs").First(&brand, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, types.ErrBrandNotFound
	}
	return &brand, err
}

func (r *BrandRepo) GetByName(ctx context.Context, name string) (*model.Brand, error) {
	var brand model.Brand
	err := r.db.Preload("SPUs").Where("name = ?", name).First(&brand).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, types.ErrBrandNotFound
	}
	return &brand, err
}

// ExistsByName 检查品牌名称是否存在
func (r *BrandRepo) ExistsByName(ctx context.Context, name string) (bool, error) {
	var count int64
	err := r.db.Model(&model.Brand{}).
		Where("name = ?", name).
		Count(&count).Error
	return count > 0, err
}

// ExistsByID 检查品牌ID是否存在
func (r *BrandRepo) ExistsByID(ctx context.Context, id uint64) (bool, error) {
	var count int64
	err := r.db.Model(&model.Brand{}).
		Where("id = ?", id).
		Count(&count).Error
	return count > 0, err
}

// List 分页查询品牌列表
func (r *BrandRepo) List(ctx context.Context, filter BrandFilter, page Pagination) ([]*model.Brand, int64, error) {
	query := r.buildQuery(ctx, filter)

	var total int64
	if err := query.Model(&model.Brand{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	var brands []*model.Brand
	err := query.Scopes(Paginate(page)).
		Order("name ASC").
		Find(&brands).Error

	return brands, total, err
}

// ListOfficial 查询官方/非官方品牌
func (r *BrandRepo) ListOfficial(ctx context.Context, onlyOfficial bool, page Pagination) ([]*model.Brand, int64, error) {
	var brands []*model.Brand
	var total int64

	query := r.db.Model(&model.Brand{}).Where("is_official = ?", onlyOfficial)
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := query.Scopes(Paginate(page)).
		Order("name ASC").
		Find(&brands).Error

	return brands, total, err
}

// GetSPUCount 获取品牌关联商品数量
func (r *BrandRepo) GetSPUCount(ctx context.Context, brandID uint64) (int64, error) {
	var count int64
	err := r.db.Model(&model.SPU{}).
		Where("brand_id = ?", brandID).
		Count(&count).Error
	return count, err
}

// ListSPUs 获取品牌商品列表
func (r *BrandRepo) ListSPUs(ctx context.Context, brandID uint64, includeInactive bool) ([]*model.SPU, error) {
	query := r.db.WithContext(ctx).Model(&model.SPU{}).Where("brand_id = ?", brandID)
	if !includeInactive {
		query = query.Where("status = ?", 1)
	}

	var spus []*model.SPU
	err := query.Find(&spus).Error
	return spus, err
}

// AddSPUs 关联商品到品牌
func (r *BrandRepo) AddSPUs(ctx context.Context, brandID uint64, spuIDs []uint64) error {
	if len(spuIDs) == 0 {
		return nil
	}

	// 检查SPU存在性
	var existing int64
	if err := r.db.Model(&model.SPU{}).
		Where("id IN ?", spuIDs).
		Count(&existing).Error; err != nil {
		return err
	}
	if int(existing) != len(spuIDs) {
		return errors.New("some SPUs not exist")
	}

	// 批量更新
	return r.db.Transaction(func(tx *gorm.DB) error {
		return tx.Model(&model.SPU{}).
			Where("id IN ?", spuIDs).
			Update("brand_id", brandID).Error
	})
}

// RemoveSPUs 解除商品关联
func (r *BrandRepo) RemoveSPUs(ctx context.Context, brandID uint64, spuIDs []uint64) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		return tx.Model(&model.SPU{}).
			Where("brand_id = ? AND id IN ?", brandID, spuIDs).
			Update("brand_id", 0).Error
	})
}

// BatchCreate 批量创建品牌
func (r *BrandRepo) BatchCreate(ctx context.Context, brands []*model.Brand) ([]uint64, error) {
	ids := make([]uint64, 0, len(brands))
	err := r.db.Transaction(func(tx *gorm.DB) error {
		// 分批处理防止SQL过长
		batchSize := 500
		for i := 0; i < len(brands); i += batchSize {
			end := i + batchSize
			if end > len(brands) {
				end = len(brands)
			}
			if err := tx.Create(brands[i:end]).Error; err != nil {
				return err
			}
			// 收集生成的ID
			for _, b := range brands[i:end] {
				ids = append(ids, b.ID)
			}
		}
		return nil
	})
	return ids, err
}

// BatchUpdateLogo 批量更新Logo
func (r *BrandRepo) BatchUpdateLogo(ctx context.Context, ids []uint64, logo []byte) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		return tx.Model(&model.Brand{}).
			Where("id IN ?", ids).
			Update("logo", logo).Error
	})
}

// BatchToggleOfficialStatus 批量切换官方状态
func (r *BrandRepo) BatchToggleOfficialStatus(ctx context.Context, ids []uint64, isOfficial bool) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		return tx.Model(&model.Brand{}).
			Where("id IN ?", ids).
			Update("is_official", isOfficial).Error
	})
}

// 私有方法
func (r *BrandRepo) buildQuery(ctx context.Context, filter BrandFilter) *gorm.DB {
	query := r.db.Model(&model.Brand{})

	if filter.Name != "" {
		query = query.Where("name LIKE ?", filter.Name+"%")
	}

	if filter.IsOfficial != nil {
		query = query.Where("is_official = ?", *filter.IsOfficial)
	}

	if filter.MinSPUCount > 0 {
		query = query.Joins(
			"LEFT JOIN spus ON spus.brand_id = brands.id").
			Group("brands.id").
			Having("COUNT(spus.id) >= ?", filter.MinSPUCount)
	}

	if !filter.CreatedAfter.IsZero() {
		query = query.Where("created_at >= ?", filter.CreatedAfter)
	}

	return query
}
