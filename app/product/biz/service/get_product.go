package service

import (
	"context"
	protuct "github.com/Whitea029/whmall/rpc_gen/kitex_gen/protuct"
)

type GetProductService struct {
	ctx context.Context
} // NewGetProductService new GetProductService
func NewGetProductService(ctx context.Context) *GetProductService {
	return &GetProductService{ctx: ctx}
}

// Run create note info
func (s *GetProductService) Run(req *protuct.GetProductReq) (resp *protuct.GetProductResp, err error) {
	// Finish your business logic.

	return
}
