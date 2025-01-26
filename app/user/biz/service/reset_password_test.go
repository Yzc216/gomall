package service

import (
	"context"
	"github.com/Yzc216/gomall/app/user/biz/dal/mysql"
	user "github.com/Yzc216/gomall/app/user/kitex_gen/user"
	"github.com/joho/godotenv"
	"testing"
)

func TestResetPassword_Run(t *testing.T) {
	godotenv.Load("../../.env")
	mysql.Init()
	ctx := context.Background()
	s := NewResetPasswordService(ctx)
	// init req and assert value

	req := &user.ResetPasswordReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
