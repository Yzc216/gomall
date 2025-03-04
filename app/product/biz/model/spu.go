package model

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"time"
)

type SPU struct {
	ID          uint64         `gorm:"primaryKey;comment:SPU ID"`
	Title       string         `gorm:"type:varchar(255);not null;comment:商品标题"`
	SubTitle    string         `gorm:"type:varchar(255);not null;comment:副标题"`
	ShopID      uint64         `gorm:"index:idx_shop;not null;comment:店铺ID"`
	Brand       string         `gorm:"index;not null;comment:品牌ID"`
	MainImages  []string       `gorm:"type:json;serializer:json;comment:主图URL列表"`
	Video       string         `gorm:"type:varchar(500);comment:商品视频URL"`
	Description string         `gorm:"type:text;comment:商品详情"`
	Status      int8           `gorm:"type:tinyint;default:0;index:idx_status;comment:状态（0-下架 1-上架 2-待审核）"`
	CreatedAt   time.Time      `gorm:"comment:创建时间"`
	UpdatedAt   time.Time      `gorm:"comment:更新时间"`
	DeletedAt   gorm.DeletedAt `gorm:"index;comment:软删除时间"`

	SKUs       []SKU      `gorm:"foreignKey:SpuID;references:ID;comment:关联SKU"`
	Categories []Category `gorm:"many2many:spu_categories;joinForeignKey:spu_id;joinReferences:category_id;comment:关联分类"`
}

func (SPU) TableName() string {
	return "spu"
}

type CachedProductQuery struct {
	productQuery *SPUQuery
	cacheClient  *redis.Client
	prefix       string
}

func NewCachedProductQuery(db *gorm.DB, cacheClient *redis.Client) *CachedProductQuery {
	return &CachedProductQuery{
		productQuery: NewSPUQuery(db),
		cacheClient:  cacheClient,
		prefix:       "shop",
	}
}

func (c CachedProductQuery) GetByID(ctx context.Context, productId uint64) (product *SPU, err error) {
	cacheKey := fmt.Sprintf("%s_%s_%d", c.prefix, "product_by_id", productId)
	cachedResult := c.cacheClient.Get(ctx, cacheKey)

	err = func() error {
		err1 := cachedResult.Err()
		if err1 != nil {
			return err1
		}
		cachedResultByte, err2 := cachedResult.Bytes()
		if err2 != nil {
			return err2
		}
		err3 := json.Unmarshal(cachedResultByte, &product)
		if err3 != nil {
			return err3
		}
		return nil
	}()
	if err != nil {
		product, err = c.productQuery.GetByID(ctx, productId)
		if err != nil {
			return nil, err
		}
		encoded, err := json.Marshal(product)
		if err != nil {
			return product, nil
		}
		_ = c.cacheClient.Set(ctx, cacheKey, encoded, time.Hour)
	}
	return
}
