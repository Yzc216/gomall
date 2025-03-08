module github.com/strings77wzq/gomall/app/mall_agent

go 1.23.6

require (
	github.com/cloudwego/eino v0.3.14
	github.com/cloudwego/eino-ext v0.3.14
	github.com/cloudwego/eino-ext/components/retriever/vikingdb v0.1.0
	github.com/cloudwego/eino-ext/devops v0.1.3
	github.com/cloudwego/hertz v0.9.6
	github.com/cloudwego/kitex v0.12.3
	github.com/go-redis/redis/v8 v8.11.5
	github.com/joho/godotenv v1.5.1
	github.com/strings77wzq/gomall v1.0.0
	github.com/cloudwego/eino-ext/retriever/redis v0.1.3
)

replace (
	github.com/chenzhuoyu/iasm => github.com/cloudwego/iasm v0.2.0
	github.com/strings77wzq/gomall/rpc_gen => github.com/Yzc216/gomall/rpc_gen v0.0.0-20250209153341-6cd64e7ead4d
	github.com/cloudwego/eino-ext/components/model/ark => github.com/cloudwego/eino-ext v0.3.14
)