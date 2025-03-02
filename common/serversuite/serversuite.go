package serversuite

import (
	"github.com/7qing/gomall/common/mtl"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/server"
	"github.com/hashicorp/consul/api"
	prometheus "github.com/kitex-contrib/monitor-prometheus"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	consul "github.com/kitex-contrib/registry-consul"
)

//serversuite是一种高级抽象封装了

type CommonServerSuite struct {
	CurrentServiceName string
	RegistryAddr       string
}

func (s CommonServerSuite) Options() []server.Option {
	opts := []server.Option{
		server.WithMetaHandler(transmeta.ServerHTTP2Handler),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: s.CurrentServiceName,
		}),
		server.WithTracer(prometheus.NewServerTracer("", "",
			prometheus.WithDisableServer(true),
			prometheus.WithRegistry(mtl.Registry)),
		),
		server.WithSuite(tracing.NewServerSuite()),
	}

	// Consul代码实现
	consulConfig := api.Config{
		Address: s.RegistryAddr,
		Scheme:  "http",
		Token:   "TEST-MY-TOKEN",
	}

	r, err := consul.NewConsulRegisterWithConfig(&consulConfig)
	if err != nil {
		klog.Fatal(err)
	}
	opts = append(opts, server.WithRegistry(r))

	return opts
}
