package service

import (
	"context"
	"github.com/7qing/gomall/app/product/biz/dal/mysql"
	"github.com/7qing/gomall/app/product/conf"
	product "github.com/7qing/gomall/rpc_gen/kitex_gen/product"
	"github.com/joho/godotenv"
	"os"
	"testing"
)

func TestDelProduct_Run(t *testing.T) {
	godotenv.Load("../../.env")
	projectRoot := "../../" // 你可以在此修改为根目录路径
	os.Chdir(projectRoot)

	conf.InitConf()
	mysql.Init()
	ctx := context.Background()
	s := NewDelProductService(ctx)
	// init req and assert value

	req := &product.DelProductReq{
		Name: "feizhouren",
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
