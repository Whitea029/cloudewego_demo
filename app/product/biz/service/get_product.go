package service

import (
	"context"

	"github.com/Whitea029/whmall/app/product/biz/dal/mysql"
	"github.com/Whitea029/whmall/app/product/biz/model"
	protuct "github.com/Whitea029/whmall/rpc_gen/kitex_gen/protuct"
	"github.com/cloudwego/kitex/pkg/kerrors"
)

type GetProductService struct {
	ctx context.Context
} // NewGetProductService new GetProductService
func NewGetProductService(ctx context.Context) *GetProductService {
	return &GetProductService{ctx: ctx}
}

// Run create note info
func (s *GetProductService) Run(req *protuct.GetProductReq) (resp *protuct.GetProductResp, err error) {
	if req.Id == 0 {
		return nil, kerrors.NewGRPCBizStatusError(2004001, "product id is required")
	}
	prodcutQuery := model.NewProductQuery(s.ctx, mysql.DB)
	p, err := prodcutQuery.GetById(uint(req.Id))
	if err != nil {
		return nil, err
	}
	return &protuct.GetProductResp{
		Product: &protuct.Product{
			Id:          uint32(p.ID),
			Name:        p.Name,
			Description: p.Description,
			Picture:     p.Picture,
			Price:       p.Price,
		},
	}, nil
}
