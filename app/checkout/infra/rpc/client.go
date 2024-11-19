package rpc

import (
	"sync"

	"github.com/Whitea029/whmall/app/checkout/conf"
	checkoutUtils "github.com/Whitea029/whmall/app/checkout/utils"
	"github.com/Whitea029/whmall/common/clientsuite"
	"github.com/Whitea029/whmall/rpc_gen/kitex_gen/cart/cartservice"
	"github.com/Whitea029/whmall/rpc_gen/kitex_gen/order/orderservice"
	"github.com/Whitea029/whmall/rpc_gen/kitex_gen/payment/paymentservice"
	"github.com/Whitea029/whmall/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/cloudwego/kitex/client"
)

var (
	CartCLient    cartservice.Client
	ProductClient productcatalogservice.Client
	PaymentClient paymentservice.Client
	OrderClient   orderservice.Client
	once          sync.Once
	ServiceName   = conf.GetConf().Kitex.Service
	RegistryAddr  = conf.GetConf().Registry.RegistryAddress[0]
	err           error
)

func InitClient() {
	once.Do(func() {
		initCartClient()
		initProductClient()
		initPaymentClient()
		initOrderClient()
	})
}

func initOrderClient() {
	opts := []client.Option{
		client.WithSuite(clientsuite.CommonClientSuite{
			CurrentServiceName: ServiceName,
			RegistryAddr:       RegistryAddr,
		}),
	}
	OrderClient, err = orderservice.NewClient("order", opts...)
	checkoutUtils.MustHandleError(err)
}

func initProductClient() {
	opts := []client.Option{
		client.WithSuite(clientsuite.CommonClientSuite{
			CurrentServiceName: ServiceName,
			RegistryAddr:       RegistryAddr,
		}),
	}
	ProductClient, err = productcatalogservice.NewClient("product", opts...)
	checkoutUtils.MustHandleError(err)
}

func initPaymentClient() {
	opts := []client.Option{
		client.WithSuite(clientsuite.CommonClientSuite{
			CurrentServiceName: ServiceName,
			RegistryAddr:       RegistryAddr,
		}),
	}
	PaymentClient, err = paymentservice.NewClient("payment", opts...)
	checkoutUtils.MustHandleError(err)
}

func initCartClient() {
	opts := []client.Option{
		client.WithSuite(clientsuite.CommonClientSuite{
			CurrentServiceName: ServiceName,
			RegistryAddr:       RegistryAddr,
		}),
	}
	CartCLient, err = cartservice.NewClient("cart", opts...)
	checkoutUtils.MustHandleError(err)
}
