package service

import (
	"context"
	"testing"
	protuct "github.com/Whitea029/whmall/rpc_gen/kitex_gen/protuct"
)

func TestSearchProducts_Run(t *testing.T) {
	ctx := context.Background()
	s := NewSearchProductsService(ctx)
	// init req and assert value

	req := &protuct.SearchProductsReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
