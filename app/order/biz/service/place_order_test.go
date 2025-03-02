package service

import (
	"context"
	"github.com/7qing/gomall/app/order/biz/dal/mysql"
	"github.com/7qing/gomall/app/order/conf"
	"github.com/7qing/gomall/rpc_gen/kitex_gen/cart"
	order "github.com/7qing/gomall/rpc_gen/kitex_gen/order"
	"github.com/joho/godotenv"
	"testing"
)

func TestPlaceOrder_Run(t *testing.T) {
	godotenv.Load(".env")

	conf.InitConf()
	mysql.Init()
	ctx := context.Background()
	s := NewPlaceOrderService(ctx)
	// init req and assert value

	var orders []*order.OrderItem
	orders = append(orders, &order.OrderItem{
		Item: &cart.CartItem{
			ProductId: 20,
			Quantity:  1,
		},
		Cost: 2,
	})
	req := &order.PlaceOrderReq{
		UserId: 2,
		Address: &order.Address{
			StreetAddress: "1",
			City:          "1",
			State:         "1",
			Country:       "1",
			ZipCode:       011111,
		},
		Email:      "1@11.com",
		OrderItems: orders,
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
