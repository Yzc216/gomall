package service

import (
	"context"
	"github.com/Yzc216/gomall/app/user/biz/dal/mysql"
	user "github.com/Yzc216/gomall/app/user/kitex_gen/user"
	"github.com/joho/godotenv"

	"testing"
)

func TestRegister_Run(t *testing.T) {
	godotenv.Load("../../.env")
	mysql.Init()

	ctx := context.Background()
	s := NewRegisterService(ctx)
	// init req and assert value
	req := &user.RegisterReq{
		UserInfo: &user.User{
			Username: "2",
			Password: "123",
			Phone:    "1",
			Email:    "bb@qq.com",
			Role:     []uint32{2},
		},
		PasswordConfirm: "123",
	}

	resp, err := s.Run(req)

	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)
	t.Logf("role: %v", resp.Role)

	// todo: edit your unit test

}
