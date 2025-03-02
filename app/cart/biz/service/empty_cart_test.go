package service

import (
	"context"
	"github.com/7qing/gomall/app/cart/biz/dal/mysql"
	"github.com/7qing/gomall/app/cart/conf"

	cart "github.com/7qing/gomall/rpc_gen/kitex_gen/cart"
	"github.com/joho/godotenv"
	"os"
	"testing"
)

func TestEmptyCart_Run(t *testing.T) {
	godotenv.Load("../../.env")
	os.Chdir("../../")
	conf.InitConf()
	mysql.Init()
	ctx := context.Background()
	s := NewEmptyCartService(ctx)
	// init req and assert value

	req := &cart.EmptyCartReq{
		UserId: 2,
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
