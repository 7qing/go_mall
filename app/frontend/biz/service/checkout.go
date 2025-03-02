package service

import (
	"context"
	"github.com/7qing/gomall/api/frontend/infra/rpc"
	frontendUtils "github.com/7qing/gomall/api/frontend/utils"
	rpccart "github.com/7qing/gomall/rpc_gen/kitex_gen/cart"
	"github.com/7qing/gomall/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"strconv"

	checkout "github.com/7qing/gomall/api/frontend/hertz_gen/frontend/checkout"
	"github.com/cloudwego/hertz/pkg/app"
)

type CheckoutService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCheckoutService(Context context.Context, RequestContext *app.RequestContext) *CheckoutService {
	return &CheckoutService{RequestContext: RequestContext, Context: Context}
}

func (h *CheckoutService) Run(req *checkout.CheckoutReq) (resp map[string]any, err error) {

	userId := frontendUtils.GetUserIdFromCtx(h.Context)

	cartResp, err := rpc.CartClient.GetCart(h.Context, &rpccart.GetCartReq{
		UserId: uint32(userId),
	})
	if err != nil {
		return nil, err
	}
	var total float64
	var items []map[string]string

	for _, item := range cartResp.Cart.Items {
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
		"title": "Checkout",
		"items": items,
		"total": strconv.FormatFloat(total, 'f', 2, 64),
	}, nil
}
