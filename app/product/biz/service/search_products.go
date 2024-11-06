package service

import (
	"context"
	protuct "github.com/Whitea029/whmall/rpc_gen/kitex_gen/protuct"
)

type SearchProductsService struct {
	ctx context.Context
} // NewSearchProductsService new SearchProductsService
func NewSearchProductsService(ctx context.Context) *SearchProductsService {
	return &SearchProductsService{ctx: ctx}
}

// Run create note info
func (s *SearchProductsService) Run(req *protuct.SearchProductsReq) (resp *protuct.SearchProductsResp, err error) {
	// Finish your business logic.

	return
}
