package service

import (
	"context"
	"errors"
	"github.com/Yzc216/gomall/app/user/biz/dal/mysql"
	"github.com/Yzc216/gomall/app/user/biz/model"
	user "github.com/Yzc216/gomall/rpc_gen/kitex_gen/user"
	"golang.org/x/crypto/bcrypt"
)

type LoginService struct {
	ctx context.Context
} // NewLoginService new LoginService
func NewLoginService(ctx context.Context) *LoginService {
	return &LoginService{ctx: ctx}
}

// Run create note info
func (s *LoginService) Run(req *user.LoginReq) (resp *user.LoginResp, err error) {

	if req.LoginInfo == "" {
		return nil, errors.New("username, email or phone is required")
	}
	if req.Password == "" {
		return nil, errors.New("password is required")
	}

	row, err := model.Login(s.ctx, mysql.DB, req.LoginInfo)
	if err != nil {
		return nil, err
	}

	//比对密码
	err = bcrypt.CompareHashAndPassword([]byte(row.Password), []byte(req.Password))
	if err != nil {
		return nil, err
	}

	var auths []uint32
	for _, v := range row.Authority {
		auths = append(auths, v.AuthorityId)
	}

	resp = &user.LoginResp{
		UserId: row.ID,
		Role:   auths,
	}

	return
}
