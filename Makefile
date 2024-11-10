PHONY: gen-frontend
gen-frontend:
	@cd app/frontend && cwgo server --type HTTP --idl ../../idl/frontend/category_page.proto --service frontend --module github.com/Whitea029/whmall/app/frontend -I ../../idl

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

