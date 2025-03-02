package service

import (
	"context"
	"github.com/7qing/gomall/api/frontend/infra/rpc"
	frontendUtils "github.com/7qing/gomall/api/frontend/utils"
	"github.com/7qing/gomall/app/frontend/types"
	"github.com/7qing/gomall/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/kitex/pkg/klog"
	"time"

	common "github.com/7qing/gomall/api/frontend/hertz_gen/frontend/common"
	"github.com/7qing/gomall/rpc_gen/kitex_gen/order"
	"github.com/cloudwego/hertz/pkg/app"
)

type OrderListService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewOrderListService(Context context.Context, RequestContext *app.RequestContext) *OrderListService {
	return &OrderListService{RequestContext: RequestContext, Context: Context}
}

func (h *OrderListService) Run(req *common.Empty) (resp map[string]any, err error) {
	userId := frontendUtils.GetUserIdFromCtx(h.Context)
	orderResp, err := rpc.OrderClient.ListOrder(h.Context, &order.ListOrderReq{
		UserId: uint32(userId),
	})

	var list []types.Order
	if err != nil || orderResp.Orders == nil {
		klog.Info("No orders found for user:", userId)
		resp = utils.H{
			"title":  "Order",
			"Orders": []types.Order{}, // 返回空列表
		}
		return
	}
	for _, rangOrder := range orderResp.Orders {
		var total float32
		var items []types.OrderItem
		for _, item := range rangOrder.OrderItems {
			productResp, err := rpc.ProductcatalogClient.GetProduct(h.Context, &product.GetProductReq{
				Id: item.Item.ProductId,
			})
			if err != nil {
				continue
			}
			total += item.Cost
			items = append(items, types.OrderItem{
				ProductName: productResp.Product.Name,
				Picture:     productResp.Product.Picture,
				Qty:         item.Item.Quantity,
				Cost:        item.Cost,
			})
		}
		list = append(list, types.Order{
			OrderId:    rangOrder.OrderId,
			CreateDate: time.Unix(rangOrder.CreatedAt, 0).Format("2006-01-02 15:04:05"),
			Cost:       total,
			Items:      items,
		})
	}
	resp = utils.H{
		"title":  "Order",
		"Orders": list,
	}
	return
}
