package main

import (
	"context"
	"github.com/Whitea029/whmall/app/product/biz/service"
	"github.com/Whitea029/whmall/rpc_gen/kitex_gen/protuct"
)

type ProductCatalogServiceImpl struct{}

// ListProducts implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) ListProducts(ctx context.Context, req *protuct.ListProductsReq) (resp *protuct.ListProductsResp, err error) {
	resp, err = service.NewListProductsService(ctx).Run(req)

	return resp, err
}

// GetProduct implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) GetProduct(ctx context.Context, req *protuct.GetProductReq) (resp *protuct.GetProductResp, err error) {
	resp, err = service.NewGetProductService(ctx).Run(req)

	return resp, err
}

// SearchProducts implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) SearchProducts(ctx context.Context, req *protuct.SearchProductsReq) (resp *protuct.SearchProductsResp, err error) {
	resp, err = service.NewSearchProductsService(ctx).Run(req)

	return resp, err
}
