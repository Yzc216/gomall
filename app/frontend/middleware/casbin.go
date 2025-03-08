package middleware

import (
	"context"
	"fmt"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/casbin"
	"github.com/hertz-contrib/jwt"
)

var (
	Casbinauth *casbin.Middleware
	casbinErr  error
)

func InitCasbin() {
	Casbinauth, casbinErr = casbin.NewCasbinMiddleware(
		"/home/wuu/vvmall/gomall/app/frontend/middleware/casbin/model.conf",
		"/home/wuu/vvmall/gomall/app/frontend/middleware/casbin/policy.csv",
		getInfo,
	)
	if casbinErr != nil {
		panic(err)
	}
}

func getInfo(ctx context.Context, c *app.RequestContext) string {
	claims := jwt.ExtractClaims(ctx, c)
	for k, v := range claims {
		if k == identity {
			vv := v.([]interface{})[0].(float64)
			v_str := "" + fmt.Sprint(vv)
			fmt.Println(v_str)
			return v_str
		}
	}
	return ""
}

func UnAuthorization(ctx context.Context, c *app.RequestContext) {
	c.AbortWithStatus(consts.StatusForbidden)
}
