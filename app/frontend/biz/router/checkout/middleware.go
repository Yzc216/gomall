// Code generated by hertz generator.

package checkout

import (
	"github.com/Yzc216/gomall/app/frontend/middleware"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/casbin"
)

func rootMw() []app.HandlerFunc {
	// your code...
	return []app.HandlerFunc{
		middleware.JwtMiddleware.MiddlewareFunc(),
		middleware.JwtOnlyParseMiddleware(),
		// 1:管理员 2:普通用户 3:商家
		middleware.Casbinauth.RequiresRoles("1 2 3", casbin.WithLogic(casbin.OR), casbin.WithUnauthorized(middleware.UnAuthorization), casbin.WithForbidden(middleware.UnAuthorization)),
	}
}

func _checkoutMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _checkout0Mw() []app.HandlerFunc {
	// your code...
	return nil
}

func _checkoutresultMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _checkoutwaitingMw() []app.HandlerFunc {
	// your code...
	return nil
}
