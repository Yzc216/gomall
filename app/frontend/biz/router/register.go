// Code generated by hertz generator. DO NOT EDIT.

package router

import (
	auth "github.com/Yzc216/gomall/app/frontend/biz/router/auth"
	home "github.com/Yzc216/gomall/app/frontend/biz/router/home"
	"github.com/cloudwego/hertz/pkg/app/server"
)

// GeneratedRegister registers routers generated by IDL.
func GeneratedRegister(r *server.Hertz) {
	//INSERT_POINT: DO NOT DELETE THIS LINE!
	auth.Register(r)

	home.Register(r)
}
