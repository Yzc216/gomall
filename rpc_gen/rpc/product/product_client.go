package product

import (
	"context"
	common "github.com/Yzc216/gomall/rpc_gen/kitex_gen/common"
	product "github.com/Yzc216/gomall/rpc_gen/kitex_gen/product"

	"github.com/Yzc216/gomall/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

type RPCClient interface {
	KitexClient() productcatalogservice.Client
	Service() string
	CreateProduct(ctx context.Context, Req *product.CreateProductReq, callOptions ...callopt.Option) (r *product.ProductResp, err error)
	UpdateProduct(ctx context.Context, Req *product.UpdateProductReq, callOptions ...callopt.Option) (r *product.ProductResp, err error)
	DeleteProduct(ctx context.Context, Req *product.DeleteProductReq, callOptions ...callopt.Option) (r *common.Empty, err error)
	ListProducts(ctx context.Context, Req *product.ListProductsReq, callOptions ...callopt.Option) (r *product.ListProductsResp, err error)
	GetProduct(ctx context.Context, Req *product.GetProductReq, callOptions ...callopt.Option) (r *product.GetProductResp, err error)
	BatchGetProducts(ctx context.Context, Req *product.BatchGetProductsReq, callOptions ...callopt.Option) (r *product.BatchGetProductsResp, err error)
	SearchProducts(ctx context.Context, Req *product.SearchProductsReq, callOptions ...callopt.Option) (r *product.SearchProductsResp, err error)
	CreateCategory(ctx context.Context, Req *product.CreateCategoryReq, callOptions ...callopt.Option) (r *product.Category, err error)
	UpdateCategory(ctx context.Context, Req *product.UpdateCategoryReq, callOptions ...callopt.Option) (r *product.Category, err error)
	DeleteCategory(ctx context.Context, Req *product.DeleteCategoryReq, callOptions ...callopt.Option) (r *common.Empty, err error)
	ListCategories(ctx context.Context, Req *product.ListCategoriesReq, callOptions ...callopt.Option) (r *product.CategoryNode, err error)
	GetCategoryTree(ctx context.Context, Req *product.GetCategoryTreeReq, callOptions ...callopt.Option) (r *product.CategoryTreeResp, err error)
}

func NewRPCClient(dstService string, opts ...client.Option) (RPCClient, error) {
	kitexClient, err := productcatalogservice.NewClient(dstService, opts...)
	if err != nil {
		return nil, err
	}
	cli := &clientImpl{
		service:     dstService,
		kitexClient: kitexClient,
	}

	return cli, nil
}

type clientImpl struct {
	service     string
	kitexClient productcatalogservice.Client
}

func (c *clientImpl) Service() string {
	return c.service
}

func (c *clientImpl) KitexClient() productcatalogservice.Client {
	return c.kitexClient
}

func (c *clientImpl) CreateProduct(ctx context.Context, Req *product.CreateProductReq, callOptions ...callopt.Option) (r *product.ProductResp, err error) {
	return c.kitexClient.CreateProduct(ctx, Req, callOptions...)
}

func (c *clientImpl) UpdateProduct(ctx context.Context, Req *product.UpdateProductReq, callOptions ...callopt.Option) (r *product.ProductResp, err error) {
	return c.kitexClient.UpdateProduct(ctx, Req, callOptions...)
}

func (c *clientImpl) DeleteProduct(ctx context.Context, Req *product.DeleteProductReq, callOptions ...callopt.Option) (r *common.Empty, err error) {
	return c.kitexClient.DeleteProduct(ctx, Req, callOptions...)
}

func (c *clientImpl) ListProducts(ctx context.Context, Req *product.ListProductsReq, callOptions ...callopt.Option) (r *product.ListProductsResp, err error) {
	return c.kitexClient.ListProducts(ctx, Req, callOptions...)
}

func (c *clientImpl) GetProduct(ctx context.Context, Req *product.GetProductReq, callOptions ...callopt.Option) (r *product.GetProductResp, err error) {
	return c.kitexClient.GetProduct(ctx, Req, callOptions...)
}

func (c *clientImpl) BatchGetProducts(ctx context.Context, Req *product.BatchGetProductsReq, callOptions ...callopt.Option) (r *product.BatchGetProductsResp, err error) {
	return c.kitexClient.BatchGetProducts(ctx, Req, callOptions...)
}

func (c *clientImpl) SearchProducts(ctx context.Context, Req *product.SearchProductsReq, callOptions ...callopt.Option) (r *product.SearchProductsResp, err error) {
	return c.kitexClient.SearchProducts(ctx, Req, callOptions...)
}

func (c *clientImpl) CreateCategory(ctx context.Context, Req *product.CreateCategoryReq, callOptions ...callopt.Option) (r *product.Category, err error) {
	return c.kitexClient.CreateCategory(ctx, Req, callOptions...)
}

func (c *clientImpl) UpdateCategory(ctx context.Context, Req *product.UpdateCategoryReq, callOptions ...callopt.Option) (r *product.Category, err error) {
	return c.kitexClient.UpdateCategory(ctx, Req, callOptions...)
}

func (c *clientImpl) DeleteCategory(ctx context.Context, Req *product.DeleteCategoryReq, callOptions ...callopt.Option) (r *common.Empty, err error) {
	return c.kitexClient.DeleteCategory(ctx, Req, callOptions...)
}

func (c *clientImpl) ListCategories(ctx context.Context, Req *product.ListCategoriesReq, callOptions ...callopt.Option) (r *product.CategoryNode, err error) {
	return c.kitexClient.ListCategories(ctx, Req, callOptions...)
}

func (c *clientImpl) GetCategoryTree(ctx context.Context, Req *product.GetCategoryTreeReq, callOptions ...callopt.Option) (r *product.CategoryTreeResp, err error) {
	return c.kitexClient.GetCategoryTree(ctx, Req, callOptions...)
}
