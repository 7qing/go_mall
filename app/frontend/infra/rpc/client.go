package rpc

import (
	"github.com/7qing/gomall/api/frontend/conf"
	"github.com/7qing/gomall/common/clientsuite"
	"github.com/7qing/gomall/rpc_gen/kitex_gen/auth/authservice"
	"github.com/7qing/gomall/rpc_gen/kitex_gen/cart/cartservice"
	"github.com/7qing/gomall/rpc_gen/kitex_gen/checkout/checkoutservice"
	"github.com/7qing/gomall/rpc_gen/kitex_gen/order/orderservice"
	"github.com/7qing/gomall/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/7qing/gomall/rpc_gen/kitex_gen/user/userservice"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/client"
	"sync"
)

// 初始化外部依赖的RPC微服务
// 客户端使用代码（前端使用）

const serviceName = "frontend"

var (
	UserClient           userservice.Client
	ProductcatalogClient productcatalogservice.Client
	CartClient           cartservice.Client
	CheckoutClient       checkoutservice.Client
	OrderClient          orderservice.Client
	AuthClient           authservice.Client
	once                 sync.Once
	ServiceName          = serviceName
	MetricsPort          = conf.GetConf().Hertz.MetricsPort
	RegistryAddress      = conf.GetConf().Hertz.RegistryAddress[0]
	err                  error
)

func Init() {
	// only once
	once.Do(func() {
		initUserClient()
		initProductCatalogClient()
		initCartClient()
		initCheckoutClient()
		initOrderClient()
		initAuthClient()
	})
}

// iniUserClient 用于初始化 UserClient 客户端并与 Consul 注册中心进行连接
func initUserClient() {
	// 使用 Consul 注册中心解析器创建用户服务的客户端
	UserClient, err = userservice.NewClient("user", client.WithSuite(clientsuite.CommonClientSuite{
		CurrentServiceName: ServiceName,
		RegistryAddr:       RegistryAddress,
	}))
	if err != nil {
		// 如果客户端创建失败，记录错误并终止程序
		hlog.Fatal(err)
	}
}

func initProductCatalogClient() {

	// 使用 Consul 注册中心解析器创建商品服务的客户端
	ProductcatalogClient, err = productcatalogservice.NewClient("product", client.WithSuite(clientsuite.CommonClientSuite{
		CurrentServiceName: ServiceName,
		RegistryAddr:       RegistryAddress,
	}))
	// 抽象出错误处理模块
	MustHandelError(err)
}

func initCartClient() {
	// 使用 Consul 注册中心解析器创建商品服务的客户端
	CartClient, err = cartservice.NewClient("cart", client.WithSuite(clientsuite.CommonClientSuite{
		CurrentServiceName: ServiceName,
		RegistryAddr:       RegistryAddress,
	}))
	// 抽象出错误处理模块
	MustHandelError(err)
}

func initCheckoutClient() {
	// 使用 Consul 注册中心解析器创建商品服务的客户端
	CheckoutClient, err = checkoutservice.NewClient("checkout", client.WithSuite(clientsuite.CommonClientSuite{
		CurrentServiceName: ServiceName,
		RegistryAddr:       RegistryAddress,
	}))
	// 抽象出错误处理模块
	MustHandelError(err)
}

func initOrderClient() {
	// 使用 Consul 注册中心解析器创建商品服务的客户端
	OrderClient, err = orderservice.NewClient("order", client.WithSuite(clientsuite.CommonClientSuite{
		CurrentServiceName: ServiceName,
		RegistryAddr:       RegistryAddress,
	}))
	// 抽象出错误处理模块
	MustHandelError(err)
}

func initAuthClient() {
	// 使用 Consul 注册中心解析器创建认证服务的客户端
	AuthClient, err = authservice.NewClient("auth", client.WithSuite(clientsuite.CommonClientSuite{
		CurrentServiceName: ServiceName,
		RegistryAddr:       RegistryAddress,
	}))
	// 抽象出错误处理模块
	MustHandelError(err)
}

func MustHandelError(err error) {
	if err != nil {
		hlog.Fatal(err)
	}
}
