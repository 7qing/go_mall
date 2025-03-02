package main

import (
	"github.com/7qing/gomall/rpc_gen/kitex_gen/user/userservice"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
)

func main() {
	// 创建 Consul 注册中心的解析器
	r, err := consul.NewConsulResolver("127.0.0.1:8500")
	if err != nil {
		// 如果解析器创建失败，记录错误并终止程序
		hlog.Fatal(err)
	}

	// 使用 Consul 注册中心解析器创建用户服务的客户端
	_, err = userservice.NewClient("user", client.WithResolver(r))
	if err != nil {
		// 如果客户端创建失败，记录错误并终止程序
		hlog.Fatal(err)
	}
}
