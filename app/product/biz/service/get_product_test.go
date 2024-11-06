package service

import (
	"context"
	"testing"
	protuct "github.com/Whitea029/whmall/rpc_gen/kitex_gen/protuct"
)

func TestGetProduct_Run(t *testing.T) {
	ctx := context.Background()
	s := NewGetProductService(ctx)
	// init req and assert value

	req := &protuct.GetProductReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
