package model

import (
	"context"
	"errors"
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	UserId uint64 `gorm:"type:bigint(11);not null;index:idx_user_id"`
	SpuId  uint64 `gorm:"type:bigint(11);not null;"`
	SkuId  uint64 `gorm:"type:bigint(11);not null;"`
	Qty    uint32 `gorm:"type:int(11);not null;"`
}

func (Cart) TableName() string {
	return "cart"
}

func AddItem(ctx context.Context, db *gorm.DB, item *Cart) error {
	var find Cart
	err := db.WithContext(ctx).Model(&Cart{}).Where(&Cart{
		UserId: item.UserId,
		SpuId:  item.SpuId,
		SkuId:  item.SkuId,
	}).First(&find).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	if find.ID != 0 {
		err = db.WithContext(ctx).Model(&Cart{}).Where(&Cart{
			UserId: item.UserId,
			SpuId:  item.SpuId,
			SkuId:  item.SkuId}).UpdateColumn("qty", gorm.Expr("qty+?", item.Qty)).Error
	} else {
		err = db.WithContext(ctx).Model(&Cart{}).Create(item).Error
	}
	return err
}

func EmptyCart(ctx context.Context, db *gorm.DB, userId uint64) error {
	if userId == 0 {
		return errors.New("user id is required")
	}
	return db.WithContext(ctx).Delete(&Cart{}, "user_id = ?", userId).Error
}

func GetCartByUserId(ctx context.Context, db *gorm.DB, userId uint64) (cartList []*Cart, err error) {
	err = db.Debug().WithContext(ctx).Model(&Cart{}).Find(&cartList, "user_id = ?", userId).Error
	return cartList, err
}
