FROM golang:1.20-alpine AS builder

WORKDIR /app

# 安装依赖
RUN apk add --no-cache gcc musl-dev

# 复制 Go 模块文件
COPY go.mod go.sum ./
RUN go mod download

# 复制源代码
COPY . .

# 构建
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o mall_agent ./cmd/mallagent

# 最终镜像
FROM alpine:latest

WORKDIR /app

# 安装运行时依赖
RUN apk --no-cache add ca-certificates tzdata
RUN cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo "Asia/Shanghai" > /etc/timezone
    
# 复制二进制文件
COPY --from=builder /app/mall_agent .

# 复制Web内容和配置
COPY --from=builder /app/cmd/mallagent/web ./web
COPY --from=builder /app/eino ./eino

# 暴露端口
EXPOSE 8080

# 设置环境变量
ENV MALL_AGENT_PORT=8080

# 运行应用
CMD ["./mall_agent"] 