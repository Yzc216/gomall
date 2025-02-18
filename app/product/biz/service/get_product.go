package service

import (
	"context"
	product "github.com/Yzc216/gomall/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/kerrors"
)

type GetProductService struct {
	ctx context.Context
} // NewGetProductService new GetProductService
func NewGetProductService(ctx context.Context) *GetProductService {
	return &GetProductService{ctx: ctx}
}

// Run create note info
func (s *GetProductService) Run(req *product.GetProductReq) (resp *product.GetProductResp, err error) {
	if req.Id == 0 {
		return nil, kerrors.NewGRPCBizStatusError(2004001, "product id is required")
	}

	//productQuery := model.NewProductQuery(s.ctx, mysql.DB)
	//
	//productRes, err := productQuery.GetById(int(req.Id))
	//if err != nil {
	//	return nil, err
	//}
	//
	//return &product.GetProductResp{
	//	Product: &product.Product{
	//		Id:          uint32(productRes.ID),
	//		Name:        productRes.Name,
	//		Price:       productRes.Price,
	//		Picture:     productRes.Picture,
	//		Description: productRes.Description,
	//	},
	//}, nil
	return &product.GetProductResp{}, nil
}
