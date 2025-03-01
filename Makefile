export ROOT_MOD=github.com/Yzc216/gomall

.PHONY: gen-demo-proto
gen-demo-proto:
	@cd demo/hello &&  cwgo server --type RPC --module ${ROOT_MOD}/demo/hello --service hellotest --idl ../../idl/echo.proto -I ../../idl/

.PHONY: gen-frontend
gen-frontend:
	@cd app/frontend && cwgo server --type HTTP --idl ../../idl/frontend/order_page.proto --service frontend -I ../../idl --module ${ROOT_MOD}/app/frontend

.PHONY: gen-frontend_category
gen-frontend_category:
	@cd app/frontend && cwgo server --type HTTP --idl ../../idl/frontend/category_page.proto --service frontend -I ../../idl --module ${ROOT_MOD}/app/frontend

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

.PHONY: gen-checkout
gen-checkout:
	@cd rpc_gen && cwgo client --type RPC  --service checkout --module  ${ROOT_MOD}/rpc_gen --I ../idl --idl ../idl/checkout.proto
	@cd app/checkout && cwgo server --type RPC  --service checkout --module  ${ROOT_MOD}/app/checkout  --pass "-use  ${ROOT_MOD}/rpc_gen/kitex_gen" -I ../../idl  --idl ../../idl/checkout.proto

.PHONY: gen-order
gen-order:
	@cd rpc_gen && cwgo client --type RPC  --service order --module  ${ROOT_MOD}/rpc_gen --I ../idl --idl ../idl/order.proto
	@cd app/order && cwgo server --type RPC  --service order --module  ${ROOT_MOD}/app/order  --pass "-use  ${ROOT_MOD}/rpc_gen/kitex_gen" -I ../../idl  --idl ../../idl/order.proto

.PHONY: gen-email
gen-email:
	@cd rpc_gen && cwgo client --type RPC  --service email --module  ${ROOT_MOD}/rpc_gen --I ../idl --idl ../idl/email.proto
	@cd app/email && cwgo server --type RPC  --service email --module  ${ROOT_MOD}/app/email  --pass "-use  ${ROOT_MOD}/rpc_gen/kitex_gen" -I ../../idl  --idl ../../idl/email.proto

.PHONY: gen-inventory
gen-inventory:
	@cd rpc_gen && cwgo client --type RPC  --service inventory --module  ${ROOT_MOD}/rpc_gen --I ../idl --idl ../idl/inventory.proto
	@cd app/inventory && cwgo server --type RPC  --service inventory --module  ${ROOT_MOD}/app/inventory  --pass "-use  ${ROOT_MOD}/rpc_gen/kitex_gen" -I ../../idl  --idl ../../idl/inventory.proto