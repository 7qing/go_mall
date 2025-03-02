package rpc

import (
	"github.com/7qing/gomall/app/cart/conf"
	"github.com/7qing/gomall/app/cart/utils"
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
	once                 sync.Once
	ServiceName          = conf.GetConf().Kitex.Service
	RegisterAddr         = conf.GetConf().Kitex.Address
	err                  error
)

func Init() {
	// only once
	once.Do(func() {
		initProductCatalogClient()
	})
}

func initProductCatalogClient() {

	// 使用 Consul 注册中心解析器创建商品服务的客户端
	//ProductcatalogClient, err = productcatalogservice.NewClient("product", client.WithSuite(clientsuite.CommonClientSuite{
	//	CurrentServiceName: ServiceName,
	//	RegistryAddr:       RegisterAddr,
	//}))
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
	utils.MustHandelError(err)
	opts = append(opts, client.WithResolver(r))
	// 使用 Consul 注册中心解析器创建商品服务的客户端
	ProductcatalogClient, err = productcatalogservice.NewClient("product", opts...)
	// 抽象出错误处理模块
	utils.MustHandelError(err)
}
