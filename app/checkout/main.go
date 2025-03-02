package main

import (
	"github.com/7qing/gomall/app/checkout/infra/rpc"
	"github.com/hashicorp/consul/api"
	consul "github.com/kitex-contrib/registry-consul"
	"net"
	"time"

	"github.com/7qing/gomall/app/checkout/conf"
	"github.com/7qing/gomall/rpc_gen/kitex_gen/checkout/checkoutservice"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	ServiceName  = conf.GetConf().Kitex.Service
	RegisterAddr = conf.GetConf().Kitex.Address
)

func main() {

	//mtl.InitMetrics(ServiceName, conf.GetConf().Kitex.MetricsPort, RegisterAddr)
	//p := mtl.InitTracing(ServiceName)
	//defer p.Shutdown(context.Background())
	rpc.Init()
	opts := kitexInit()

	svr := checkoutservice.NewServer(new(CheckoutServiceImpl), opts...)

	err := svr.Run()
	if err != nil {
		klog.Error(err.Error())
	}
}

func kitexInit() (opts []server.Option) {
	// address
	addr, err := net.ResolveTCPAddr("tcp", conf.GetConf().Kitex.Address)
	if err != nil {
		panic(err)
	}
	opts = append(opts, server.WithServiceAddr(addr))

	consulConfig := api.Config{
		Address: conf.GetConf().Registry.RegistryAddress[0],
		Scheme:  "http",
		Token:   "TEST-MY-TOKEN",
	}
	//r, err := consul.NewConsulRegisterWithConfig(&consulConfig, consul.WithCheck(&api.AgentServiceCheck{
	//	Interval:                       "7s",
	//	Timeout:                        "5s",
	//	DeregisterCriticalServiceAfter: "1m",
	//}))
	r, err := consul.NewConsulRegisterWithConfig(&consulConfig)
	if err != nil {
		klog.Fatal(err)
	}

	// service info 添加注册信息
	opts = append(opts, server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
		ServiceName: conf.GetConf().Kitex.Service,
	}), server.WithRegistry(r))

	//opts = append(opts, server.WithServiceAddr(addr), server.WithSuite(serversuite.CommonServerSuite{
	//	CurrentServiceName: ServiceName,
	//	RegistryAddr:       RegisterAddr,
	//}))

	// klog
	logger := kitexlogrus.NewLogger()
	klog.SetLogger(logger)
	klog.SetLevel(conf.LogLevel())
	asyncWriter := &zapcore.BufferedWriteSyncer{
		WS: zapcore.AddSync(&lumberjack.Logger{
			Filename:   conf.GetConf().Kitex.LogFileName,
			MaxSize:    conf.GetConf().Kitex.LogMaxSize,
			MaxBackups: conf.GetConf().Kitex.LogMaxBackups,
			MaxAge:     conf.GetConf().Kitex.LogMaxAge,
		}),
		FlushInterval: time.Minute,
	}
	klog.SetOutput(asyncWriter)
	server.RegisterShutdownHook(func() {
		asyncWriter.Sync()
	})
	return
}
