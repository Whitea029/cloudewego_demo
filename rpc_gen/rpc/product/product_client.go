package product

import (
	"context"
	protuct "github.com/Whitea029/whmall/rpc_gen/kitex_gen/protuct"

	"github.com/Whitea029/whmall/rpc_gen/kitex_gen/protuct/productcatalogservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

type RPCClient interface {
	KitexClient() productcatalogservice.Client
	Service() string
	ListProducts(ctx context.Context, Req *protuct.ListProductsReq, callOptions ...callopt.Option) (r *protuct.ListProductsResp, err error)
	GetProduct(ctx context.Context, Req *protuct.GetProductReq, callOptions ...callopt.Option) (r *protuct.GetProductResp, err error)
	SearchProducts(ctx context.Context, Req *protuct.SearchProductsReq, callOptions ...callopt.Option) (r *protuct.SearchProductsResp, err error)
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

func (c *clientImpl) ListProducts(ctx context.Context, Req *protuct.ListProductsReq, callOptions ...callopt.Option) (r *protuct.ListProductsResp, err error) {
	return c.kitexClient.ListProducts(ctx, Req, callOptions...)
}

func (c *clientImpl) GetProduct(ctx context.Context, Req *protuct.GetProductReq, callOptions ...callopt.Option) (r *protuct.GetProductResp, err error) {
	return c.kitexClient.GetProduct(ctx, Req, callOptions...)
}

func (c *clientImpl) SearchProducts(ctx context.Context, Req *protuct.SearchProductsReq, callOptions ...callopt.Option) (r *protuct.SearchProductsResp, err error) {
	return c.kitexClient.SearchProducts(ctx, Req, callOptions...)
}
