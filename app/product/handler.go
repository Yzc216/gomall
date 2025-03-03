package main

import (
	"context"
	"github.com/Yzc216/gomall/app/product/biz/service"
	"github.com/Yzc216/gomall/rpc_gen/kitex_gen/common"
	"github.com/Yzc216/gomall/rpc_gen/kitex_gen/product"
)

// ProductCatalogServiceImpl implements the last service interface defined in the IDL.
type ProductCatalogServiceImpl struct{}

// ListProducts implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) ListProducts(ctx context.Context, req *product.ListProductsReq) (resp *product.ListProductsResp, err error) {
	resp, err = service.NewListProductsService(ctx).Run(req)

	return resp, err
}

// GetProduct implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) GetProduct(ctx context.Context, req *product.GetProductReq) (resp *product.GetProductResp, err error) {
	resp, err = service.NewGetProductService(ctx).Run(req)

	return resp, err
}

// SearchProducts implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) SearchProducts(ctx context.Context, req *product.SearchProductsReq) (resp *product.SearchProductsResp, err error) {
	resp, err = service.NewSearchProductsService(ctx).Run(req)

	return resp, err
}

// CreateProduct implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) CreateProduct(ctx context.Context, req *product.CreateProductReq) (resp *product.ProductResp, err error) {
	resp, err = service.NewCreateProductService(ctx).Run(req)

	return resp, err
}

// UpdateProduct implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) UpdateProduct(ctx context.Context, req *product.UpdateProductReq) (resp *product.ProductResp, err error) {
	resp, err = service.NewUpdateProductService(ctx).Run(req)

	return resp, err
}

// DeleteProduct implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) DeleteProduct(ctx context.Context, req *product.DeleteProductReq) (resp *common.Empty, err error) {
	resp, err = service.NewDeleteProductService(ctx).Run(req)

	return resp, err
}

// ListCategories implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) ListCategories(ctx context.Context, req *product.ListCategoriesReq) (resp *product.CategoryNode, err error) {
	resp, err = service.NewListCategoriesService(ctx).Run(req)

	return resp, err
}

// CreateCategory implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) CreateCategory(ctx context.Context, req *product.CreateCategoryReq) (resp *product.Category, err error) {
	resp, err = service.NewCreateCategoryService(ctx).Run(req)

	return resp, err
}

// UpdateCategory implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) UpdateCategory(ctx context.Context, req *product.UpdateCategoryReq) (resp *product.Category, err error) {
	resp, err = service.NewUpdateCategoryService(ctx).Run(req)

	return resp, err
}

// DeleteCategory implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) DeleteCategory(ctx context.Context, req *product.DeleteCategoryReq) (resp *common.Empty, err error) {
	resp, err = service.NewDeleteCategoryService(ctx).Run(req)

	return resp, err
}

// GetCategoryTree implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) GetCategoryTree(ctx context.Context, req *product.GetCategoryTreeReq) (resp *product.CategoryTreeResp, err error) {
	resp, err = service.NewGetCategoryTreeService(ctx).Run(req)

	return resp, err
}

// BatchGetProducts implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) BatchGetProducts(ctx context.Context, req *product.BatchGetProductsReq) (resp *product.BatchGetProductsResp, err error) {
	resp, err = service.NewBatchGetProductsService(ctx).Run(req)

	return resp, err
}
