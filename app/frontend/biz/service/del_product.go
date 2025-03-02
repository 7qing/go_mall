package service

import (
	"context"
	"github.com/7qing/gomall/api/frontend/infra/rpc"
	rpcProduct "github.com/7qing/gomall/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/klog"

	common "github.com/7qing/gomall/api/frontend/hertz_gen/frontend/common"
	product "github.com/7qing/gomall/api/frontend/hertz_gen/frontend/product"
	"github.com/cloudwego/hertz/pkg/app"
)

type DelProductService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewDelProductService(Context context.Context, RequestContext *app.RequestContext) *DelProductService {
	return &DelProductService{RequestContext: RequestContext, Context: Context}
}

func (h *DelProductService) Run(req *product.DelProductReq) (resp *common.Empty, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	res, err := rpc.ProductcatalogClient.DelProduct(h.Context, &rpcProduct.DelProductReq{
		Name: req.Name,
	})
	if err != nil {
		klog.Info("Del product err: ", err)
	}
	if res.Res == true {

	}
	resp = &common.Empty{}
	return
}
