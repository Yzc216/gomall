package middleware

import (
	"context"
	"github.com/Yzc216/gomall/app/frontend/utils"
	"github.com/cloudwego/hertz/pkg/app"
	"strconv"
)

func JwtOnlyParseMiddleware() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		//获取claims
		claimsPayload, err := JwtMiddleware.GetClaimsFromJWT(ctx, c)
		if err != nil {
			c.Next(ctx)
		}
		//获取payload中的userID信息
		user := claimsPayload[utils.IdentityKey].(map[string]interface{})
		UserIdStr := user[utils.UserIdKey].(string)
		UserId, err := strconv.ParseUint(UserIdStr, 10, 64)
		if err != nil {
			c.Next(ctx)
		}
		//存入上下文
		ctx = context.WithValue(ctx, utils.IdentityKey, UserId)
		c.Next(ctx)
	}
}
