package service

import (
	"context"
	user "github.com/Yzc216/gomall/rpc_gen/kitex_gen/user"
	"testing"
)

func TestBanUser_Run(t *testing.T) {
	ctx := context.Background()
	s := NewBanUserService(ctx)
	// init req and assert value

	req := &user.BanUserReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
