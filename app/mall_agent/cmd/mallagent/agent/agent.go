package agent

import (
	"context"
	"fmt"
	"os"
	"sync"

	"github.com/cloudwego/eino/compose"
	"github.com/cloudwego/eino/flow/agent/react"
	"github.com/cloudwego/eino/schema"
	"github.com/cloudwego/kitex/pkg/discovery"

	"github.com/strings77wzq/gomall/app/mall_agent/cmd/mallagent/agent/callbacks"
	"github.com/strings77wzq/gomall/app/mall_agent/eino/mallagent"
)

type Agent struct {
	agent *react.Agent
	once  sync.Once
}

// 全局Agent
var agent *Agent
var memory *mallagent.Memory

// Init 初始化Agent和相关资源
func Init() error {
	// 初始化内存管理
	if err := mallagent.InitMemory(os.Getenv("REDIS_ADDR")); err != nil {
		return fmt.Errorf("init memory failed: %v", err)
	}
	memory = mallagent.GetMemory()

	// 初始化Agent
	var err error
	agent, err = NewAgent(context.Background(), nil)
	return err
}

func NewAgent(ctx context.Context, registry discovery.Registry) (*Agent, error) {
	a := &Agent{}
	var initErr error

	a.once.Do(func() {
		// 初始化服务客户端
		if err := mallagent.InitClients(ctx, registry); err != nil {
			initErr = fmt.Errorf("init clients failed: %v", err)
			return
		}

		// 初始化ChatModel
		chatModel, err := mallagent.NewArkChatModel(ctx, nil)
		if err != nil {
			initErr = fmt.Errorf("create chat model failed: %v", err)
			return
		}

		// 获取所有工具
		tools, err := mallagent.GetTools(ctx)
		if err != nil {
			initErr = fmt.Errorf("get tools failed: %v", err)
			return
		}

		// 创建Agent
		a.agent, err = react.NewAgent(ctx, &react.AgentConfig{
			Model: chatModel,
			ToolsConfig: compose.ToolsNodeConfig{
				Tools: tools,
			},
			MessageModifier: react.NewPersonaModifier(mallagent.SystemPromptTemplate),
			MaxStep:         10,
		})
		if err != nil {
			initErr = fmt.Errorf("create agent failed: %v", err)
			return
		}
	})

	return a, initErr
}

// HandleChatStream 处理聊天请求并返回流式响应
func HandleChatStream(ctx context.Context, req *ChatRequest) (*schema.StreamReader[*schema.Message], error) {
	if agent == nil {
		return nil, fmt.Errorf("agent not initialized")
	}

	// 获取历史消息
	history := memory.Get(req.ID)
	messages := make([]*schema.Message, 0)

	// 添加历史消息
	for _, h := range history {
		messages = append(messages, h...)
	}

	// 添加新消息
	messages = append(messages, schema.UserMessage(req.Message))

	// 调用Agent处理
	return agent.Stream(ctx, messages, react.WithCallbacks(callbacks.DefaultHandler))
}

// GetHistory 获取历史记录
func GetHistory(id string) ([]*schema.Message, error) {
	history := memory.Get(id)
	if len(history) == 0 {
		return nil, nil
	}
	return history[len(history)-1], nil
}
