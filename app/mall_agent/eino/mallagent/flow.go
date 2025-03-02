package mallagent

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/cloudwego/eino/components/flow/agent/react"
	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/compose"
	"github.com/cloudwego/eino/schema"
	"github.com/joho/godotenv"

	"github.com/strings77wzq/gomall/app/mall_agent/config"
	"github.com/strings77wzq/gomall/app/mall_agent/discovery"
	"github.com/strings77wzq/gomall/app/mall_agent/pkg/clients"
)

var (
	// 全局服务实例
	serviceWrappers *ServiceWrappers
	// 全局工具列表
	toolsList []tool.BaseTool
)

// 系统提示词
const systemPrompt = `你是一个电商平台的智能助手。你可以:
1. 查询商品信息和库存
2. 管理购物车
3. 处理订单
4. 回答用户关于商品和服务的问题

请根据用户需求调用相应工具完成任务。在执行操作前，请确保理解用户意图并在必要时确认重要信息。`

// InitAgentServices 初始化代理服务
func InitAgentServices(ctx context.Context) error {
	// 加载环境变量
	if err := godotenv.Load(); err != nil {
		log.Printf("警告: .env 文件未找到: %v", err)
	}

	// 加载配置
	cfg := config.MustLoad()

	// 初始化服务发现
	registry, err := discovery.NewEtcdRegistry(cfg.Etcd)
	if err != nil {
		return fmt.Errorf("初始化服务注册中心失败: %v", err)
	}

	// 初始化RPC客户端
	rpcClients, err := clients.NewRPCClients(ctx, registry)
	if err != nil {
		return fmt.Errorf("初始化RPC客户端失败: %v", err)
	}

	// 创建服务适配器
	serviceWrappers = NewServiceWrappers(rpcClients)

	// 初始化工具
	if err := initTools(ctx); err != nil {
		return fmt.Errorf("初始化工具失败: %v", err)
	}

	return nil
}

// initTools 初始化工具列表
func initTools(ctx context.Context) error {
	var tools []tool.BaseTool

	// 添加商品服务工具
	productTools := NewProductTools(serviceWrappers.ProductSvc)
	tools = append(tools, productTools...)

	// 添加购物车服务工具
	cartTools := NewCartTools(serviceWrappers.CartSvc)
	tools = append(tools, cartTools...)

	// 添加订单服务工具
	orderTools := NewOrderTools(serviceWrappers.OrderSvc)
	tools = append(tools, orderTools...)

	// 保存全局工具列表
	toolsList = tools
	return nil
}

// GetMallAgentTools 获取全局工具列表
func GetMallAgentTools() []tool.BaseTool {
	return toolsList
}

// ProcessQuery 处理用户查询
func ProcessQuery(ctx context.Context, userMsg *UserMessage) (map[string]interface{}, error) {
	// 创建聊天模型
	chatModel, err := NewArkChatModel(ctx)
	if err != nil {
		return nil, fmt.Errorf("创建聊天模型失败: %v", err)
	}

	// 从知识库召回相关上下文
	retrievedContext, err := RetrieveKnowledge(ctx, userMsg.Query)
	if err != nil {
		log.Printf("从知识库召回上下文失败: %v", err)
		// 继续处理，即使知识库查询失败
	}

	// 创建 ReAct 代理配置
	reactConfig := &react.Config{
		ChatModel:  chatModel,
		Tools:      GetMallAgentTools(),
		MaxSteps:   25,
		AgentCache: NewMemoryCache(),
	}

	// 创建 ReAct 代理
	agent, err := react.NewAgent(ctx, reactConfig)
	if err != nil {
		return nil, fmt.Errorf("创建ReAct代理失败: %v", err)
	}

	// 准备系统提示词
	systemPrompt := LoadSystemPrompt()

	// 准备用户消息
	userMessages := []*schema.Message{
		{
			Role:    schema.MessageRoleSystem,
			Content: systemPrompt,
		},
	}

	// 添加知识库上下文（如果有）
	if retrievedContext != "" {
		contextMsg := &schema.Message{
			Role:    schema.MessageRoleSystem,
			Content: fmt.Sprintf("根据以下与该查询相关的知识信息处理用户请求：\n%s", retrievedContext),
		}
		userMessages = append(userMessages, contextMsg)
	}

	// 添加用户查询
	userMessages = append(userMessages, &schema.Message{
		Role:    schema.MessageRoleUser,
		Content: userMsg.Query,
	})

	// 添加上下文信息（如果有）
	if len(userMsg.Context) > 0 {
		contextJSON, _ := json.Marshal(userMsg.Context)
		contextMsg := &schema.Message{
			Role:    schema.MessageRoleSystem,
			Content: fmt.Sprintf("用户的上下文信息: %s", string(contextJSON)),
		}
		userMessages = append(userMessages, contextMsg)
	}

	// 运行代理
	response, err := agent.Run(ctx, userMessages)
	if err != nil {
		return nil, fmt.Errorf("执行代理失败: %v", err)
	}

	// 返回结果
	return map[string]interface{}{
		"message": response.Content,
		"success": true,
	}, nil
}

// LoadSystemPrompt 加载系统提示词
func LoadSystemPrompt() string {
	return `你是一个专业的电商平台客服助手。你可以：
1. 搜索和查询商品信息
2. 管理购物车
3. 创建和查询订单
4. 提供购物建议

请根据用户的需求，调用相应的工具来完成任务。如果用户想要添加商品到购物车，你应该先搜索商品，然后获取详情，检查库存，最后添加到购物车。
如果用户要下单，你应该先确认用户的购物车内容，然后创建订单。

请始终以礼貌、专业的方式与用户交流。`
}

// NewReactAgent 创建 ReAct Agent 实例
func NewReactAgent(ctx context.Context, config *react.AgentConfig) (*compose.Lambda, error) {
	if config == nil {
		defaultConfig, err := defaultReactAgentConfig(ctx)
		if err != nil {
			return nil, err
		}
		config = defaultConfig
	}

	// 添加重试机制
	config.RetryConfig = &react.RetryConfig{
		MaxRetries: 3,
		RetryDelay: time.Second,
	}

	// 添加错误处理
	config.ErrorHandler = func(err error) *schema.Message {
		return schema.SystemMessage(fmt.Sprintf("操作出现错误: %v, 请重试或联系客服", err))
	}

	// 创建 Agent
	agent, err := react.NewAgent(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("create agent failed: %v", err)
	}

	return compose.AnyLambda(agent.Generate, agent.Stream, nil, nil)
}

// defaultReactAgentConfig 默认配置
func defaultReactAgentConfig(ctx context.Context) (*react.AgentConfig, error) {
	// 初始化聊天模型
	chatModel, err := NewArkChatModel(ctx, nil)
	if err != nil {
		return nil, err
	}

	// 获取工具列表
	tools, err := GetTools(ctx)
	if err != nil {
		return nil, err
	}

	return &react.AgentConfig{
		Model: chatModel,
		ToolsConfig: compose.ToolsNodeConfig{
			Tools: tools,
		},
		MessageModifier: react.NewPersonaModifier(systemPrompt),
		MaxStep:         10,
	}, nil
}
