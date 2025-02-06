// Code generated by hertz generator. DO NOT EDIT.

package user

import (
	user "github.com/Yzc216/gomall/app/frontend/biz/handler/user"
	"github.com/cloudwego/hertz/pkg/app/server"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	{
		_admin := root.Group("/admin", _adminMw()...)
		_admin.GET("/users", append(_admin0Mw(), user.Admin)...)
		_users := _admin.Group("/users", _usersMw()...)
		_users.POST("/ban", append(_banuserMw(), user.BanUser)...)
		_users.POST("/role", append(_setroleMw(), user.SetRole)...)
	}
	{
		_user := root.Group("/user", _userMw()...)
		_user.POST("/password", append(_resetpasswordMw(), user.ResetPassword)...)
		_user.GET("/profile", append(_user0Mw(), user.User)...)
		_user.POST("/profile", append(_updateuserMw(), user.UpdateUser)...)
	}
}
