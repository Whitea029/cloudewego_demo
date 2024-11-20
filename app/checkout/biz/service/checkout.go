package service

import (
	"context"
	"strconv"

	"github.com/Whitea029/whmall/app/checkout/infra/mq"
	"github.com/Whitea029/whmall/app/checkout/infra/rpc"
	"github.com/Whitea029/whmall/rpc_gen/kitex_gen/cart"
	checkout "github.com/Whitea029/whmall/rpc_gen/kitex_gen/checkout"
	"github.com/Whitea029/whmall/rpc_gen/kitex_gen/email"
	rpcorder "github.com/Whitea029/whmall/rpc_gen/kitex_gen/order"
	"github.com/Whitea029/whmall/rpc_gen/kitex_gen/payment"
	"github.com/Whitea029/whmall/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/nats-io/nats.go"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"google.golang.org/protobuf/proto"
)

type CheckoutService struct {
	ctx context.Context
} // NewCheckoutService new CheckoutService
func NewCheckoutService(ctx context.Context) *CheckoutService {
	return &CheckoutService{ctx: ctx}
}

// Run create note info
func (s *CheckoutService) Run(req *checkout.CheckoutReq) (resp *checkout.CheckoutResp, err error) {
	// get cart
	cartResp, err := rpc.CartCLient.GetCart(s.ctx, &cart.GetCartReq{
		UserId: req.UserId,
	})
	if err != nil {
		return nil, kerrors.NewGRPCBizStatusError(5005001, err.Error())
	}
	if cartResp == nil || cartResp.Cart.Items == nil {
		return nil, kerrors.NewGRPCBizStatusError(5004001, "cart is empty")
	}

	var (
		total float32
		oi    []*rpcorder.OrderItem
	)

	for _, item := range cartResp.Cart.Items {
		// get product
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

		oi = append(oi, &rpcorder.OrderItem{
			Item: &cart.CartItem{
				ProductId: item.ProductId,
				Quantity:  item.Quantity,
			},
			Cost: p.Price * float32(item.Quantity),
		})
	}

	var orderId string
	zipCodeInt, _ := strconv.Atoi(req.Address.ZipCode)

	// place order
	orderResp, err := rpc.OrderClient.PlaceOrder(s.ctx, &rpcorder.PlaceOrderReq{
		UserId: req.UserId,
		Email:  req.Email,
		Address: &rpcorder.Address{
			StreetAddress: req.Address.StreetAddress,
			City:          req.Address.City,
			State:         req.Address.State,
			ZipCode:       int32(zipCodeInt),
			Country:       req.Address.Country,
		},
		OrderItems: oi,
	})
	if err != nil {
		return nil, kerrors.NewGRPCBizStatusError(5005002, err.Error())
	}
	if orderResp != nil && orderResp.Order != nil {
		orderId = orderResp.Order.OrderId
	}

	// charge
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
	data, _ := proto.Marshal(&email.EmailReq{
		From:        "whitea0029@gmail.com",
		To:          req.Email,
		ContentType: "text/plain",
		Subject:     "You have just created an order in Pomelo shop",
		Content:     "You have just created an order in Pomelo shop, your order id is " + orderId,
	})
	msg := &nats.Msg{Subject: "email", Data: data, Header: make(nats.Header)}
	otel.GetTextMapPropagator().Inject(s.ctx, propagation.HeaderCarrier(msg.Header))
	_ = mq.Nc.PublishMsg(msg)
	klog.Info("paymentResp: ", paymentResp)
	return &checkout.CheckoutResp{OrderId: orderId, TransactionId: paymentResp.TransactionId}, nil
}
