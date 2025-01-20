module github.com/Yzc216/gomall/app/cart

go 1.22.11

replace (
	github.com/Yzc216/gomall/rpc_gen => ../../rpc_gen
	github.com/apache/thrift => github.com/apache/thrift v0.13.0
)

require github.com/golang/protobuf v1.5.4 // indirect
