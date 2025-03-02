package service

import (
	"context"
	"github.com/7qing/gomall/api/frontend/infra/rpc"
	rpcProduct "github.com/7qing/gomall/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/common/utils"

	product "github.com/7qing/gomall/api/frontend/hertz_gen/frontend/product"
	"github.com/cloudwego/hertz/pkg/app"
)

type SearchProducsService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewSearchProducsService(Context context.Context, RequestContext *app.RequestContext) *SearchProducsService {
	return &SearchProducsService{RequestContext: RequestContext, Context: Context}
}

func (h *SearchProducsService) Run(req *product.SearchProductsReq) (resp map[string]any, err error) {

	// todo edit your code
	NewProduct, err := rpc.ProductcatalogClient.SearchProducts(h.Context, &rpcProduct.SearchProductsReq{
		Query: req.Q,
	})
	if err != nil {
		return nil, err
	}
	resp = utils.H{
		"items": NewProduct.Results,
		"q":     req.Q,
	}
	return resp, nil
}
