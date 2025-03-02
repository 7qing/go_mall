package service

import (
	"context"

	cart "github.com/7qing/gomall/api/frontend/hertz_gen/frontend/cart"
	common "github.com/7qing/gomall/api/frontend/hertz_gen/frontend/common"
	"github.com/7qing/gomall/api/frontend/infra/rpc"
	frontendUtils "github.com/7qing/gomall/api/frontend/utils"
	rpccart "github.com/7qing/gomall/rpc_gen/kitex_gen/cart"
	"github.com/cloudwego/hertz/pkg/app"
)

type AddCartItemService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewAddCartItemService(Context context.Context, RequestContext *app.RequestContext) *AddCartItemService {
	return &AddCartItemService{RequestContext: RequestContext, Context: Context}
}

func (h *AddCartItemService) Run(req *cart.AddCartReq) (resp *common.Empty, err error) {
	_, err = rpc.CartClient.AddItem(h.Context, &rpccart.AddItemReq{
		UserId: uint32(frontendUtils.GetUserIdFromCtx(h.Context)),
		Item: &rpccart.CartItem{
			ProductId: req.ProductId,
			Quantity:  req.ProductNum,
		},
	})
	if err != nil {
		return nil, err
	}
	return
}
