package service

import (
	"context"
	"github.com/7qing/gomall/app/cart/biz/dal/mysql"
	"github.com/7qing/gomall/app/cart/conf"
	"github.com/7qing/gomall/app/cart/rpc"

	cart "github.com/7qing/gomall/rpc_gen/kitex_gen/cart"
	"github.com/joho/godotenv"
	"testing"
)

func TestAddItem_Run(t *testing.T) {
	godotenv.Load(".env")

	conf.InitConf()
	mysql.Init()
	rpc.Init()
	ctx := context.Background()
	s := NewAddItemService(ctx)
	// init req and assert value

	Newcart := &cart.CartItem{
		ProductId: 20,
		Quantity:  1,
	}
	req := &cart.AddItemReq{
		UserId: 2,
		Item:   Newcart,
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
