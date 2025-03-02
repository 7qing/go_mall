package rpc

import (
	"github.com/7qing/gomall/app/checkout/conf"
	"github.com/7qing/gomall/rpc_gen/kitex_gen/cart/cartservice"
	"github.com/7qing/gomall/rpc_gen/kitex_gen/order/orderservice"
	"github.com/7qing/gomall/rpc_gen/kitex_gen/payment/paymentservice"
	"github.com/7qing/gomall/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/cloudwego/kitex/client"
	consulapi "github.com/hashicorp/consul/api"
	consul "github.com/kitex-contrib/registry-consul"
	"sync"
)

// 初始化外部依赖的RPC微服务
// 客户端使用代码（前端使用）

var (
	ProductcatalogClient productcatalogservice.Client
	CartClient           cartservice.Client
	PaymentClient        paymentservice.Client
	OrdermentClient      orderservice.Client
	once                 sync.Once
	ServiceName          = conf.GetConf().Kitex.Service
	RegisterAddr         = conf.GetConf().Kitex.Address
	err                  error
)

func Init() {
	// only once
	once.Do(func() {
		initCartClient()
		initProductCatalogClient()
		initPaymentClient()
		initOrderClient()
	})
}

// iniUserClient 用于初始化 PaymentClient 客户端并与 Consul 注册中心进行连接
func initPaymentClient() {
	// 使用 Consul 注册中心解析器创建用户服务的客户端
	//PaymentClient, err = paymentservice.NewClient("payment", client.WithSuite(clientsuite.CommonClientSuite{
	//	CurrentServiceName: ServiceName,
	//	RegistryAddr:       RegisterAddr,
	//}))
	//if err != nil {
	//	panic(err)
	//}
	var opts []client.Option
	// 创建 Consul 注册中心的解析器
	consulConfig := consulapi.Config{
		Address: conf.GetConf().Registry.RegistryAddress[0],
		Scheme:  "http",
		Token:   "TEST-MY-TOKEN",
	}
	r, err := consul.NewConsulResolverWithConfig(&consulConfig)

	// 如果解析器创建失败，记录错误并终止程序
	// 抽象出错误处理模块
	MustHandelError(err)
	opts = append(opts, client.WithResolver(r))
	// 使用 Consul 注册中心解析器创建商品服务的客户端
	PaymentClient, err = paymentservice.NewClient("payment", opts...)
	// 抽象出错误处理模块
	MustHandelError(err)
}

func initProductCatalogClient() {

	//// 使用 Consul 注册中心解析器创建用户服务的客户端
	//ProductcatalogClient, err = productcatalogservice.NewClient("product", client.WithSuite(clientsuite.CommonClientSuite{
	//	CurrentServiceName: ServiceName,
	//	RegistryAddr:       RegisterAddr,
	//}))
	//if err != nil {
	//	panic(err)
	//}
	var opts []client.Option
	// 创建 Consul 注册中心的解析器
	consulConfig := consulapi.Config{
		Address: conf.GetConf().Registry.RegistryAddress[0],
		Scheme:  "http",
		Token:   "TEST-MY-TOKEN",
	}
	r, err := consul.NewConsulResolverWithConfig(&consulConfig)

	// 如果解析器创建失败，记录错误并终止程序
	// 抽象出错误处理模块
	MustHandelError(err)
	opts = append(opts, client.WithResolver(r))
	// 使用 Consul 注册中心解析器创建商品服务的客户端
	ProductcatalogClient, err = productcatalogservice.NewClient("product", opts...)
	// 抽象出错误处理模块
	MustHandelError(err)
}

func initCartClient() {
	//// 使用 Consul 注册中心解析器创建用户服务的客户端
	//CartClient, err = cartservice.NewClient("cart", client.WithSuite(clientsuite.CommonClientSuite{
	//	CurrentServiceName: ServiceName,
	//	RegistryAddr:       RegisterAddr,
	//}))
	//if err != nil {
	//	panic(err)
	//}
	var opts []client.Option
	// 创建 Consul 注册中心的解析器
	consulConfig := consulapi.Config{
		Address: conf.GetConf().Registry.RegistryAddress[0],
		Scheme:  "http",
		Token:   "TEST-MY-TOKEN",
	}
	r, err := consul.NewConsulResolverWithConfig(&consulConfig)

	// 如果解析器创建失败，记录错误并终止程序
	// 抽象出错误处理模块
	MustHandelError(err)
	opts = append(opts, client.WithResolver(r))
	// 使用 Consul 注册中心解析器创建商品服务的客户端
	CartClient, err = cartservice.NewClient("cart", opts...)
	// 抽象出错误处理模块
	MustHandelError(err)
}

func initOrderClient() {
	//// 使用 Consul 注册中心解析器创建用户服务的客户端
	//OrdermentClient, err = orderservice.NewClient("order", client.WithSuite(clientsuite.CommonClientSuite{
	//	CurrentServiceName: ServiceName,
	//	RegistryAddr:       RegisterAddr,
	//}))
	//if err != nil {
	//	panic(err)
	//}
	var opts []client.Option
	// 创建 Consul 注册中心的解析器
	consulConfig := consulapi.Config{
		Address: conf.GetConf().Registry.RegistryAddress[0],
		Scheme:  "http",
		Token:   "TEST-MY-TOKEN",
	}
	r, err := consul.NewConsulResolverWithConfig(&consulConfig)

	// 如果解析器创建失败，记录错误并终止程序
	// 抽象出错误处理模块
	MustHandelError(err)
	opts = append(opts, client.WithResolver(r))
	// 使用 Consul 注册中心解析器创建商品服务的客户端
	OrdermentClient, err = orderservice.NewClient("order", opts...)
	// 抽象出错误处理模块
	MustHandelError(err)
}

func MustHandelError(err error) {
	if err != nil {
		panic(err)
	}
}
