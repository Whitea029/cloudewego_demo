package service

import (
	"context"

	"github.com/Whitea029/whmall/app/order/biz/dal/mysql"
	"github.com/Whitea029/whmall/app/order/biz/model"
	order "github.com/Whitea029/whmall/rpc_gen/kitex_gen/order"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PlaceOrderService struct {
	ctx context.Context
} // NewPlaceOrderService new PlaceOrderService
func NewPlaceOrderService(ctx context.Context) *PlaceOrderService {
	return &PlaceOrderService{ctx: ctx}
}

// Run create note info
func (s *PlaceOrderService) Run(req *order.PlaceOrderReq) (resp *order.PlaceOrderResp, err error) {
	if len(req.OrderItems) == 0 {
		return nil, kerrors.NewBizStatusError(500001, "order items is empty")
	}
	err = mysql.DB.Transaction(func(tx *gorm.DB) error {
		orderId := uuid.New().String()
		o := &model.Order{
			OrderId:      orderId,
			UserId:       req.UserId,
			UserCurrency: req.UserCurrency,
			Consignee: model.Consignee{
				Email: req.Email,
			},
		}
		if req.Address != nil {
			a := req.Address
			o.Consignee.StreetAddress = a.StreetAddress
			o.Consignee.City = a.City
			o.Consignee.State = a.State
			o.Consignee.Country = a.Country
		}
		if err = tx.Create(o).Error; err != nil {
			return err
		}
		var itemList []*model.OrderItem
		for _, v := range req.OrderItems {
			itemList = append(itemList, &model.OrderItem{
				OrderIdRefer: o.OrderId,
				ProductId:    v.Item.ProductId,
				Quantity:     uint32(v.Item.Quantity),
				Cost:         v.Cost,
			})
		}
		if err := tx.Create(&itemList).Error; err != nil {
			return err
		}
		resp = &order.PlaceOrderResp{
			Order: &order.OrderResult{
				OrderId: orderId,
			},
		}
		return nil
	})
	return
}
