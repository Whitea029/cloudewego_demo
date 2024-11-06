PHONY: gen-frontend
gen-frontend:
	@cd app/frontend && cwgo server --type HTTP --idl ../../idl/frontend/auth_page.proto --service frontend --module github.com/Whitea029/whmall/app/frontend -I ../../idl

PHONY: gen-user
gen-user:
	@cd app/user && cwgo server --type RPC --service user --module github.com/Whitea029/whmall/app/user --pass "-use github.com/Whitea029/whmall/rpc_gen/kitex_gen" -I ../../idl --idl ../../idl/user.proto
	@cd rpc_gen && cwgo client --type RPC --service user --module github.com/Whitea029/whmall/rpc_gen --I ../idl --idl ../idl/user.proto