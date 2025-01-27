package service

import (
	"context"
	"github.com/Yzc216/gomall/app/user/biz/dal/mysql"
	user "github.com/Yzc216/gomall/app/user/kitex_gen/user"
	"github.com/joho/godotenv"
	"testing"
)

func TestUpdateUserInfo_Run(t *testing.T) {
	godotenv.Load("../../.env")
	mysql.Init()

	ctx := context.Background()
	s := NewUpdateUserInfoService(ctx)
	// init req and assert value

	req := &user.UpdateUserInfoReq{
		UserId: 551017802534813694,
		UserInfo: &user.User{
			Username: "yzc",
			Avatar:   "https://qmplusimg.henrongyi.top/gva_header.jpg",
			Email:    "yangzc216@163.com",
			Phone:    "19121763510",
		},
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
