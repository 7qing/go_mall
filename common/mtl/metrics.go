package mtl

import (
	"github.com/cloudwego/kitex/pkg/registry"
	"github.com/cloudwego/kitex/server"
	consul "github.com/kitex-contrib/registry-consul"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net"
	"net/http"
)

var Registry *prometheus.Registry

func InitMetrics(serverName string, metricsPort string, registryAddr string) (registry.Registry, *registry.Info) {
	Registry = prometheus.NewRegistry()
	Registry.MustRegister(collectors.NewGoCollector())                                       //go相关的指标
	Registry.MustRegister(collectors.NewProcessCollector(collectors.ProcessCollectorOpts{})) //注册进程相关的指标
	//consul注册中心注册
	r, _ := consul.NewConsulRegister(registryAddr)
	addr, _ := net.ResolveTCPAddr("tcp", metricsPort)
	registryInfo := &registry.Info{
		ServiceName: "prometheus",
		Addr:        addr,
		Weight:      1,
		Tags:        map[string]string{"service": serverName},
	}
	_ = r.Register(registryInfo)

	// 服务关闭时候，下线注册信息：
	server.RegisterShutdownHook(func() {
		err := r.Deregister(registryInfo)
		if err != nil {
			return
		}
	},
	)

	//启动metrics server
	http.Handle("/metrics", promhttp.HandlerFor(Registry, promhttp.HandlerOpts{}))
	//异步启动一个seever供prometheus拉取指标
	go http.ListenAndServe(metricsPort, nil)
	return r, registryInfo
}
