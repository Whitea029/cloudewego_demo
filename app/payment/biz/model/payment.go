package model

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type PaymentLog struct {
	gorm.Model
	UserId        uint32    `json:"user_id"`
	OrderId       string    `json:"order_id"`
	TransactionId string    `json:"transaction_id"`
	Amount        float32   `json:"amount"`
	PayAt         time.Time `json:"pay_at"`
}

type PaymentLogQuery struct {
	db  gorm.DB
	ctx context.Context
}

func (p PaymentLog) TableName() string {
	return "payment_log"
}

func (qp *PaymentLogQuery) CreatePaymentLog(payment *PaymentLog) error {
	return qp.db.WithContext(qp.ctx).Model(&PaymentLog{}).Create(payment).Error
}

func NewPaymentLogQuery(db gorm.DB, ctx context.Context) *PaymentLogQuery {
	return &PaymentLogQuery{db: db, ctx: ctx}
}
