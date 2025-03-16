package repo

import (
	"context"
	"github.com/Yzc216/gomall/app/product/biz/model"
	"gorm.io/gorm"
	"strings"
)

type SKUQuery struct {
	db *gorm.DB
}

func NewSKUQuery(db *gorm.DB) *SKUQuery {
	return &SKUQuery{db: db}
}

func (r *SKUQuery) GetSKUs(ctx context.Context, spuID uint64) ([]*model.SKU, error) {
	var skus []*model.SKU
	err := r.db.WithContext(ctx).
		Where("spu_id = ?", spuID).
		Find(&skus).Error
	return skus, err
}

func (r *SKUQuery) GetSKUCount(ctx context.Context, spuID uint64) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).
		Model(&model.SKU{}).
		Where("spu_id = ?", spuID).
		Count(&count).Error
	return count, err
}

func (q *SKUQuery) buildListQuery(ctx context.Context, filter *SPUFilter) *gorm.DB {
	query := q.db.WithContext(ctx).Model(&model.SPU{})
	builders := []QueryBuilder{
		q.buildShopQuery(filter.ShopID),
		q.buildBrandQuery(filter.Brand),
		q.buildStatusQuery(filter.Status),
		q.buildKeywordQuery(filter.Keyword),
		q.buildOrderByQuery(filter.OrderBy),
		//q.buildPriceRangeQuery(filter.MinPrice, filter.MaxPrice),
	}
	for _, builder := range builders {
		query = builder(query)
	}
	return query
}

func (q *SKUQuery) buildKeywordQuery(keyword string) QueryBuilder {
	return func(db *gorm.DB) *gorm.DB {
		// 关键词搜索（标题+副标题）
		if keyword = strings.TrimSpace(keyword); keyword == "" {
			return db
		}
		return db.Where("title LIKE ? OR sub_title LIKE ?",
			"%"+keyword+"%", "%"+keyword+"%")
	}
}

func (q *SKUQuery) buildShopQuery(shopID uint64) QueryBuilder {
	return func(db *gorm.DB) *gorm.DB {
		if shopID <= 0 {
			return db
		}
		return db.Where("shop_id = ?", shopID)
	}
}

func (q *SKUQuery) buildBrandQuery(brand string) QueryBuilder {
	return func(db *gorm.DB) *gorm.DB {
		if brand == "" {
			return db
		}
		return db.Where("brand = ?", brand)
	}
}

func (q *SKUQuery) buildStatusQuery(status int8) QueryBuilder {
	return func(db *gorm.DB) *gorm.DB {
		if status < 0 {
			return db
		}
		return db.Where("status = ?", status)
	}
}

func (q *SKUQuery) buildOrderByQuery(orderBy string) QueryBuilder {
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

func (q *SKUQuery) buildPriceRangeQuery(min, max float64) QueryBuilder {
	return func(db *gorm.DB) *gorm.DB {
		// 参数校验
		if min < 0 || max < 0 {
			return db
		}
		if min > 0 && max > 0 && min > max {
			return db
		}
		return db
	}
}
