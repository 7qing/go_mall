package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/utils"

	category "github.com/7qing/gomall/api/frontend/hertz_gen/frontend/category"

	"github.com/7qing/gomall/api/frontend/infra/rpc"
	rpcProduct "github.com/7qing/gomall/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/app"
)

type CategoryService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCategoryService(Context context.Context, RequestContext *app.RequestContext) *CategoryService {
	return &CategoryService{RequestContext: RequestContext, Context: Context}
}

func (h *CategoryService) Run(req *category.CategoryReq) (resp map[string]any, err error) {
	// todo edit your code
	NewCategory, err := rpc.ProductcatalogClient.ListProducts(h.Context, &rpcProduct.ListProductsReq{
		CategoryName: req.Category,
	})
	if NewCategory.Products == nil {
		return resp, err
	}
	if err != nil {
		return nil, err
	}
	resp = utils.H{
		"title": "Category",
		"items": NewCategory.Products,
	}
	return resp, nil
}
