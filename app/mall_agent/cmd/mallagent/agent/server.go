package agent

import (
	"context"
	"mime"
	"path/filepath"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/middlewares/server/cors"
	"github.com/cloudwego/hertz/pkg/app/server/render/sse"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/cloudwego/hertz/pkg/route"
)

// 中间件
type middleware struct{}

// CORS 跨域中间件
func (m *middleware) CORS() app.HandlerFunc {
	return cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
		ExposeHeaders:    []string{"Content-Length"},
		MaxAge:           86400,
	})
}

// AccessLog 访问日志中间件
func (m *middleware) AccessLog() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		c.Next(ctx)
	}
}

// 全局中间件实例
var middleware = &middleware{}

// ChatRequest 定义聊天请求结构
type ChatRequest struct {
	ID      string   `query:"id"`      // 会话ID
	Message string   `query:"msg"`     // 用户消息
	History []string `query:"history"` // 历史消息
}

// BindRoutes 绑定所有路由
func BindRoutes(r *route.RouterGroup) error {
	// 初始化Agent
	if err := Init(); err != nil {
		return err
	}

	// 添加CORS中间件
	r.Use(middleware.CORS())

	// 添加请求日志
	r.Use(middleware.AccessLog())

	// 健康检查
	r.GET("/health", func(ctx context.Context, c *app.RequestContext) {
		if err := GetClients().HealthCheck(ctx); err != nil {
			c.JSON(consts.StatusServiceUnavailable, map[string]interface{}{
				"status": "unhealthy",
				"error":  err.Error(),
			})
			return
		}
		c.JSON(consts.StatusOK, map[string]interface{}{
			"status": "healthy",
		})
	})

	// 聊天API - SSE流式响应
	r.GET("/api/chat", func(ctx context.Context, c *app.RequestContext) {
		req := &ChatRequest{
			ID:      c.Query("id"),
			Message: c.Query("msg"),
			History: c.QueryArgs().PeekMulti("history"),
		}

		// 设置SSE响应头
		c.Response.Header.Set("Content-Type", "text/event-stream")
		c.Response.Header.Set("Cache-Control", "no-cache")
		c.Response.Header.Set("Connection", "keep-alive")
		c.Response.Header.Set("Transfer-Encoding", "chunked")

		s := sse.NewStreamWriter(c.GetWriter())

		// 调用Agent处理请求
		sr, err := HandleChatStream(ctx, req)
		if err != nil {
			s.Publish(&sse.Event{
				Data: []byte("Error: " + err.Error()),
			})
			return
		}

		// 流式返回响应
		for {
			msg, err := sr.Recv()
			if err != nil {
				break
			}
			err = s.Publish(&sse.Event{
				Data: []byte(msg.Content),
			})
			if err != nil {
				break
			}
		}
	})

	// 历史记录API
	r.GET("/api/history", func(ctx context.Context, c *app.RequestContext) {
		id := c.Query("id")
		history, err := GetHistory(id)
		if err != nil {
			c.JSON(consts.StatusInternalServerError, map[string]interface{}{
				"error": err.Error(),
			})
			return
		}
		c.JSON(consts.StatusOK, history)
	})

	// 静态文件服务
	r.GET("/", func(ctx context.Context, c *app.RequestContext) {
		content, err := webContent.ReadFile("web/index.html")
		if err != nil {
			c.String(consts.StatusNotFound, "File not found")
			return
		}
		c.Header("Content-Type", "text/html")
		c.Write(content)
	})

	// 其他静态资源
	r.GET("/:file", func(ctx context.Context, c *app.RequestContext) {
		file := c.Param("file")
		content, err := webContent.ReadFile("web/" + file)
		if err != nil {
			c.String(consts.StatusNotFound, "File not found")
			return
		}
		contentType := mime.TypeByExtension(filepath.Ext(file))
		if contentType == "" {
			contentType = "application/octet-stream"
		}
		c.Header("Content-Type", contentType)
		c.Write(content)
	})

	return nil
}
