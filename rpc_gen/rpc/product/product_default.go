package product

import (
	"context"
	common "github.com/Yzc216/gomall/rpc_gen/kitex_gen/common"
	product "github.com/Yzc216/gomall/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
)

func CreateProduct(ctx context.Context, req *product.CreateProductReq, callOptions ...callopt.Option) (resp *product.ProductResp, err error) {
	resp, err = defaultClient.CreateProduct(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "CreateProduct call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func UpdateProduct(ctx context.Context, req *product.UpdateProductReq, callOptions ...callopt.Option) (resp *product.ProductResp, err error) {
	resp, err = defaultClient.UpdateProduct(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "UpdateProduct call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func DeleteProduct(ctx context.Context, req *product.DeleteProductReq, callOptions ...callopt.Option) (resp *common.Empty, err error) {
	resp, err = defaultClient.DeleteProduct(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "DeleteProduct call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func ListProducts(ctx context.Context, req *product.ListProductsReq, callOptions ...callopt.Option) (resp *product.ListProductsResp, err error) {
	resp, err = defaultClient.ListProducts(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "ListProducts call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func GetProduct(ctx context.Context, req *product.GetProductReq, callOptions ...callopt.Option) (resp *product.GetProductResp, err error) {
	resp, err = defaultClient.GetProduct(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "GetProduct call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func BatchGetProducts(ctx context.Context, req *product.BatchGetProductsReq, callOptions ...callopt.Option) (resp *product.BatchGetProductsResp, err error) {
	resp, err = defaultClient.BatchGetProducts(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "BatchGetProducts call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func SearchProducts(ctx context.Context, req *product.SearchProductsReq, callOptions ...callopt.Option) (resp *product.SearchProductsResp, err error) {
	resp, err = defaultClient.SearchProducts(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "SearchProducts call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func CreateCategory(ctx context.Context, req *product.CreateCategoryReq, callOptions ...callopt.Option) (resp *product.Category, err error) {
	resp, err = defaultClient.CreateCategory(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "CreateCategory call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func UpdateCategory(ctx context.Context, req *product.UpdateCategoryReq, callOptions ...callopt.Option) (resp *product.Category, err error) {
	resp, err = defaultClient.UpdateCategory(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "UpdateCategory call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func DeleteCategory(ctx context.Context, req *product.DeleteCategoryReq, callOptions ...callopt.Option) (resp *common.Empty, err error) {
	resp, err = defaultClient.DeleteCategory(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "DeleteCategory call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func ListCategories(ctx context.Context, req *product.ListCategoriesReq, callOptions ...callopt.Option) (resp *product.CategoryNode, err error) {
	resp, err = defaultClient.ListCategories(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "ListCategories call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func GetCategoryTree(ctx context.Context, req *product.GetCategoryTreeReq, callOptions ...callopt.Option) (resp *product.CategoryTreeResp, err error) {
	resp, err = defaultClient.GetCategoryTree(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "GetCategoryTree call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
