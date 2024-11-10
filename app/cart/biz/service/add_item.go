package service

import (
	"context"

	"github.com/Whitea029/whmall/app/cart/biz/dal/mysql"
	"github.com/Whitea029/whmall/app/cart/biz/model"
	"github.com/Whitea029/whmall/app/cart/infra/rpc"
	cart "github.com/Whitea029/whmall/rpc_gen/kitex_gen/cart"
	"github.com/Whitea029/whmall/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/kerrors"
)

type AddItemService struct {
	ctx context.Context
} // NewAddItemService new AddItemService
func NewAddItemService(ctx context.Context) *AddItemService {
	return &AddItemService{ctx: ctx}
}

// Run create note info
func (s *AddItemService) Run(req *cart.AddItemReq) (resp *cart.AddItemResp, err error) {
	// Finish your business logic.
	productResp, err := rpc.ProductClient.GetProduct(s.ctx, &product.GetProductReq{Id: req.Item.ProductId})
	if err != nil {
		return nil, err
	}
	if productResp == nil || productResp.Product.Id == 0 {
		return nil, kerrors.NewBizStatusError(40004, "product not found")
	}
	cartItem := model.Cart{
		UserId:    req.UserId,
		ProductId: req.Item.ProductId,
		Qty:       uint32(req.Item.Quantity),
	}
	err = model.NewCartQuery(s.ctx, mysql.DB).AddCart(&cartItem)
	if err != nil {
		return nil, kerrors.NewBizStatusError(50000, err.Error())
	}
	return &cart.AddItemResp{}, nil
}
