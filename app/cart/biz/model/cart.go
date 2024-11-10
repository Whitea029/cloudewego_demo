package model

import (
	"context"
	"errors"

	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	UserId    uint32 `json:"user_id"`
	ProductId uint32 `json:"product_id"`
	Qty       uint32 `json:"qty"`
}

func (Cart) TableName() string {
	return "cart"
}

type CartQuery struct {
	ctx context.Context
	db  *gorm.DB
}

func (cq CartQuery) GetCartById(userId uint32) (cartList []*Cart, err error) {
	err = cq.db.Debug().WithContext(cq.ctx).Model(&Cart{}).Find(&cartList, "user_id = ?", userId).Error
	return
}

func (cq CartQuery) AddCart(c *Cart) (err error) {
	var find Cart
	err = cq.db.WithContext(cq.ctx).Model(&Cart{}).Where(&Cart{UserId: c.UserId, ProductId: c.ProductId}).First(&find).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return
	}
	if find.ID > 0 {
		err = cq.db.WithContext(cq.ctx).Model(&Cart{}).Where(&Cart{UserId: c.UserId, ProductId: c.ProductId}).UpdateColumn("qty", gorm.Expr("qty+?", c.Qty)).Error
		return
	}
	return cq.db.WithContext(cq.ctx).Model(&Cart{}).Create(c).Error
}

func (cq CartQuery) EmptyCart(userId uint32) (err error) {
	if userId == 0 {
		return errors.New("user_is is required")
	}
	return cq.db.WithContext(cq.ctx).Delete(&Cart{}, "user_id = ?", userId).Error
}

func NewCartQuery(ctx context.Context, db *gorm.DB) *CartQuery {
	return &CartQuery{ctx: ctx, db: db}
}
