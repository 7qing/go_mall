package service

import (
	"context"
	"github.com/7qing/gomall/app/order/biz/dal/mysql"
	"github.com/7qing/gomall/app/order/conf"
	order "github.com/7qing/gomall/rpc_gen/kitex_gen/order"
	"github.com/joho/godotenv"
	"testing"
)

func TestListOrder_Run(t *testing.T) {
	godotenv.Load(".env")

	conf.InitConf()
	mysql.Init()
	ctx := context.Background()
	s := NewListOrderService(ctx)
	// init req and assert value

	req := &order.ListOrderReq{
		UserId: 2,
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
