package rpc

import (
	"sync"

	"github.com/Whitea029/whmall/app/cart/conf"
	cartUtils "github.com/Whitea029/whmall/app/cart/utils"
	"github.com/Whitea029/whmall/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/Whitea029/whmall/rpc_gen/kitex_gen/user/userservice"
	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
)

var (
	UserClient    userservice.Client
	ProductClient productcatalogservice.Client
	once          sync.Once
)

func Init() {
	once.Do(func() {
		initProductClient()
	})
}

func initProductClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	cartUtils.MustHandleError(err)
	opts = append(opts, client.WithResolver(r))

	ProductClient, err = productcatalogservice.NewClient("product", opts...)
	cartUtils.MustHandleError(err)
}
