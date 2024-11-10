package service

import (
	"context"

	"github.com/Whitea029/whmall/app/product/biz/dal/mysql"
	"github.com/Whitea029/whmall/app/product/biz/model"
	product "github.com/Whitea029/whmall/rpc_gen/kitex_gen/product"
)

type SearchProductsService struct {
	ctx context.Context
} // NewSearchProductsService new SearchProductsService
func NewSearchProductsService(ctx context.Context) *SearchProductsService {
	return &SearchProductsService{ctx: ctx}
}

// Run create note info
func (s *SearchProductsService) Run(req *product.SearchProductsReq) (resp *product.SearchProductsResp, err error) {
	productQuery := model.NewProductQuery(s.ctx, mysql.DB)
	p, err := productQuery.SearchProduct(req.Query)
	if err != nil {
		return nil, err
	}
	results := make([]*product.Product, 0, len(p))
	for _, v := range p {
		results = append(results, &product.Product{
			Id:          uint32(v.ID),
			Name:        v.Name,
			Description: v.Description,
			Picture:     v.Picture,
			Price:       v.Price,
		})
	}
	return &product.SearchProductsResp{Results: results}, nil
}
