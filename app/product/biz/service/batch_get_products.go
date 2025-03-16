package service

import (
	"context"
	"github.com/Yzc216/gomall/app/product/biz/dal/mysql"
	"github.com/Yzc216/gomall/app/product/biz/dal/redis"
	"github.com/Yzc216/gomall/app/product/biz/repo"
	product "github.com/Yzc216/gomall/rpc_gen/kitex_gen/product"
)

type BatchGetProductsService struct {
	ctx  context.Context
	repo *repo.CachedProductQuery
} // NewBatchGetProductsService new BatchGetProductsService
func NewBatchGetProductsService(ctx context.Context) *BatchGetProductsService {
	return &BatchGetProductsService{ctx: ctx, repo: repo.NewCachedProductQuery(mysql.DB, redis.RedisClient)}
}

// Run create note info
func (s *BatchGetProductsService) Run(req *product.BatchGetProductsReq) (resp *product.BatchGetProductsResp, err error) {
	// 去重
	ids := deduplicateIDs(req.GetIds())

	// 获取spu
	spuMap, err := s.repo.SPUQuery.BatchGetByIDs(s.ctx, ids)
	if err != nil {
		return nil, err
	}

	// 遍历原始 ID 列表，提取失败 ID
	var SPUs = make(map[uint64]*product.SPU, len(spuMap))
	var failedIDs []uint64
	for _, id := range ids {
		if _, ok := spuMap[id]; !ok {
			failedIDs = append(failedIDs, id)
			continue
		}

		SPUs[id], err = convertToProtoSPU(spuMap[id])
		if err != nil {
			failedIDs = append(failedIDs, id)
		}
	}

	return &product.BatchGetProductsResp{
		Products:  SPUs,
		FailedIds: failedIDs,
	}, nil
}

func deduplicateIDsInOrder(ids []uint64) []uint64 {
	seen := make(map[uint64]struct{}, len(ids))
	result := make([]uint64, 0, len(ids)) // 预分配原切片容量
	for _, id := range ids {
		if _, exists := seen[id]; !exists {
			seen[id] = struct{}{}
			result = append(result, id) // 保留首次出现顺序
		}
	}
	return result
}

func deduplicateIDs(ids []uint64) []uint64 {
	seen := make(map[uint64]struct{}, len(ids)) // 预分配足够空间
	for _, id := range ids {
		seen[id] = struct{}{} // 空结构体不占内存
	}
	result := make([]uint64, 0, len(seen))
	for id := range seen {
		result = append(result, id) // 无序
	}
	return result
}
