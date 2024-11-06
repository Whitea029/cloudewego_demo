package service

import (
	"context"

	"github.com/Whitea029/whmall/app/product/biz/dal/mysql"
	"github.com/Whitea029/whmall/app/product/biz/model"
	protuct "github.com/Whitea029/whmall/rpc_gen/kitex_gen/protuct"
)

type ListProductsService struct {
	ctx context.Context
} // NewListProductsService new ListProductsService
func NewListProductsService(ctx context.Context) *ListProductsService {
	return &ListProductsService{ctx: ctx}
}

// Run create note info
func (s *ListProductsService) Run(req *protuct.ListProductsReq) (resp *protuct.ListProductsResp, err error) {
	categoryQuery := model.NewCategoryQuery(s.ctx, mysql.DB)
	c, err := categoryQuery.GetProductsByCategoryName(req.CategoryName)
	if err != nil {
		return nil, err
	}
	resp = &protuct.ListProductsResp{}
	for _, v1 := range c {
		for _, v2 := range v1.Products {
			resp.Products = append(resp.Products, &protuct.Product{
				Id:          uint32(v2.ID),
				Name:        v2.Name,
				Description: v2.Description,
				Picture:     v2.Picture,
				Price:       v2.Price,
			})
		}
	}
	return resp, nil
}
