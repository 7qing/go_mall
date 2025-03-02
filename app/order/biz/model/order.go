package model

import (
	"context"
	"gorm.io/gorm"
)

type Consignee struct {
	Email         string
	StreetAddress string
	City          string
	State         string
	Country       string
	ZipCode       int32
}

type Order struct {
	gorm.Model
	OrderId      string      `gorm:"type:varchar(100);uniqueIndex"`
	UserId       uint32      `gorm:"type:int(11)"`
	UserCurrency string      `gorm:"type:varchar(10)"` //币种
	Consignee    Consignee   `gorm:"embedded;"`        //告诉 GORM，将嵌套结构体的字段“嵌入”到当前结构体中，而不是将其当作一个单独的表来处理。
	OrderItems   []OrderItem `gorm:"foreignkey:OrderIdRefer;references:OrderId"`
}

func (Order) TableName() string {
	return "order"
}

//message Order {
//repeated OrderItem order_items = 1;
//string order_id = 2;
//uint32 user_id = 3;
//string user_currency = 4;
//Address address = 5;
//string email = 6;
//int32 created_at = 7;
//}

func ListOrders(ctx context.Context, db *gorm.DB, userId uint32) (orders []*Order, err error) {
	err = db.WithContext(ctx).Where("user_id = ?", userId).Preload("OrderItems").Find(&orders).Error
	if err != nil {
		return nil, err
	}
	return orders, nil
}
