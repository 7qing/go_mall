package model

import (
	"context"
	"errors"
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	UserID    uint32 `gorm:"type:int(11);not null;index:idx_user_id"`
	ProductID uint32 `gorm:"type:int(11);not null;"`
	Qty       uint32 `gorm:"type:int(11);not null;"` //数量
}

func (Cart) TableName() string {
	return "cart"
}

func AddItem(ctx context.Context, db *gorm.DB, cart *Cart) error {
	var cartResult Cart
	err := db.WithContext(ctx).Model(&Cart{}).Where(&Cart{UserID: cart.UserID}).First(&cartResult).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return db.WithContext(ctx).Create(cart).Error
	}
	if cartResult.ID != 0 {
		return db.WithContext(ctx).Model(&Cart{}).Where(&Cart{UserID: cart.UserID, ProductID: cart.ProductID}).
			UpdateColumn("qty", gorm.Expr("qty+?", cart.Qty)).Error
	}
	// 没有就创建一个新的cart类给他用
	return nil
}

// 删除
func EmptyCart(ctx context.Context, db *gorm.DB, userID uint32) error {

	if userID == 0 {
		return errors.New("userID can not be 0")
	}
	return db.WithContext(ctx).Delete(&Cart{}, "user_id = ?", userID).Error
}

func GetCartByUserId(ctx context.Context, db *gorm.DB, userID uint32) ([]*Cart, error) {
	if userID == 0 {
		return nil, errors.New("userID can not be 0")
	}
	var rows []*Cart
	err := db.WithContext(ctx).Model(&Cart{}).
		Where(&Cart{UserID: userID}).Find(&rows).Error
	return rows, err
}
func CreateUser(db *gorm.DB, cart *Cart) error {
	return db.Create(cart).Error
}
