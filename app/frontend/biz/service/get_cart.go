package service

import (
	"context"
	common "github.com/7qing/gomall/api/frontend/hertz_gen/frontend/common"
	"github.com/7qing/gomall/api/frontend/infra/rpc"
	frontendUtils "github.com/7qing/gomall/api/frontend/utils"
	"github.com/7qing/gomall/rpc_gen/kitex_gen/cart"
	"github.com/7qing/gomall/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"strconv"
)

type GetCartService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetCartService(Context context.Context, RequestContext *app.RequestContext) *GetCartService {
	return &GetCartService{RequestContext: RequestContext, Context: Context}
}

func (h *GetCartService) Run(req *common.Empty) (resp map[string]any, err error) {
	NewCartsRes, err := rpc.CartClient.GetCart(h.Context, &cart.GetCartReq{
		UserId: uint32(frontendUtils.GetUserIdFromCtx(h.Context)),
	})
	if err != nil {
		return nil, err
	}
	var total float64
	var items []map[string]string
	for _, item := range NewCartsRes.Cart.Items {
		NewProductResp, err := rpc.ProductcatalogClient.GetProduct(h.Context, &product.GetProductReq{
			Id: item.ProductId,
		})
		if err != nil {
			continue
		}
		items = append(items, map[string]string{
			"Name":        NewProductResp.Product.Name,
			"Description": NewProductResp.Product.Description,
			"Price":       strconv.FormatFloat(float64(NewProductResp.Product.Price), 'f', 2, 64),
			"Qty":         strconv.Itoa(int(item.Quantity)),
		})
		total += float64(NewProductResp.Product.Price) * float64(item.Quantity)
	}
	return utils.H{
		"title": "Cart",
		"items": items,
		"total": strconv.FormatFloat(total, 'f', 2, 64),
	}, nil
}
