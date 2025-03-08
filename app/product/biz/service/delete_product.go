package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/Yzc216/gomall/app/product/biz/dal/mysql"
	"github.com/Yzc216/gomall/app/product/biz/model"
	"github.com/Yzc216/gomall/app/product/biz/types"
	"github.com/Yzc216/gomall/rpc_gen/kitex_gen/common"
	product "github.com/Yzc216/gomall/rpc_gen/kitex_gen/product"
)

type DeleteProductService struct {
	ctx  context.Context
	repo *model.SPURepo
} // NewDeleteProductService new DeleteProductService
func NewDeleteProductService(ctx context.Context) *DeleteProductService {
	return &DeleteProductService{ctx: ctx, repo: model.NewSPURepo(mysql.DB)}
}

// Run create note info
func (s *DeleteProductService) Run(req *product.DeleteProductReq) (resp *common.Empty, err error) {
	if req.GetId() == 0 {
		return nil, errors.New("spuId is required")
	}
	if !req.Force {
		hasSKUs, err := s.repo.GetSKUCount(s.ctx, req.Id)
		if err != nil {
			return nil, fmt.Errorf("check SKU associations failed: %w", err)
		}
		if hasSKUs > 0 {
			return nil, types.ErrHasAssociatedSKUs
		}
	}

	// 2. 执行删除操作（包含事务管理）
	if err := s.repo.Delete(s.ctx, req.Id); err != nil {
		return nil, fmt.Errorf("delete SPU failed: %w", err)
	}

	// 3. 发送删除事件（如清理缓存、更新搜索索引等）
	//go func() {
	//	data, _ := proto.Marshal(&inventory.ProductDeleteEvent{
	//
	//	})
	//	msg := &nats.Msg{Subject: "inventory", Data: data, Header: make(nats.Header)}
	//
	//	// otel inject
	//	//otel.GetTextMapPropagator().Inject(s.ctx, propagation.HeaderCarrier(msg.Header))
	//
	//	err = mq.Nc.PublishMsg(msg)
	//	if err != nil {
	//		klog.Error(err.Error())
	//	}
	//	// go s.publishDeleteEvent(req.Id)
	//}()

	return nil, nil
}
