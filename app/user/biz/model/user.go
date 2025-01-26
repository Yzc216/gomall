package model

import (
	"context"
	"errors"
	"fmt"
	"github.com/Yzc216/gomall/app/user/util"
	"gorm.io/gorm"
	"time"
)

type User struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	ID        uint64         `gorm:"primarykey;uniqueIndex;type:bigint;comment:用户ID"`
	Username  string         `gorm:"index;type:varchar(50);comment:用户登录名"`
	Password  string         `gorm:"type:char(60);comment:用户登录密码"`
	Avatar    string         `gorm:"default:https://qmplusimg.henrongyi.top/gva_header.jpg;comment:用户头像"`
	Phone     string         `gorm:"index;type:varchar(255);comment:用户手机号"`
	Email     string         `gorm:"uniqueIndex;type:varchar(255);comment:用户邮箱"`
	Enable    int            `gorm:"default:1;comment:用户是否被冻结 1正常 2冻结"`

	Authority []Authority `gorm:"many2many:user_authority;comment:用户角色"`
}

func (u *User) TableName() string {
	return "user"
}

func (u *User) GetUint32Auth() (auths []uint32) {
	for _, v := range u.Authority {
		auths = append(auths, v.AuthorityId)
	}
	return
}

// CreateUser
//
//	@Author: YZC 2025-01-23 20:56:04
//	@Description: 创建用户
func CreateUser(ctx context.Context, db *gorm.DB, u *User) (userInter *User, err error) {
	var user User
	if !errors.Is(db.WithContext(ctx).
		Where("username = ?", u.Username).
		First(&user).Error,
		gorm.ErrRecordNotFound) { // 判断用户名是否注册
		return &user, errors.New("用户名已注册")
	}

	u.ID = util.GenID()
	u.Password, err = util.BcryptHash(u.Password)
	if err != nil {
		return nil, errors.New("密码哈希加密错误")
	}

	return u, db.WithContext(ctx).Create(u).Error
}

// UpdateUser
//
//	@Author: YZC 2025-01-23 20:56:20
//	@Description: 更新用户信息
func UpdateUser(ctx context.Context, db *gorm.DB, u *User) error {
	return db.WithContext(ctx).
		Model(&User{}).
		Select("updated_at", "avatar", "phone", "email", "username").
		Where("id=?", u.ID).
		Updates(map[string]interface{}{
			"updated_at": time.Now(),
			"avatar":     u.Avatar,
			"phone":      u.Phone,
			"email":      u.Email,
			"username":   u.Username,
		}).Error

}

// UpdatePassword
//
//	@Author: YZC 2025-01-23 21:10:14
//	@Description:  更新密码
func UpdatePassword(ctx context.Context, db *gorm.DB, u *User, newPassword string) (err error) {
	var user User
	if err = db.WithContext(ctx).Where("id = ?", u.ID).First(&user).Error; err != nil {
		return err
	}
	if ok := util.BcryptCheck(u.Password, user.Password); !ok {
		return errors.New("原密码错误")
	}
	user.Password, err = util.BcryptHash(newPassword)

	return db.WithContext(ctx).Save(&user).Error
}

// UpdateAuthority
//
//	@Author: YZC 2025-01-23 21:10:21
//	@Description:  更新权限
func UpdateAuthority(ctx context.Context, db *gorm.DB, id uint64, auth []uint32) (err error) {
	// 1. 查询用户（确保用户存在）
	var user User
	if err = db.WithContext(ctx).Where("id = ?", id).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("用户不存在（id=%d）", id)
		}
		return fmt.Errorf("查询用户失败: %w", err)
	}
	fmt.Println(user.Authority)

	var NewAuths []Authority
	for _, v := range auth {
		NewAuths = append(NewAuths, Authority{
			AuthorityId:   v,
			AuthorityName: AuthorityTypeMap[v],
		})
	}

	user.Authority = NewAuths
	fmt.Println(user.Authority)

	return db.Model(&user).Association("Authority").Replace(NewAuths)
}

// BanUser
//
//	@Author: YZC 2025-01-23 22:23:10
//	@Description: 封禁用户
func BanUser(ctx context.Context, db *gorm.DB, id uint64) (err error) {
	var user User
	err = db.WithContext(ctx).Where("id = ?", id).First(&user).Error
	user.Enable = 2
	return db.WithContext(ctx).Save(&user).Error
}

// AddAuthority
//
//	@Author: YZC 2025-01-23 21:44:08
//	@Description: 增加权限
func AddAuthority(ctx context.Context, db *gorm.DB, id uint64, auth []uint32) (err error) {
	// 1. 查询用户（确保用户存在）
	var user User
	if err = db.WithContext(ctx).Where("id = ?", id).Preload("Authority").First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("用户不存在（id=%d）", id)
		}
		return fmt.Errorf("查询用户失败: %w", err)
	}

	//2.创建需要添加的权限
	var NewAuths []Authority
	for _, v := range auth {
		//判断是否存在,存在则跳过
		if authContains(user.Authority, v) {
			continue
		}
		NewAuths = append(NewAuths, Authority{
			AuthorityId:   v,
			AuthorityName: AuthorityTypeMap[v],
		})
	}

	//3.添加
	user.Authority = append(user.Authority, NewAuths...)
	return db.Model(&user).Association("Authority").Append(NewAuths)
	//db.WithContext(ctx).Save(&user).Error
}
func authContains(auth []Authority, id uint32) bool {
	for _, v := range auth {
		if v.AuthorityId == id {
			return true
		}
	}
	return false
}

// GetById
//
//	@Author: YZC 2025-01-23 21:10:25
//	@Description:   根据id获取用户信息
func GetById(ctx context.Context, db *gorm.DB, id uint64) (user *User, err error) {
	err = db.WithContext(ctx).Where("id = ?", id).Preload("Authority").First(&user).Error
	if user.Enable == 2 {
		return nil, errors.New("用户已被封禁")
	}
	return user, err
}

func GetBatchById(ctx context.Context, db *gorm.DB, page, pageSize int, ids []uint64) (users []*User, err error) {
	if len(ids) == 0 {
		return nil, nil
	}

	if page != 0 || pageSize != 0 {
		offset := pageSize * (page - 1)
		err = db.WithContext(ctx).Limit(pageSize).Offset(offset).Where(ids).Preload("Authority").Find(&users).Error
	}
	err = db.WithContext(ctx).Where(ids).Preload("Authority").Find(&users).Error
	return users, err
}

// GetByUsername
//
//	@Author: YZC 2025-01-23 21:10:28
//	@Description:   根据用户名获取用户信息
func GetByUsername(ctx context.Context, db *gorm.DB, username string) (user *User, err error) {
	err = db.WithContext(ctx).Where("username = ?", username).First(&user).Error
	if user.Enable == 2 {
		return nil, errors.New("用户已被封禁")
	}
	return user, err
}

// GetByEmail
//
//	@Author: YZC 2025-01-23 21:10:31
//	@Description:   根据邮箱获取用户信息
func GetByEmail(ctx context.Context, db *gorm.DB, email string) (user *User, err error) {
	err = db.WithContext(ctx).Where("email = ?", email).First(&user).Error
	if user.Enable == 2 {
		return nil, errors.New("用户已被封禁")
	}
	return user, err
}

// GetByPhone
//
//	@Author: YZC 2025-01-23 20:57:36
//	@Description: 根据电话获取用户信息
func GetByPhone(ctx context.Context, db *gorm.DB, phone string) (user *User, err error) {
	err = db.WithContext(ctx).Where("phone = ?", phone).First(&user).Error
	if user.Enable == 2 {
		return nil, errors.New("用户已被封禁")
	}
	return user, err
}

// DeleteUser
//
//	@Author: YZC 2025-01-23 21:10:35
//	@Description:  删除用户
func DeleteUser(ctx context.Context, db *gorm.DB, id uint64) (err error) {
	//return db.WithContext(ctx).Where("id = ?", id).Delete(&User{}).Error
	return db.WithContext(ctx).Select("authority").Where("id = ?", id).Delete(&User{}).Error

}
