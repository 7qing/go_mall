package service

import (
	"context"
	"github.com/7qing/gomall/app/order/biz/dal/mysql"
	"github.com/7qing/gomall/app/order/biz/model"
	order "github.com/7qing/gomall/rpc_gen/kitex_gen/order"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PlaceOrderService struct {
	ctx context.Context
} // NewPlaceOrderService new PlaceOrderService
func NewPlaceOrderService(ctx context.Context) *PlaceOrderService {
	return &PlaceOrderService{ctx: ctx}
}

// Run create note info
func (s *PlaceOrderService) Run(req *order.PlaceOrderReq) (resp *order.PlaceOrderResp, err error) {
	// Step 1: Validate the request
	// 检查订单项是否为空，如果为空，则返回错误
	if req.OrderItems == nil {
		return nil, kerrors.NewBizStatusError(500001, "orderItems is empty")
	}

	// Step 2: Begin a transaction
	// 开始数据库事务
	mysql.DB.Transaction(func(tx *gorm.DB) error {
		// Generate a new UUID for the order
		// 生成新的订单ID
		orderId, _ := uuid.NewUUID()

		// Create an order object with the provided details
		// 创建订单对象，初始化订单的基础信息
		o := &model.Order{
			OrderId:      orderId.String(), // 订单ID
			UserId:       req.UserId,       // 用户ID
			UserCurrency: req.UserCurrency, // 用户货币
			Consignee: model.Consignee{ // 收货人信息
				Email: req.Email, // 收货人邮箱
			},
		}

		// Step 3: Handle the address, if provided
		// 如果提供了地址信息，将其存储在 Consignee 中
		if req.Address != nil {
			a := req.Address
			o.Consignee.StreetAddress = a.StreetAddress
			o.Consignee.City = a.City
			o.Consignee.State = a.State
			o.Consignee.ZipCode = a.ZipCode
			o.Consignee.Country = a.Country
		}

		// Step 4: Save the order object to the database
		// 将订单保存到数据库中
		if err := tx.Create(o).Error; err != nil {
			return err // 如果保存订单失败，返回错误
		}

		// Step 5: Prepare order items for insertion
		// 构建订单项数组，将每个商品转换为 OrderItem 对象
		var items []model.OrderItem
		for _, item := range req.OrderItems {
			items = append(items, model.OrderItem{
				ProductId:    item.Item.ProductId,        // 商品ID
				OrderIdRefer: orderId.String(),           // 订单ID
				Quantity:     uint32(item.Item.Quantity), // 商品数量
				Cost:         item.Cost,                  // 商品成本
			})
		}

		// Step 6: Save the order items to the database
		// 将订单项保存到数据库中
		if err := tx.Create(items).Error; err != nil {
			// 可以在此处处理错误，或者日志记录
		}

		// Step 7: Create and return the response
		// 创建响应对象并返回
		resp = &order.PlaceOrderResp{
			Order: &order.OrderResult{
				OrderId: orderId.String(), // 返回订单ID
			},
		}

		// Commit transaction
		// 提交事务
		return nil
	})

	// Step 8: Return the response
	// 返回订单响应
	return
}
