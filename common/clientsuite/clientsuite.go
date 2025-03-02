package clientsuite

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/transport"
	consulapi "github.com/hashicorp/consul/api"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	consul "github.com/kitex-contrib/registry-consul"
)

type CommonClientSuite struct {
	CurrentServiceName string
	RegistryAddr       string
}

func (s CommonClientSuite) Options() []client.Option {
	opts := []client.Option{
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: s.CurrentServiceName,
		}),
		client.WithMetaHandler(transmeta.ClientHTTP2Handler),
		client.WithTransportProtocol(transport.GRPC),
		client.WithSuite(tracing.NewClientSuite()),
	}
	// 创建 Consul 注册中心的解析器
	consulConfig := consulapi.Config{
		Address: s.RegistryAddr,
		Scheme:  "http",
		Token:   "TEST-MY-TOKEN",
	}
	r, err := consul.NewConsulResolverWithConfig(&consulConfig)

	// 如果解析器创建失败，记录错误并终止程序
	// 抽象出错误处理模块
	MustHandelError(err)

	opts = append(opts, client.WithResolver(r))
	return opts
}

func MustHandelError(err error) {
	if err != nil {
		hlog.Fatal(err)
	}
}
