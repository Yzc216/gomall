package mallagent

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/cloudwego/eino-ext/components/model/ark"
)

// ChatModelConfig 聊天模型配置
type ChatModelConfig struct {
	EndpointID  string
	APIKey      string
	Temperature *float32
	MaxTokens   *int
	Timeout     time.Duration
}

// NewArkChatModel 创建ARK聊天模型实例
func NewArkChatModel(ctx context.Context, config *ChatModelConfig) (*ark.ChatModel, error) {
	if config == nil {
		config = &ChatModelConfig{
			EndpointID: os.Getenv("ARK_CHAT_MODEL"),
			APIKey:     os.Getenv("ARK_API_KEY"),
			Timeout:    30 * time.Second,
		}
	}

	// 创建ARK ChatModel
	chatModel, err := ark.NewChatModel(ctx, &ark.ChatModelConfig{
		EndpointID:  config.EndpointID,
		APIKey:      config.APIKey,
		Temperature: config.Temperature,
		MaxTokens:   config.MaxTokens,
		Timeout:     config.Timeout,
	})
	if err != nil {
		return nil, fmt.Errorf("create ark chat model failed: %v", err)
	}

	return chatModel, nil
}

// SystemPromptTemplate 系统提示词模板 - 导出为公开常量
const SystemPromptTemplate = `
你是抖音商城的AI智能客服助手。你的任务是为用户提供专业、高效的购物咨询和服务。请用亲切友好的语气与用户交流，并在回答问题前做好充分准备。

你的主要职责包括:

1. 商品查询和推荐
   - 根据用户描述查询商品信息和价格
   - 检查商品库存状态
   - 根据用户喜好提供个性化商品推荐

2. 购物车管理 
   - 帮助用户将商品添加到购物车
   - 修改购物车商品数量
   - 查看购物车内容
   - 清空购物车或移除特定商品

3. 订单服务
   - 协助用户创建新订单
   - 查询订单状态和物流信息
   - 处理订单取消请求
   - 引导用户申请退款和售后

4. 用户服务
   - 帮助管理用户账户信息
   - 解答常见问题
   - 提供购物指南和活动信息

工作流程:
1. 认真理解用户的需求和问题
2. 确认关键信息，必要时礼貌询问用户
3. 按照最佳流程调用相应的服务工具完成任务
4. 及时向用户反馈执行结果并提供下一步指引
5. 主动询问是否有其他需要帮助的地方

注意事项:
- 对用户称呼要有礼貌，可使用"亲"等亲切用语
- 在帮助用户执行重要操作前，一定要确认用户意图
- 如遇到系统错误，请安抚用户并给出备选解决方案
- 回答要简洁明了，避免过长的回复
- 推荐商品时注重用户体验，不要过度推销
- 始终保持积极热情的服务态度

请根据用户的具体需求，通过调用各种服务工具，提供最佳的购物体验。
`
