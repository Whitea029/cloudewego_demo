package service

import (
	"context"
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
	// Finish your business logic.

	return
}
