package repo

import (
	"context"
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Yzc216/gomall/app/product/biz/model"
	"github.com/allegro/bigcache"
	"github.com/redis/go-redis/v9"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"
	"math/rand"
	"time"
)

const (
	localTTL = 3 * time.Minute  // 本地缓存TTL
	redisTTL = 30 * time.Minute // Redis缓存TTL
)

type CachedSPUQuery interface {
	SPUQueryRepository

	//TODO
	RefreshCache(ctx context.Context, spuID uint64) error
	BatchRefreshCache(ctx context.Context, spuIDs []uint64) error
}

type CachedProductQuery struct {
	localCache   *bigcache.BigCache
	redisCache   *redis.Client
	SPUQuery     *SPUQuery
	SKUQuery     *SKUQuery
	singleFlight singleflight.Group

	prefix string
}

func NewCachedProductQuery(db *gorm.DB, cacheClient *redis.Client) *CachedProductQuery {
	cache, err := bigcache.NewBigCache(bigcache.Config{
		Shards:      1024,
		LifeWindow:  localTTL,
		CleanWindow: time.Minute,
	})
	if err != nil {
		return nil
	}
	return &CachedProductQuery{
		localCache:   cache,
		redisCache:   cacheClient,
		SPUQuery:     NewSPUQuery(db),
		singleFlight: singleflight.Group{},

		prefix: "shop",
	}
}

func (c *CachedProductQuery) GetByID(ctx context.Context, productId uint64) (product *model.SPU, err error) {
	cacheKey := fmt.Sprintf("%s_%s_%d", c.prefix, "product_by_id", productId)

	// 1. 查询本地缓存
	if cached, err := c.localCache.Get(cacheKey); err == nil {
		if len(cached) > 0 {
			var spu model.SPU
			if err = json.Unmarshal(cached, &spu); err == nil {
				return &spu, nil
			}
		}
	}

	// 2. 本地缓存未命中，尝试从 Redis 读取
	redisResult, err := c.redisCache.Get(ctx, cacheKey).Bytes()
	if err == nil && len(redisResult) > 0 {
		// 解码 Redis 数据
		var spu model.SPU
		if err := json.Unmarshal(redisResult, &spu); err == nil {
			// 异步回填本地缓存
			go func() {
				if data, err := json.Marshal(spu); err == nil {
					_ = c.localCache.Set(cacheKey, data)
				}
			}()
			return &spu, nil
		}
	}

	// 使用 singleflight 防止缓存击穿
	result, err, _ := c.singleFlight.Do(cacheKey, func() (interface{}, error) {
		// 3. 从数据库查询
		spu, err := c.SPUQuery.GetByID(ctx, productId)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				// 防止缓存穿透：缓存空值
				emptyData, _ := json.Marshal(model.SPU{})
				_ = c.localCache.Set(cacheKey, emptyData)
				c.redisCache.Set(ctx, cacheKey, emptyData, 5*time.Minute)
			}
			return nil, err
		}

		// 序列化数据
		data, err := json.Marshal(spu)
		if err != nil {
			return nil, err
		}

		// 异步更新缓存
		go func() {
			// 更新本地缓存
			_ = c.localCache.Set(cacheKey, data)

			// 更新 Redis 缓存，设置 TTL
			randomOffset := time.Duration(rand.Intn(5)) * time.Minute
			c.redisCache.Set(ctx, cacheKey, data, redisTTL+randomOffset)
		}()

		return spu, nil
	})

	if err != nil {
		return nil, err
	}

	return result.(*model.SPU), nil
}

func (c *CachedProductQuery) List(ctx context.Context, filter *SPUFilter, page *Pagination) ([]*model.SPU, int64, error) {
	// 生成唯一缓存键（示例序列化方式，可根据实际调整）
	filterHash := fmt.Sprintf("%v", struct {
		KeyWord string
		ShopID  uint64
		Brand   string
		Status  int8
	}{filter.Keyword, filter.ShopID, filter.Brand, filter.Status})

	cacheKey := fmt.Sprintf("%s_list_%s_p%d_s%d",
		c.prefix,
		md5.Sum([]byte(filterHash)),
		page.Page,
		page.PageSize,
	)

	// 尝试从本地缓存获取
	if cached, err := c.localCache.Get(cacheKey); err == nil {
		var cachedResult struct {
			List  []*model.SPU
			Total int64
		}
		if err := json.Unmarshal(cached, &cachedResult); err == nil {
			return cachedResult.List, cachedResult.Total, nil
		}
	}

	// 尝试从Redis获取
	redisResult, err := c.redisCache.Get(ctx, cacheKey).Bytes()
	if err == nil {
		var cachedResult struct {
			List  []*model.SPU
			Total int64
		}
		if err := json.Unmarshal(redisResult, &cachedResult); err == nil {
			// 异步回填本地缓存
			go func() {
				if data, err := json.Marshal(cachedResult); err == nil {
					_ = c.localCache.Set(cacheKey, data)
				}
			}()
			return cachedResult.List, cachedResult.Total, nil
		}
	}

	// 使用singleflight合并请求
	result, err, _ := c.singleFlight.Do(cacheKey, func() (interface{}, error) {
		list, total, err := c.SPUQuery.List(ctx, filter, page)
		if err != nil {
			return nil, err
		}

		// 构造缓存结果
		cacheData := struct {
			List  []*model.SPU
			Total int64
		}{List: list, Total: total}

		// 异步更新缓存
		go func() {
			if data, err := json.Marshal(cacheData); err == nil {
				// 本地缓存设置短时间（5分钟+随机）
				_ = c.localCache.Set(cacheKey, data)

				// Redis设置基础30分钟 + 随机防雪崩
				baseTTL := 30 * time.Minute
				randomOffset := time.Duration(rand.Intn(300)) * time.Second // 0-5分钟随机
				c.redisCache.Set(ctx, cacheKey, data, baseTTL+randomOffset)
			}
		}()

		return &cacheData, nil
	})

	if err != nil {
		return nil, 0, err
	}

	cachedResult := result.(*struct {
		List  []*model.SPU
		Total int64
	})
	return cachedResult.List, cachedResult.Total, nil
}
