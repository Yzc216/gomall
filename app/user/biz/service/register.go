package service

import (
	"context"
	"errors"
	"github.com/Yzc216/gomall/app/user/biz/dal/mysql"
	"github.com/Yzc216/gomall/app/user/biz/model"
	user "github.com/Yzc216/gomall/app/user/kitex_gen/user"
	"github.com/cloudwego/kitex/pkg/kerrors"
)

type RegisterService struct {
	ctx context.Context
} // NewRegisterService new RegisterService
func NewRegisterService(ctx context.Context) *RegisterService {
	return &RegisterService{ctx: ctx}
}

// Run create note info
func (s *RegisterService) Run(req *user.RegisterReq) (resp *user.RegisterResp, err error) {
	// Finish your business logic.

	if req.UserInfo.Username == "" {
		return nil, kerrors.NewBizStatusError(400, "username is required")
	}
	if req.UserInfo.Phone == "" || req.UserInfo.Email == "" {
		return nil, errors.New("email or phone is required")
	}
	if req.UserInfo.Password == "" || req.PasswordConfirm == "" {
		return nil, errors.New("password is empty")
	}
	if req.UserInfo.Password != req.PasswordConfirm {
		return nil, errors.New("password confirmation failed")
	}

	var Auths []model.Authority
	for _, val := range req.UserInfo.Role {
		switch val {
		case model.AdminType:
			Auths = append(Auths, model.Authority{
				AuthorityId:   model.AdminType,
				AuthorityName: "管理员",
			})
		case model.UserType:
			Auths = append(Auths, model.Authority{
				AuthorityId:   model.UserType,
				AuthorityName: "普通用户",
			})
		case model.MerchantType:
			Auths = append(Auths, model.Authority{
				AuthorityId:   model.MerchantType,
				AuthorityName: "商家",
			})
		default:
			Auths = append(Auths, model.Authority{
				AuthorityId:   model.UserType,
				AuthorityName: "普通用户",
			})
		}
	}

	newUser := &model.User{
		Username:  req.UserInfo.Username,
		Password:  req.UserInfo.Password,
		Phone:     req.UserInfo.Phone,
		Email:     req.UserInfo.Email,
		Avatar:    req.UserInfo.Avatar,
		Authority: Auths,
	}

	userInter, err := model.CreateUser(s.ctx, mysql.DB, newUser)
	if err != nil {
		return nil, err
	}

	var roles []uint32
	for _, val := range userInter.Authority {
		roles = append(roles, val.AuthorityId)
	}

	return &user.RegisterResp{
		UserId: userInter.ID,
		Role:   roles,
	}, nil
}
