package repo

import (
	"context"
	"errors"
	"github.com/Yzc216/gomall/app/product/biz/model"
	"github.com/Yzc216/gomall/app/product/biz/types"
	"gorm.io/gorm"
	"strings"
)

var (
// _ SPUQueryRepository = (*SPUQuery)(nil)
)

type SPUQueryRepository interface {
	GetByID(ctx context.Context, id uint64) (*model.SPU, error)
	GetByName(ctx context.Context, name string) (*model.SPU, error)
	ExistsByID(ctx context.Context, id uint64) (bool, error)
	ExistByTitle(ctx context.Context, title string) (bool, error)

	// 复杂查询
	List(ctx context.Context, filter *SPUFilter, page *Pagination) ([]*model.SPU, int64, error)
	BatchGetByIDs(ctx context.Context, ids []uint64) (map[uint64]*model.SPU, error)
}

type SPUQuery struct {
	db *gorm.DB
}

func NewSPUQuery(db *gorm.DB) *SPUQuery {
	return &SPUQuery{db: db}
}

func (q *SPUQuery) GetByID(ctx context.Context, id uint64) (*model.SPU, error) {
	var spu model.SPU
	err := q.db.WithContext(ctx).
		Preload("SKUs").
		Preload("Categories").
		First(&spu, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, types.ErrSPUNotFound
	}
	return &spu, err
}

func (q *SPUQuery) GetByName(ctx context.Context, name string) (*model.SPU, error) {
	var spu model.SPU
	err := q.db.WithContext(ctx).
		Where("title = ?", name).
		Preload("SKUs").
		Preload("Categories").
		First(&spu).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, types.ErrSPUNotFound
	}
	return &spu, err
}

func (q *SPUQuery) ExistsByID(ctx context.Context, id uint64) (bool, error) {
	var count int64
	err := q.db.WithContext(ctx).
		Model(&model.SPU{}).
		Where("id = ?", id).
		Count(&count).Error
	return count > 0, err
}

func (q *SPUQuery) ExistByTitle(ctx context.Context, title string) (bool, error) {
	var count int64
	err := q.db.WithContext(ctx).
		Model(&model.SPU{}).
		Where("title = ?", title).
		Count(&count).Error
	return count > 0, err
}

// 复杂查询
func (q *SPUQuery) List(ctx context.Context, filter *SPUFilter, page *Pagination) ([]*model.SPU, int64, error) {
	query := q.buildListQuery(ctx, filter)

	var total int64
	if err := query.Model(&model.SPU{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if page.Page != 0 && page.PageSize != 0 {
		query = query.Offset(page.Offset()).Limit(page.PageSize)
	}

	var spus []*model.SPU
	err := query.Preload("SKUs", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, spu_id, price, stock").Where("is_active = ?", true)
	}).
		Find(&spus).Error

	return spus, total, err
}

func (q *SPUQuery) BatchGetByIDs(ctx context.Context, ids []uint64) (map[uint64]*model.SPU, error) {
	var SPUs []*model.SPU
	if err := q.db.WithContext(ctx).
		Where("id IN ?", ids).
		Preload("SKUs").
		Preload("Categories").
		Find(&SPUs).Error; err != nil {
		return nil, err
	}

	// 转换为 map 结构
	spuMap := make(map[uint64]*model.SPU, len(SPUs))
	for _, spu := range SPUs {
		spuMap[spu.ID] = spu
	}
	return spuMap, nil
}

// 辅助函数
func (q *SPUQuery) buildListQuery(ctx context.Context, filter *SPUFilter) *gorm.DB {
	query := q.db.WithContext(ctx).Model(&model.SPU{})
	builders := []QueryBuilder{
		//q.buildShopQuery(filter.ShopID),
		//q.buildBrandQuery(filter.Brand),
		//q.buildStatusQuery(filter.Status),
		q.buildKeywordQuery(filter.Keyword),
		//q.buildOrderByQuery(filter.OrderBy),
		//q.buildPriceRangeQuery(filter.MinPrice, filter.MaxPrice),
	}
	for _, builder := range builders {
		query = builder(query)
	}
	//if !filter.CreateStart.IsZero() {
	//	query = query.Where("created_at >= ?", filter.CreateStart)
	//}
	//if !filter.CreateEnd.IsZero() {
	//	query = query.Where("created_at <= ?", filter.CreateEnd)
	//}
	return query
}

func (q *SPUQuery) buildKeywordQuery(keyword string) QueryBuilder {
	return func(db *gorm.DB) *gorm.DB {
		// 关键词搜索（标题+副标题）
		if keyword = strings.TrimSpace(keyword); keyword == "" {
			return db
		}
		return db.Where("title LIKE ? OR sub_title LIKE ?",
			"%"+keyword+"%", "%"+keyword+"%")
	}
}

func (q *SPUQuery) buildShopQuery(shopID uint64) QueryBuilder {
	return func(db *gorm.DB) *gorm.DB {
		if shopID <= 0 {
			return db
		}
		return db.Where("shop_id = ?", shopID)
	}
}

func (q *SPUQuery) buildBrandQuery(brand string) QueryBuilder {
	return func(db *gorm.DB) *gorm.DB {
		if brand == "" {
			return db
		}
		return db.Where("brand = ?", brand)
	}
}

func (q *SPUQuery) buildStatusQuery(status int8) QueryBuilder {
	return func(db *gorm.DB) *gorm.DB {
		if status <= 0 {
			return db
		}
		return db.Where("status = ?", status)
	}
}

func (q *SPUQuery) buildOrderByQuery(orderBy string) QueryBuilder {
	return func(db *gorm.DB) *gorm.DB {
		if orderBy == "" {
			return db
		}
		switch orderBy {
		case "sales_desc":
			return db.Order("sales_count DESC")
		case "price_asc":
			return db.Order("min_price ASC")
		default:
			return db.Order("created_at DESC")
		}
	}
}

func (q *SPUQuery) buildPriceRangeQuery(min, max float64) QueryBuilder {
	return func(db *gorm.DB) *gorm.DB {
		// 参数校验
		if min <= 0 || max <= 0 {
			return db
		}
		if min > 0 && max > 0 && min > max {
			return db
		}
		return db
	}
}

func (q *SPUQuery) FullTextSearch(ctx context.Context, keyword string, page Pagination) ([]*model.SPU, int64, error) {
	query := q.db.WithContext(ctx).
		Model(&model.SPU{}).
		Where("MATCH(title, description) AGAINST(? IN BOOLEAN MODE)", keyword+"*")

	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	var spus []*model.SPU
	err := query.Scopes(Paginate(page)).
		Find(&spus).Error

	return spus, total, err
}

func (q *SPUQuery) GetCategories(ctx context.Context, spuID uint64) ([]*model.Category, error) {
	var categories []*model.Category
	err := q.db.WithContext(ctx).Model(&model.SPU{ID: spuID}).
		Association("Categories").
		Find(&categories)
	return categories, err
}
