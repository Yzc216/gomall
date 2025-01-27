package model

import (
	"gorm.io/gorm"
	"time"
)

const (
	AdminType uint32 = iota + 1
	UserType
	MerchantType
)

var AuthorityTypeMap = map[uint32]string{
	AdminType:    "管理员",
	UserType:     "普通用户",
	MerchantType: "商家",
}

type Authority struct {
	AuthorityId   uint32 `gorm:"primarykey;comment:角色ID;size:90"` // 角色ID
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
	AuthorityName string         `gorm:"comment:角色名"` // 角色名

	Users []User `gorm:"many2many:user_authority"`
}
