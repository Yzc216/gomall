package middleware

import (
	"context"
	"fmt"
	"github.com/Yzc216/gomall/app/frontend/utils"

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
		"middleware/casbin/model.conf",
		"middleware/casbin/policy.csv",
		subjectFromJWT,
	)
	if casbinErr != nil {
		panic(err)
	}
}

func subjectFromJWT(ctx context.Context, c *app.RequestContext) string {
	claims := jwt.ExtractClaims(ctx, c)
	for k, v := range claims {
		if k == utils.IdentityKey {
			userInfo := v.(map[string]interface{})
			role := userInfo[utils.UserRole].([]interface{})
			a := "" + fmt.Sprint(role[0].(float64))
			return a
		}
	}
	return ""
}

func UnAuthorization(ctx context.Context, c *app.RequestContext) {
	c.AbortWithStatus(consts.StatusForbidden)
}
