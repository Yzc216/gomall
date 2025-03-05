package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/cloudwego/eino-ext/devops"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/joho/godotenv"

	"github.com/strings77wzq/gomall/app/mall_agent/cmd/mallagent/agent"
	"github.com/strings77wzq/gomall/app/mall_agent/config"
	"github.com/strings77wzq/gomall/app/mall_agent/discovery"
	"github.com/strings77wzq/gomall/app/mall_agent/eino/mallagent"
	"github.com/strings77wzq/gomall/app/mall_agent/utils/env"
)

func init() {
	// 加载 .env 文件
	if err := godotenv.Load(); err != nil {
		log.Printf("警告: .env 文件未找到")
	}

	// 开发模式下初始化 devops
	if os.Getenv("MALL_AGENT_DEBUG") != "false" {
		err := devops.Init(context.Background())
		if err != nil {
			log.Printf("[mall agent dev] 初始化失败, err=%v", err)
		}
	}

	// 检查必要的环境变量
	env.MustHasEnvs("ARK_CHAT_MODEL", "ARK_EMBEDDING_MODEL", "ARK_API_KEY")
}

func main() {
	flag.Parse()

	// 初始化配置
	cfg := config.MustLoad()

	// 初始化服务发现
	registry, err := discovery.NewEtcdRegistry(cfg.Etcd)
	if err != nil {
		log.Fatalf("初始化服务注册中心失败: %v", err)
	}

	// 初始化知识库
	ctx := context.Background()
	if err := mallagent.LoadAndIndexKnowledge(ctx); err != nil {
		log.Printf("警告: 知识库索引失败: %v", err)
	}

	// 初始化Agent
	if err := agent.Init(); err != nil {
		log.Fatalf("初始化Agent失败: %v", err)
	}

	// 启动RPC服务（在后台运行）
	go func() {
		if err := startRPCServer(cfg); err != nil {
			log.Fatalf("RPC服务启动失败: %v", err)
		}
	}()

	// 创建 Hertz 服务器
	h := server.Default(
		server.WithHostPorts(cfg.Server.Address),
	)

	// 注册路由
	if err := agent.BindRoutes(h.Group("/mall")); err != nil {
		log.Fatalf("注册路由失败: %v", err)
	}

	// 优雅关闭
	go func() {
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
		<-sigCh
		h.Close()
	}()

	// 启动服务器
	log.Printf("mall_agent HTTP服务启动，监听地址 %s", cfg.Server.Address)
	h.Spin()
}

// startRPCServer 启动RPC服务器
// 修改startRPCServer函数

func startRPCServer(cfg *config.Config) error {
	// 初始化服务发现
	registry, err := discovery.NewEtcdRegistry(cfg.Etcd)
	if err != nil {
		return fmt.Errorf("初始化服务注册中心失败: %v", err)
	}

	// 初始化RPC客户端
	ctx := context.Background()
	if err := mallagent.InitClients(ctx, registry); err != nil {
		return fmt.Errorf("初始化RPC客户端失败: %v", err)
	}

	// 启动RPC服务器
	return rpc.StartRPCServer(cfg, registry)
}

	// 检查服务健康状态
	clients := mallagent.GetClients()
	if err := clients.HealthCheck(ctx); err != nil {
		log.Printf("警告: 服务健康检查失败: %v", err)
		// 不返回错误，允许即使某些服务不可用也能启动
	}

	log.Printf("RPC客户端初始化成功，已连接到所有微服务")
	return nil
}
