package service

import (
	"context"

	"github.com/Whitea029/whmall/app/order/biz/dal/mysql"
	"github.com/Whitea029/whmall/app/order/biz/model"
	"github.com/Whitea029/whmall/rpc_gen/kitex_gen/cart"
	order "github.com/Whitea029/whmall/rpc_gen/kitex_gen/order"
	"github.com/cloudwego/kitex/pkg/kerrors"
)

type ListOrderService struct {
	ctx context.Context
} // NewListOrderService new ListOrderService
func NewListOrderService(ctx context.Context) *ListOrderService {
	return &ListOrderService{ctx: ctx}
}

// Run create note info
func (s *ListOrderService) Run(req *order.ListOrderReq) (resp *order.ListOrderResp, err error) {
	list, err := model.NewOrderQuery(s.ctx, mysql.DB).ListOrder(req.UserId)
	if err != nil {
		return nil, kerrors.NewBizStatusError(500001, err.Error())
	}
	var orders []*order.Order
	for _, o := range list {
		var items []*order.OrderItem
		for _, oi := range o.OrderItems {
			items = append(items, &order.OrderItem{
				Item: &cart.CartItem{
					ProductId: oi.ProductId,
					Quantity:  int32(oi.Quantity),
				},
				Cost: oi.Cost,
			})
		}
		orders = append(orders, &order.Order{
			OrderId:      o.OrderId,
			UserId:       o.UserId,
			UserCurrency: o.UserCurrency,
			Email:        o.Consignee.Email,
			Address: &order.Address{
				StreetAddress: o.Consignee.StreetAddress,
				City:          o.Consignee.City,
				State:         o.Consignee.State,
				Country:       o.Consignee.Country,
			},
			OrderItems: items,
		})
	}
	resp = &order.ListOrderResp{
		Orders: orders,
	}
	return
}
