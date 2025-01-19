.PHONY: gen-demo-proto
gen-demo-proto:
	@cd demo/hello &&  cwgo server --type RPC --module github.com/Yzc216/gomall/demo/hello --service hellotest --idl ../../idl/echo.proto -I ../../idl/

.PHONY: gen-frontend
gen-frontend:
	@cd app/frontend && cwgo server --type HTTP --idl ../../idl/frontend/common.proto --service frontend -I ../../idl -module github.com/Yzc216/gomall/app/frontend  && cwgo server --type HTTP --idl ../../idl/frontend/auth_page.proto --service frontend -I ../../idl -module github.com/Yzc216/gomall/app/frontend && cwgo server --type HTTP --idl ../../idl/frontend/home.proto --service frontend -I ../../idl -module github.com/Yzc216/gomall/app/frontend