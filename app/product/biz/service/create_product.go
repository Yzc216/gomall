package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Yzc216/gomall/app/product/biz/dal/mysql"
	"github.com/Yzc216/gomall/app/product/biz/model"
	"github.com/Yzc216/gomall/app/product/infra/mq"
	utils "github.com/Yzc216/gomall/common/utils"
	"github.com/Yzc216/gomall/rpc_gen/kitex_gen/inventory"
	product "github.com/Yzc216/gomall/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/proto"
	"sort"
	"strings"
)

type CreateProductService struct {
	ctx          context.Context
	spuRepo      *model.SPURepo
	categoryRepo *model.CategoryRepo
} // NewCreateProductService new CreateProductService
func NewCreateProductService(ctx context.Context) *CreateProductService {
	return &CreateProductService{
		ctx:          ctx,
		spuRepo:      model.NewSPURepo(mysql.DB),
		categoryRepo: model.NewCategoryRepo(mysql.DB),
	}
}

// Run create note info
func (s *CreateProductService) Run(req *product.CreateProductReq) (resp *product.ProductResp, err error) {
	// 1. 参数校验
	if err = validateCreateRequest(req); err != nil {
		return nil, err
	}

	// 2. 数据转换
	spuModel, err := s.convertToSPUModel(req)
	if err != nil {
		return nil, err
	}

	// 3. 业务校验
	if err = s.validateBusinessRules(s.ctx, spuModel, req); err != nil {
		return nil, err
	}

	// 4. 持久化操作
	if err = s.spuRepo.Create(s.ctx, spuModel); err != nil {
		return nil, fmt.Errorf("创建商品失败: %w", err)
	}

	go func() {
		for _, sku := range spuModel.SKUs {
			data, _ := proto.Marshal(&inventory.ProductCreatedEvent{
				SkuId:        sku.ID,
				SkuName:      sku.Title,
				InitialStock: sku.Stock,
			})
			msg := &nats.Msg{Subject: "inventory", Data: data, Header: make(nats.Header)}

			// otel inject
			//otel.GetTextMapPropagator().Inject(s.ctx, propagation.HeaderCarrier(msg.Header))

			err = mq.Nc.PublishMsg(msg)
			if err != nil {
				klog.Error(err.Error())
			}
		}
	}()

	// 5. 返回结果转换
	return &product.ProductResp{Success: true}, nil
}

// 参数校验示例
func validateCreateRequest(req *product.CreateProductReq) error {
	if req.GetBasicInfo().GetTitle() == "" {
		return errors.New("商品标题不能为空")
	}
	if len(req.GetMedia().GetMainImages()) == 0 {
		return errors.New("至少需要一张主图")
	}
	if req.GetCategoryRelation().GetCategoryIds() == nil {
		return errors.New("商品分类不应为空")
	}
	if req.GetBasicInfo().GetShopId() == 0 {
		return errors.New("店铺编号不能为空")
	}
	// 更多校验规则...
	return nil
}

// 业务规则校验
func (s *CreateProductService) validateBusinessRules(ctx context.Context, spu *model.SPU, req *product.CreateProductReq) error {
	// 校验分类是否存在
	if exist, err := s.categoryRepo.ExistByIDs(ctx, req.CategoryRelation.CategoryIds); err != nil || !exist {
		return errors.New("包含无效的商品分类")
	}

	//// 校验品牌是否存在
	//if _, err := s.brandRepo.GetBrand(ctx, spu.BrandID); err != nil {
	//	return status.Error(codes.InvalidArgument, "品牌不存在")
	//}
	//
	//// 校验店铺合法性
	//if _, err := s.shopRepo.GetShop(ctx, spu.ShopID); err != nil {
	//	return status.Error(codes.PermissionDenied, "无权操作该店铺")
	//}

	// 校验SKU规格唯一性
	specCombinations := make(map[string]bool)
	for _, sku := range req.GetSKUs() {
		keys := make([]string, 0)
		for k, v := range sku.GetSpecs() {
			keys = append(keys, fmt.Sprintf("%s:%s", k, v))
		}
		sort.Strings(keys)
		combination := strings.Join(keys, "|")
		if specCombinations[combination] {
			return errors.New("存在重复的SKU规格组合")
		}
		specCombinations[combination] = true
	}
	return nil
}

// 请求转模型示例
func (s *CreateProductService) convertToSPUModel(req *product.CreateProductReq) (*model.SPU, error) {
	spu := &model.SPU{
		ID:          utils.GenID(),
		Title:       req.BasicInfo.Title,
		SubTitle:    req.BasicInfo.SubTitle,
		ShopID:      req.BasicInfo.ShopId,
		Brand:       req.BasicInfo.Brand,
		MainImages:  req.Media.MainImages,
		Video:       req.Media.VideoUrl,
		Description: req.BasicInfo.Description,
		Status:      int8(req.BasicInfo.Status),
		SKUs:        make([]model.SKU, 0),
	}

	// 处理分类关联
	for _, cid := range req.CategoryRelation.CategoryIds {
		spu.Categories = append(spu.Categories, model.Category{ID: cid})
	}

	// 处理SKU
	for _, skuReq := range req.SKUs {
		specs, _ := json.Marshal(skuReq.Specs)
		spu.SKUs = append(spu.SKUs, model.SKU{
			Title:    skuReq.Title,
			Price:    skuReq.Price,
			Stock:    skuReq.Stock,
			Specs:    specs,
			IsActive: true, // 默认激活
		})
	}

	return spu, nil
}
