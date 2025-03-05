package mallagent

import (
	"context"
	"time"

	"github.com/cloudwego/eino/compose"
	"github.com/cloudwego/eino/flow/agent/react"
)

func NewReactAgent(ctx context.Context) (*react.Agent, error) {
    // 创建聊天模型
    chatModel, err := NewArkChatModel(ctx, nil)
    if err != nil {
        return nil, err
    }

    // 获取工具列表
    tools := GetMallAgentTools()

    // 创建Agent配置
    config := &react.AgentConfig{
        Model: chatModel,
        ToolsConfig: compose.ToolsNodeConfig{
            Tools:           tools,
            StreamableTools: []tool.StreamableTool{&OrderTool{}}, // 添加流式工具
        },
        MessageModifier: react.NewPersonaModifier(SystemPromptTemplate),
        MaxStep:         10,
        RetryConfig: &react.RetryConfig{
            MaxRetries: 3,
            RetryDelay: time.Second,
        },
    }

    return react.NewAgent(ctx, config)
}
