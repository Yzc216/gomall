package agent

import (
	"context"
	"fmt"
	"log"

	"github.com/cloudwego/eino/schema"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// 请求结构
type AgentRequest struct {
	UserID  string            `json:"user_id"`
	Query   string            `json:"query"`
	Context map[string]string `json:"context,omitempty"`
}

// 响应结构
type AgentResponse struct {
	Message      string                 `json:"message"`
	Data         map[string]interface{} `json:"data,omitempty"`
	Success      bool                   `json:"success"`
	ErrorMessage string                 `json:"error_message,omitempty"`
}

// HandleQuery 处理查询请求
func HandleQuery(ctx context.Context, c *app.RequestContext) {
	var req AgentRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(consts.StatusBadRequest, AgentResponse{
			Success:      false,
			ErrorMessage: "请求格式错误: " + err.Error(),
		})
		return
	}

	// 验证必需字段
	if req.UserID == "" {
		c.JSON(consts.StatusBadRequest, AgentResponse{
			Success:      false,
			ErrorMessage: "用户ID不能为空",
		})
		return
	}

	if req.Query == "" {
		c.JSON(consts.StatusBadRequest, AgentResponse{
			Success:      false,
			ErrorMessage: "查询内容不能为空",
		})
		return
	}

	// 1. 获取相关知识
	docs, err := RetrieveKnowledge(ctx, req.Query)
	if err != nil {
		log.Printf("获取相关知识失败: %v", err)
		c.JSON(consts.StatusInternalServerError, AgentResponse{
			Success:      false,
			ErrorMessage: "获取相关知识失败: " + err.Error(),
		})
		return
	}

	// 2. 构建上下文
	messages := buildContextMessages(docs)
	messages = append(messages, schema.UserMessage(req.Query))

	// 3. 调用Agent处理
	agent, err := NewReactAgent(ctx)
	if err != nil {
		log.Printf("创建ReactAgent失败: %v", err)
		c.JSON(consts.StatusInternalServerError, AgentResponse{
			Success:      false,
			ErrorMessage: "创建ReactAgent失败: " + err.Error(),
		})
		return
	}

	// 4. 流式处理
	stream, err := agent.Stream(ctx, messages)
	if err != nil {
		log.Printf("流式处理失败: %v", err)
		c.JSON(consts.StatusInternalServerError, AgentResponse{
			Success:      false,
			ErrorMessage: "流式处理失败: " + err.Error(),
		})
		return
	}

	// 5. 返回SSE响应
	c.Stream(func(w *app.ResponseWriter) bool {
		msg, err := stream.Read()
		if err != nil {
			return false
		}
		w.Write([]byte(msg.Content))
		return true
	})
}

// HandleHistory 处理获取历史会话请求
func HandleHistory(ctx context.Context, c *app.RequestContext) {
	userID := c.Query("user_id")
	if userID == "" {
		c.JSON(consts.StatusBadRequest, AgentResponse{
			Success:      false,
			ErrorMessage: "用户ID不能为空",
		})
		return
	}

	// 从缓存中获取历史
	history, err := mallagent.GetSessionHistory(ctx, userID)
	if err != nil {
		c.JSON(consts.StatusInternalServerError, AgentResponse{
			Success:      false,
			ErrorMessage: "获取会话历史失败: " + err.Error(),
		})
		return
	}

	c.JSON(consts.StatusOK, map[string]interface{}{
		"history": history,
		"success": true,
	})
}

// HandleClear 处理清除会话请求
func HandleClear(ctx context.Context, c *app.RequestContext) {
	userID := c.Query("user_id")
	if userID == "" {
		c.JSON(consts.StatusBadRequest, AgentResponse{
			Success:      false,
			ErrorMessage: "用户ID不能为空",
		})
		return
	}

	// 清除会话历史
	if err := mallagent.ClearSession(ctx, userID); err != nil {
		c.JSON(consts.StatusInternalServerError, AgentResponse{
			Success:      false,
			ErrorMessage: "清除会话历史失败: " + err.Error(),
		})
		return
	}

	c.JSON(consts.StatusOK, map[string]interface{}{
		"message": "会话历史已清除",
		"success": true,
	})
}

// 添加SSE流式响应处理
func HandleStream(ctx context.Context, c *app.RequestContext) {
	userID := c.Query("user_id")

	c.SetStatusCode(consts.StatusOK)
	c.Response.Header.Set("Content-Type", "text/event-stream")
	c.Response.Header.Set("Cache-Control", "no-cache")
	c.Response.Header.Set("Connection", "keep-alive")

	// 创建消息通道
	msgChan := make(chan string)
	defer close(msgChan)

	go func() {
		// 调用Agent处理逻辑
		resp, err := ProcessRequest(ctx, &AgentRequest{
			UserID: userID,
			Query:  c.Query("query"),
		})
		if err != nil {
			msgChan <- fmt.Sprintf("错误: %v", err)
			return
		}

		for message := range resp.Stream {
			msgChan <- message.Content
		}
	}()

	for {
		select {
		case msg := <-msgChan:
			c.Write([]byte(fmt.Sprintf("data: %s\n\n", msg)))
			c.Flush()
		case <-ctx.Done():
			return
		}
	}
}

// BindRoutes 绑定路由
func BindRoutes(group *app.RouterGroup) error {
	// 初始化服务
	if err := mallagent.InitAgentServices(context.Background()); err != nil {
		return err
	}

	// 注册路由
	group.POST("/query", HandleQuery)
	group.GET("/history", HandleHistory)
	group.POST("/clear", HandleClear)
	group.GET("/stream", HandleStream)

	return nil
}
