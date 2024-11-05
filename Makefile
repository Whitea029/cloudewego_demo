PHONY: gen-frontend
gen-frontend:
	@cd app/frontend && cwgo server --type HTTP --idl ../../idl/frontend/auth_page.proto --service frontend --module github.com/Whitea029/whmall/app/frontend -I ../../idl