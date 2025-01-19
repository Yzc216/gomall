package service

import (
	"context"
	"errors"
	"github.com/Yzc216/gomall/app/user/biz/dal/mysql"
	"github.com/Yzc216/gomall/app/user/biz/model"
	user "github.com/Yzc216/gomall/app/user/kitex_gen/user"
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
	// Finish your business logic.
	if req.Email == "" || req.Password == "" {
		return nil, errors.New("email or password is empty")
	}

	row, err := model.GetByEmail(context.Background(), mysql.DB, req.Email)
	if err != nil {
		return nil, err
	}

	//比对密码
	err = bcrypt.CompareHashAndPassword([]byte(row.PasswordHashed), []byte(req.Password))
	if err != nil {
		return nil, err
	}

	resp = &user.LoginResp{
		UserId: int32(row.ID),
	}

	return
}
