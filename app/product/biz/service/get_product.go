package service

import (
	"context"
	"github.com/Yzc216/gomall/app/product/biz/dal/mysql"
	"github.com/Yzc216/gomall/app/product/biz/model"
	product "github.com/Yzc216/gomall/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/kerrors"
)

type GetProductService struct {
	ctx  context.Context
	repo *model.SPURepo
} // NewGetProductService new GetProductService
func NewGetProductService(ctx context.Context) *GetProductService {
	return &GetProductService{ctx: ctx, repo: model.NewSPURepo(mysql.DB)}
}

// Run create note info
func (s *GetProductService) Run(req *product.GetProductReq) (resp *product.GetProductResp, err error) {
	if req.Id == 0 {
		return nil, kerrors.NewBizStatusError(40000, "product id is required")
	}
	spu, err := s.repo.GetByID(s.ctx, req.Id)
	if err != nil {
		return nil, err
	}
	protoSPU, err := convertToProtoSPU(spu)
	if err != nil {
		return nil, err
	}

	return &product.GetProductResp{Product: protoSPU}, nil
}
