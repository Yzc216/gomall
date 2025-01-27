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
	root.POST("/user", append(_updateuserMw(), user.UpdateUser)...)
	{
		_admin := root.Group("/admin", _adminMw()...)
		_admin.GET("/user", append(_admin0Mw(), user.Admin)...)
		_user := _admin.Group("/user", _userMw()...)
		_user.POST("/role", append(_setroleMw(), user.SetRole)...)
		_admin.POST("/user", append(_banuserMw(), user.BanUser)...)
	}
	root.GET("/user", append(_user1Mw(), user.User)...)
	_user0 := root.Group("/user", _user0Mw()...)
	_user0.POST("/password", append(_resetpasswordMw(), user.ResetPassword)...)
}
