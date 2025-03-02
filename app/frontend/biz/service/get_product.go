package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/utils"

	product "github.com/7qing/gomall/api/frontend/hertz_gen/frontend/product"
	"github.com/7qing/gomall/api/frontend/infra/rpc"
	rpcProduct "github.com/7qing/gomall/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/app"
)

type GetProductService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetProductService(Context context.Context, RequestContext *app.RequestContext) *GetProductService {
	return &GetProductService{RequestContext: RequestContext, Context: Context}
}

func (h *GetProductService) Run(req *product.ProductReq) (resp map[string]any, err error) {

	// todo edit your code
	p, err := rpc.ProductcatalogClient.GetProduct(h.Context, &rpcProduct.GetProductReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	resp = utils.H{
		"item": p.Product,
	}
	return resp, nil
}
