package middleware

import (
	"context"
	"fmt"
	frontendUtils "github.com/Yzc216/gomall/app/frontend/utils"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/sessions"
)

func GlobalAuth() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		s := sessions.Default(c)
		//把session内容存入ctx，后续业务只需从ctx取
		ctx = context.WithValue(ctx, frontendUtils.SessionUserId, s.Get("user_id"))
		c.Next(ctx)
	}
}

func Auth() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		s := sessions.Default(c)
		userId := s.Get("user_id")
		if userId == nil {
			fmt.Println(c.FullPath())
			c.Redirect(consts.StatusFound, []byte("/sign-in?next="+c.FullPath()))
			c.Abort()
			return
		}
		c.Next(ctx)
	}
}
