export ROOT_MOD=github.com/Yzc216/gomall

.PHONY: gen-demo-proto
gen-demo-proto:
	@cd demo/hello &&  cwgo server --type RPC --module ${ROOT_MOD}/demo/hello --service hellotest --idl ../../idl/echo.proto -I ../../idl/

.PHONY: gen-frontend
gen-frontend:
	@cd app/frontend && cwgo server --type HTTP --idl ../../idl/frontend/user_page.proto --service frontend -I ../../idl --module ${ROOT_MOD}/app/frontend

.PHONY: gen-user
gen-user:
	@cd rpc_gen && cwgo client --type RPC  --service user --module  ${ROOT_MOD}/rpc_gen --I ../idl --idl ../idl/user.proto
	@cd app/user && cwgo server --type RPC  --service user --module  ${ROOT_MOD}/app/user  --pass "-use  ${ROOT_MOD}/rpc_gen/kitex_gen" -I ../../idl  --idl ../../idl/user.proto

.PHONY: gen-product
gen-product:
	@cd rpc_gen && cwgo client --type RPC  --service product --module  ${ROOT_MOD}/rpc_gen --I ../idl --idl ../idl/product.proto
	@cd app/product && cwgo server --type RPC  --service product --module  ${ROOT_MOD}/app/product  --pass "-use  ${ROOT_MOD}/rpc_gen/kitex_gen" -I ../../idl  --idl ../../idl/product.proto

.PHONY: gen-cart
gen-cart:
	@cd rpc_gen && cwgo client --type RPC  --service cart --module  ${ROOT_MOD}/rpc_gen --I ../idl --idl ../idl/cart.proto
	@cd app/cart && cwgo server --type RPC  --service cart --module  ${ROOT_MOD}/app/cart  --pass "-use  ${ROOT_MOD}/rpc_gen/kitex_gen" -I ../../idl  --idl ../../idl/cart.proto

.PHONY: gen-payment
gen-payment:
	@cd rpc_gen && cwgo client --type RPC  --service payment --module  ${ROOT_MOD}/rpc_gen --I ../idl --idl ../idl/payment.proto
	@cd app/payment && cwgo server --type RPC  --service payment --module  ${ROOT_MOD}/app/payment  --pass "-use  ${ROOT_MOD}/rpc_gen/kitex_gen" -I ../../idl  --idl ../../idl/payment.proto