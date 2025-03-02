package service

import (
	"context"
	"github.com/7qing/gomall/app/cart/biz/dal/mysql"
	"github.com/7qing/gomall/app/cart/biz/model"
	cart "github.com/7qing/gomall/rpc_gen/kitex_gen/cart"
	"github.com/cloudwego/kitex/pkg/kerrors"
)

type GetCartService struct {
	ctx context.Context
} // NewGetCartService new GetCartService
func NewGetCartService(ctx context.Context) *GetCartService {
	return &GetCartService{ctx: ctx}
}

// Run create note info
func (s *GetCartService) Run(req *cart.GetCartReq) (resp *cart.GetCartResp, err error) {
	// Step 1: Retrieve the user's cart items from the database
	// 根据用户ID获取该用户的购物车项列表
	list, err := model.GetCartByUserId(s.ctx, mysql.DB, req.UserId)
	if err != nil {
		// 如果获取购物车项时发生错误，返回错误
		// todo: 随便定义一个错误码
		return nil, kerrors.NewBizStatusError(50002, err.Error())
	}

	// Step 2: Convert database cart items into response format
	// 将数据库中获取到的购物车项转换为响应格式
	var items []*cart.CartItem
	for _, item := range list {
		// 遍历数据库中的购物车项，将其转换为响应中的 CartItem
		items = append(items, &cart.CartItem{
			ProductId: item.ProductID,  // 产品ID
			Quantity:  int32(item.Qty), // 产品数量
		})
	}
	
	// Step 3: Create response
	// 构建最终的响应对象
	resp = &cart.GetCartResp{
		Cart: &cart.Cart{
			UserId: req.UserId, // 用户ID
			Items:  items,      // 购物车中的产品项
		},
	}

	// Step 4: Return response
	// 返回响应对象
	return resp, nil
}
