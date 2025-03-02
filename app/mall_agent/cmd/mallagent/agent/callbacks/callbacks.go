package callbacks

import (
	"context"
	"log"

	"github.com/cloudwego/eino/flow/agent/react"
	"github.com/cloudwego/eino/schema"
)

// DefaultHandler 默认回调处理器
var DefaultHandler = &react.Callbacks{
	OnThought: func(ctx context.Context, thought string) {
		log.Printf("[THOUGHT] %s", thought)
	},
	OnAction: func(ctx context.Context, action *react.Action) {
		log.Printf("[ACTION] Tool: %s, Input: %+v", action.Tool, action.Input)
	},
	OnActionError: func(ctx context.Context, action *react.Action, err error) {
		log.Printf("[ACTION ERROR] Tool: %s, Error: %v", action.Tool, err)
	},
	OnActionResult: func(ctx context.Context, action *react.Action, result interface{}) {
		log.Printf("[ACTION RESULT] Tool: %s, Result: %+v", action.Tool, result)
	},
	OnFinish: func(ctx context.Context, m *schema.Message) {
		log.Printf("[FINISH] Content length: %d", len(m.Content))
	},
}
