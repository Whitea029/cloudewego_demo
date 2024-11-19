PHONY: gen-frontend
gen-frontend:
	@cd app/frontend && cwgo server --type HTTP --idl ../../idl/frontend/checkout_page.proto --service frontend --module github.com/Whitea029/whmall/app/frontend -I ../../idl

PHONY: gen-user
gen-user:
	@cd rpc_gen && cwgo client --type RPC --service user --module github.com/Whitea029/whmall/rpc_gen --I ../idl --idl ../idl/user.proto
	@cd app/user && cwgo server --type RPC --service user --module github.com/Whitea029/whmall/app/user --pass "-use github.com/Whitea029/whmall/rpc_gen/kitex_gen" -I ../../idl --idl ../../idl/user.proto
	

PHONY: gen-product
gen-product:
	@cd rpc_gen && cwgo client --type RPC --service product --module github.com/Whitea029/whmall/rpc_gen --I ../idl --idl ../idl/product.proto
	@cd app/product && cwgo server --type RPC --service product --module github.com/Whitea029/whmall/app/product --pass "-use github.com/Whitea029/whmall/rpc_gen/kitex_gen" -I ../../idl --idl ../../idl/product.proto

PHONY: gen-cart
gen-cart:
	@cd rpc_gen && cwgo client --type RPC --service cart --module github.com/Whitea029/whmall/rpc_gen --I ../idl --idl ../idl/cart.proto
	@cd app/cart && cwgo server --type RPC --service cart --module github.com/Whitea029/whmall/app/cart --pass "-use github.com/Whitea029/whmall/rpc_gen/kitex_gen" -I ../../idl --idl ../../idl/cart.proto

PHONY: gen-payment
gen-payment:
	@cd rpc_gen && cwgo client --type RPC --service payment --module github.com/Whitea029/whmall/rpc_gen --I ../idl --idl ../idl/payment.proto
	@cd app/payment && cwgo server --type RPC --service payment --module github.com/Whitea029/whmall/app/payment --pass "-use github.com/Whitea029/whmall/rpc_gen/kitex_gen" -I ../../idl --idl ../../idl/payment.proto

PHONY: gen-checkout
gen-checkout:
	@cd rpc_gen && cwgo client --type RPC --service checkout --module github.com/Whitea029/whmall/rpc_gen --I ../idl --idl ../idl/checkout.proto
	@cd app/checkout && cwgo server --type RPC --service checkout --module github.com/Whitea029/whmall/app/checkout --pass "-use github.com/Whitea029/whmall/rpc_gen/kitex_gen" -I ../../idl --idl ../../idl/checkout.proto


PHONY: gen-order
gen-order:
	@cd rpc_gen && cwgo client --type RPC --service order --module github.com/Whitea029/whmall/rpc_gen --I ../idl --idl ../idl/order.proto
	@cd app/order && cwgo server --type RPC --service order --module github.com/Whitea029/whmall/app/order --pass "-use github.com/Whitea029/whmall/rpc_gen/kitex_gen" -I ../../idl --idl ../../idl/order.proto


PHONY: gen-email
gen-email:
	@cd rpc_gen && cwgo client --type RPC --service email --module github.com/Whitea029/whmall/rpc_gen --I ../idl --idl ../idl/email.proto
	@cd app/email && cwgo server --type RPC --service email --module github.com/Whitea029/whmall/app/email --pass "-use github.com/Whitea029/whmall/rpc_gen/kitex_gen" -I ../../idl --idl ../../idl/email.proto

