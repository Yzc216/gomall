package middleware

import (
	"context"
	"fmt"
	"github.com/Yzc216/gomall/app/frontend/types"
	"github.com/Yzc216/gomall/app/frontend/utils"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"strconv"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/jwt"
)

var (
	JwtMiddleware *jwt.HertzJWTMiddleware
	err           error
)

func InitJWT() {
	JwtMiddleware, err = jwt.New(
		&jwt.HertzJWTMiddleware{
			Realm:       "Gomall",
			Key:         []byte("secret key"),
			Timeout:     time.Minute * 1,
			MaxRefresh:  time.Hour,
			IdentityKey: utils.IdentityKey,
			Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
				// 用于验证登陆信息的函数
				c.Next(ctx) // 先走登录逻辑拿到id和role

				userId, exists := c.Get("user_id")
				if !exists {
					return "", jwt.ErrFailedAuthentication
				}

				role, exists := c.Get("role")
				if !exists {
					return "", jwt.ErrFailedAuthentication
				}

				return &types.User{
					UserId: strconv.FormatUint(userId.(uint64), 10),
					Role:   role.([]uint32),
				}, nil
			},
			PayloadFunc: func(data interface{}) jwt.MapClaims {
				if u, ok := data.(*types.User); ok {
					return jwt.MapClaims{
						utils.IdentityKey: u,
					}
				}
				return jwt.MapClaims{}
			},
			LoginResponse: func(ctx context.Context, c *app.RequestContext, code int, token string, expire time.Time) {
				hlog.Info("LoginToken: ", token)
				hlog.Info("ExpireTime: ", expire.Format(time.RFC3339))
			},
			IdentityHandler: func(ctx context.Context, c *app.RequestContext) interface{} {
				claims := jwt.ExtractClaims(ctx, c)
				user := claims[utils.IdentityKey].(map[string]interface{})
				return &types.User{
					UserId: user[utils.UserIdKey].(string),
				}
			},
			Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
				fmt.Println("Unauthorized: ", message)
				byteRef := c.GetHeader("Referer")
				ref := string(byteRef)
				next := "/sign-in"
				if ref != "" {
					if utils.ValidateNext(ref) {
						next = fmt.Sprintf("%s?next=%s", next, ref)
					}
				}
				c.Redirect(consts.StatusFound, []byte(next))
			},
			SendCookie:     true,
			TokenLookup:    "cookie: jwt-cookie",
			SecureCookie:   false,
			CookieHTTPOnly: false,
			CookieName:     "jwt-cookie",
		},
	)
	if err != nil {
		panic(err)
	}
}

func JwtOnlyParseMiddleware() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		//获取claims
		claimsPayload, err := JwtMiddleware.GetClaimsFromJWT(ctx, c)
		if err != nil || claimsPayload == nil {
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
