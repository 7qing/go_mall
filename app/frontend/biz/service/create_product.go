package service

import (
	"context"
	"github.com/7qing/gomall/api/frontend/infra/rpc"
	"github.com/cloudwego/kitex/pkg/klog"

	common "github.com/7qing/gomall/api/frontend/hertz_gen/frontend/common"
	product "github.com/7qing/gomall/api/frontend/hertz_gen/frontend/product"
	rpcproduct "github.com/7qing/gomall/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/app"
)

type CreateProductService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCreateProductService(Context context.Context, RequestContext *app.RequestContext) *CreateProductService {
	return &CreateProductService{RequestContext: RequestContext, Context: Context}
}

func (h *CreateProductService) Run(req *product.CreateProductReq) (resp *common.Empty, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	var categories []string
	categories = append(categories, req.Cate)
	CreateProduct := rpcproduct.Product{
		Name:        req.Name,
		Description: req.Description,
		Picture:     req.Picture,
		Price:       req.Price,
		Categories:  categories,
	}
	_, err = rpc.ProductcatalogClient.CreateProduct(h.Context, &rpcproduct.CreateProductReq{
		Product: &CreateProduct,
	})
	if err != nil {
	}
	klog.Info("create product success")
	resp = &common.Empty{}
	return
}
