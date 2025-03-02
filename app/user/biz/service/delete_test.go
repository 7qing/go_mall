package service

import (
	"context"
	"github.com/7qing/gomall/app/user/biz/dal/mysql"
	"github.com/7qing/gomall/app/user/conf"
	user "github.com/7qing/gomall/rpc_gen/kitex_gen/user"
	"github.com/joho/godotenv"
	"os"
	"testing"
)

func TestDelete_Run(t *testing.T) {
	godotenv.Load("../../.env")
	projectRoot := "../../" // 你可以在此修改为根目录路径
	os.Chdir(projectRoot)

	conf.InitConf()
	mysql.Init()
	ctx := context.Background()
	s := NewDeleteService(ctx)

	// init req and assert value

	req := &user.DeleteReq{
		UserId: 5,
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
