package service

import (
	"context"
	"github.com/7qing/gomall/app/cart/biz/dal/mysql"
	modelcart "github.com/7qing/gomall/app/cart/biz/model"
	"github.com/7qing/gomall/app/cart/rpc"
	cart "github.com/7qing/gomall/rpc_gen/kitex_gen/cart"
	rpcProduct "github.com/7qing/gomall/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/kerrors"
)

type AddItemService struct {
	ctx context.Context
} // NewAddItemService new AddItemService
func NewAddItemService(ctx context.Context) *AddItemService {
	return &AddItemService{ctx: ctx}
}

// Run create note info
func (s *AddItemService) Run(req *cart.AddItemReq) (resp *cart.AddItemResp, err error) {
	// Step 1: Get product details from Product Catalog service
	// 从产品目录服务获取指定产品的详细信息
	NewProduct, err := rpc.ProductcatalogClient.GetProduct(s.ctx, &rpcProduct.GetProductReq{
		Id: req.Item.ProductId, // 使用请求中的产品ID来获取产品详细信息
	})
	if err != nil {
		// 如果获取产品信息时发生错误，返回错误
		return nil, err
	}
	if NewProduct == nil || NewProduct.Product.Id == 0 {
		// 如果未能找到产品，返回错误，表示产品不存在
		// todo: 随便定义一个错误码
		return nil, kerrors.NewBizStatusError(40000, "product not found")
	}

	// Step 2: Create cart item
	// 根据请求中的数据创建一个购物车项
	cartItem := &modelcart.Cart{
		UserID:    req.UserId,                // 用户ID
		ProductID: req.Item.ProductId,        // 产品ID
		Qty:       uint32(req.Item.Quantity), // 产品数量
	}

	// Step 3: Add item to cart
	// 将购物车项添加到数据库
	err = modelcart.AddItem(s.ctx, mysql.DB, cartItem)
	if err != nil {
		// 如果添加购物车项时发生错误，返回错误
		// todo: 随便定义一个错误码
		return nil, kerrors.NewBizStatusError(50000, err.Error())
	}

	// Step 4: Return successful response
	// 添加成功，返回空响应
	return &cart.AddItemResp{}, nil
}
