package service

import (
	"context"
	"github.com/7qing/gomall/app/order/biz/dal/mysql"
	"github.com/7qing/gomall/app/order/biz/model"
	"github.com/7qing/gomall/rpc_gen/kitex_gen/cart"
	order "github.com/7qing/gomall/rpc_gen/kitex_gen/order"
	"github.com/cloudwego/kitex/pkg/kerrors"
)

type ListOrderService struct {
	ctx context.Context
} // NewListOrderService new ListOrderService
func NewListOrderService(ctx context.Context) *ListOrderService {
	return &ListOrderService{ctx: ctx}
}

// Run create note info
func (s *ListOrderService) Run(req *order.ListOrderReq) (resp *order.ListOrderResp, err error) {
	// Finish your business logic.
	ordersresp, err := model.ListOrders(s.ctx, mysql.DB, req.UserId)
	if err != nil {
		return nil, kerrors.NewBizStatusError(500001, "orderItems is empty")
	}

	var orders []*order.Order
	for _, rangeorder := range ordersresp {
		var items []*order.OrderItem
		for _, item := range rangeorder.OrderItems {
			items = append(items, &order.OrderItem{
				Item: &cart.CartItem{
					ProductId: item.ProductId,
					Quantity:  int32(item.Quantity),
				},
				Cost: item.Cost,
			})
		}

		orders = append(orders, &order.Order{
			OrderItems:   items,
			OrderId:      rangeorder.OrderId,
			UserId:       rangeorder.UserId,
			UserCurrency: rangeorder.UserCurrency,
			Address: &order.Address{
				StreetAddress: rangeorder.Consignee.StreetAddress,
				City:          rangeorder.Consignee.City,
				State:         rangeorder.Consignee.State,
				Country:       rangeorder.Consignee.Country,
				ZipCode:       rangeorder.Consignee.ZipCode,
			},
			Email:     rangeorder.Consignee.Email,
			CreatedAt: rangeorder.Model.CreatedAt.Unix(),
		})
	}
	resp = &order.ListOrderResp{
		Orders: orders,
	}
	return
}
