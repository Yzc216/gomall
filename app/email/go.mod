module github.com/Yzc216/gomall/app/email

go 1.22.11

replace (
	github.com/Yzc216/gomall/rpc_gen => ../../rpc_gen
	github.com/apache/thrift => github.com/apache/thrift v0.13.0
)

require (
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/klauspost/compress v1.17.9 // indirect
	github.com/nats-io/nats.go v1.39.0 // indirect
	github.com/nats-io/nkeys v0.4.9 // indirect
	github.com/nats-io/nuid v1.0.1 // indirect
	golang.org/x/crypto v0.31.0 // indirect
	golang.org/x/sys v0.28.0 // indirect
)
