package middleware

import (
	"context"
	jwtUtils "github.com/Yzc216/gomall/app/frontend/biz/utils"
	frontendUtils "github.com/Yzc216/gomall/app/frontend/utils"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/sessions"
	"strings"
)

func JWTAuth() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		// 从Cookie或Header获取Token
		tokenByte := c.Cookie("token")
		tokenString := string(tokenByte)
		if tokenString == "" {
			authHeaderByte := c.GetHeader("Authorization")
			authHeader := string(authHeaderByte)
			if authHeader != "" {
				tokenString = strings.Replace(authHeader, "Bearer ", "", 1)
			}
		}

		// 解析Token
		claims, err := jwtUtils.ParseJWT(tokenString)
		if err != nil || claims == nil {
			c.JSON(401, utils.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// 存储用户信息到上下文
		c.Set("user_id", claims.UserID)
		c.Set("role", claims.Roles)
		c.Next(ctx)
	}
}

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
			c.Redirect(consts.StatusFound, []byte("/sign-in?next="+c.FullPath()))
			c.Abort()
			return
		}
		c.Next(ctx)
	}
}
