package service

import (
	"context"

	"github.com/Whitea029/whmall/app/checkout/infra/rpc"
	"github.com/Whitea029/whmall/rpc_gen/kitex_gen/cart"
	checkout "github.com/Whitea029/whmall/rpc_gen/kitex_gen/checkout"
	"github.com/Whitea029/whmall/rpc_gen/kitex_gen/payment"
	"github.com/Whitea029/whmall/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/google/uuid"
)

type CheckoutService struct {
	ctx context.Context
} // NewCheckoutService new CheckoutService
func NewCheckoutService(ctx context.Context) *CheckoutService {
	return &CheckoutService{ctx: ctx}
}

// Run create note info
func (s *CheckoutService) Run(req *checkout.CheckoutReq) (resp *checkout.CheckoutResp, err error) {
	cartResp, err := rpc.CartCLient.GetCart(s.ctx, &cart.GetCartReq{
		UserId: req.UserId,
	})
	if err != nil {
		return nil, kerrors.NewGRPCBizStatusError(5005001, err.Error())
	}
	if cartResp == nil || cartResp.Cart.Items == nil {
		return nil, kerrors.NewGRPCBizStatusError(5004001, "cart is empty")
	}
	var total float32
	for _, item := range cartResp.Cart.Items {
		productResp, err := rpc.ProductClient.GetProduct(s.ctx, &product.GetProductReq{
			Id: item.ProductId,
		})
		if err != nil {
			return nil, err
		}
		if productResp == nil {
			continue
		}
		p := productResp.Product
		total += p.Price * float32(item.Quantity)
	}
	var orderId string
	u, _ := uuid.NewRandom()
	orderId = u.String()
	payReq := &payment.ChargeReq{
		Amount:  total,
		OrderId: orderId,
		UserId:  req.UserId,
		CreditCard: &payment.CreditCardInfo{
			CreditCardNumber:          req.CreditCard.CreditCardNumber,
			CreditCardExpirationMonth: req.CreditCard.CreditCardExpirationMonth,
			CreditCardExpirationYear:  req.CreditCard.CreditCardExpirationYear,
			CreditCardCvv:             req.CreditCard.CreditCardCvv,
		},
	}
	_, err = rpc.CartCLient.EmptyCart(s.ctx, &cart.EmptyCartReq{UserId: req.UserId})
	if err != nil {
		klog.Error(err.Error())
	}
	paymentResp, err := rpc.PaymentClient.Charge(s.ctx, payReq)
	if err != nil {
		return nil, kerrors.NewGRPCBizStatusError(5005002, err.Error())
	}
	klog.Info("paymentResp: ", paymentResp)
	return &checkout.CheckoutResp{OrderId: orderId, TransactionId: paymentResp.TransactionId}, nil
}
