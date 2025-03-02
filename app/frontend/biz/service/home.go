package service

import (
	"context"
	common "github.com/7qing/gomall/api/frontend/hertz_gen/frontend/common"
	"github.com/7qing/gomall/api/frontend/infra/rpc"
	rpcProduct "github.com/7qing/gomall/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/kitex/pkg/klog"
)

type HomeService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewHomeService(Context context.Context, RequestContext *app.RequestContext) *HomeService {
	return &HomeService{RequestContext: RequestContext, Context: Context}
}

func (h *HomeService) Run(req *common.Empty) (resp map[string]any, err error) {
	// todo edit your code product api
	// 留个坑
	//这里还需要做页面得展示（各个东西）
	//var cartNum int
	//return utils.H{
	//	"title":    "Hot sale",
	//	"cart_num": cartNum,
	//	"items":    p.Products,
	//}, nil
	p, err := rpc.ProductcatalogClient.ListProducts(h.Context, &rpcProduct.ListProductsReq{
		CategoryName: "test",
	})
	if err != nil {
		klog.Error(err)
	}
	//resp = make(map[string]any)
	resp = map[string]any{
		"title": "Hot sale",
		"items": p.Products,
	}
	return resp, nil
}
