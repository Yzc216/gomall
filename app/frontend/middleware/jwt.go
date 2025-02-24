package middleware

import (
	"context"
	"fmt"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/jwt"
)

var (
	JwtMiddleware *jwt.HertzJWTMiddleware
	err           error
	identity      = "privilege"
)

func InitJWT() {
	JwtMiddleware, err = jwt.New(
		&jwt.HertzJWTMiddleware{
			Realm:       "govall",
			Key:         []byte("secret key"),
			Timeout:     time.Minute * 1,
			MaxRefresh:  time.Hour,
			IdentityKey: identity,
			Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
				// 用于验证登陆信息的函数
				c.Next(ctx) // 先走逻辑后再处理
				role, exits := c.Get("role")
				fmt.Println("auth: ", role, exits)
				if !exits {
					return "", jwt.ErrFailedAuthentication
				}

				return role, nil
			},
			PayloadFunc: func(data interface{}) jwt.MapClaims {
				return jwt.MapClaims{
					identity: data,
				}
			},
			LoginResponse: func(ctx context.Context, c *app.RequestContext, code int, token string, expire time.Time) {
				c.Redirect(consts.StatusFound, []byte("/sign-up"))
			},
			// Authorizator: func(data interface{}, ctx context.Context, c *app.RequestContext) bool {
			// 	role := data.([]interface{})[0]
			// 	allpath := c.HandlerName()
			// 	index := strings.LastIndex(allpath, "/")
			// 	if index != -1 {
			// 		allpath = allpath[index+1:]
			// 	}
			// 	fmt.Println(role, allpath)
			// 	// 拿到了role和调用的方法,开始鉴权
			// 	return true
			// },
			Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
				fmt.Println("Unauthorized: ", message)
				c.Redirect(consts.StatusFound, []byte("/sign-in"))
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
