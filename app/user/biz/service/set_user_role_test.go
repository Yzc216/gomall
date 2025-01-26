package service

import (
	"context"
	"github.com/Yzc216/gomall/app/user/biz/dal/mysql"
	user "github.com/Yzc216/gomall/app/user/kitex_gen/user"
	"github.com/joho/godotenv"
	"testing"
)

func TestSetUserRole_Run(t *testing.T) {
	godotenv.Load("../../.env")
	mysql.Init()

	ctx := context.Background()
	s := NewSetUserRoleService(ctx)
	// init req and assert value

	req := &user.SetUserRoleReq{
		UserId:  550607018575069182,
		SetType: 1,
		NewRole: []uint32{2, 3},
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
