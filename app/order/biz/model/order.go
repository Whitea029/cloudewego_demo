package model

import (
	"context"

	"gorm.io/gorm"
)

type Consignee struct {
	Email         string
	StreetAddress string
	City          string
	State         string
	Country       string
	ZipCode       string
}

type Order struct {
	gorm.Model
	OrderId      string      `gorm:"type:varchar(100);uniqueIndex"`
	UserId       uint32      `gorm:"type:int(11)"`
	UserCurrency string      `gorm:"type:varchar(10)"`
	Consignee    Consignee   `gorm:"embedded"`
	OrderItems   []OrderItem `gorm:"foreignKey:OrderIdRefer;references:OrderId"`
}

func (o *Order) TableName() string {
	return "order"
}

func (oq OrderQuery) ListOrder(userId uint32) ([]*Order, error) {
	var orders []*Order
	err := oq.db.WithContext(*oq.ctx).Where("user_id = ?", userId).Preload("OrderItems").Find(&Order{}).Error
	if err != nil {
		return nil, err
	}
	return orders, nil
}

type OrderQuery struct {
	db  *gorm.DB
	ctx *context.Context
}

func NewOrderQuery(ctx context.Context, db *gorm.DB) *OrderQuery {
	return &OrderQuery{ctx: &ctx, db: db}
}
