package agent

import (
	"context"
	"log"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"

	"github.com/strings77wzq/gomall/app/mall_agent/eino/mallagent"
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

	// 创建用户消息
	userMsg := &mallagent.UserMessage{
		UserID:  req.UserID,
		Query:   req.Query,
		Context: req.Context,
	}

	// 处理查询
	result, err := mallagent.ProcessQuery(ctx, userMsg)
	if err != nil {
		log.Printf("处理用户查询失败: %v", err)
		c.JSON(consts.StatusInternalServerError, AgentResponse{
			Success:      false,
			ErrorMessage: "处理查询失败: " + err.Error(),
		})
		return
	}

	// 构建响应
	response := AgentResponse{
		Message: result["message"].(string),
		Data:    result,
		Success: true,
	}

	c.JSON(consts.StatusOK, response)
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

	return nil
}
