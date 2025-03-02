package service

import (
	"context"
	"github.com/7qing/gomall/app/product/biz/dal/mysql"
	"github.com/7qing/gomall/app/product/biz/model"
	product "github.com/7qing/gomall/rpc_gen/kitex_gen/product"
	"log"
)

type ListProductsService struct {
	ctx context.Context
} // NewListProductsService new ListProductsService
func NewListProductsService(ctx context.Context) *ListProductsService {
	return &ListProductsService{ctx: ctx}
}

// Run create note info
func (s *ListProductsService) Run(req *product.ListProductsReq) (resp *product.ListProductsResp, err error) {
	// Step 1: Log request data
	// 记录接收到的请求数据，便于调试和日志追踪
	log.Printf("Request received: %+v", req)

	// Step 2: Validate input
	// 检查请求是否为空，或者请求中的分类名称是否为空
	if req == nil || req.CategoryName == "" {
		// 如果分类名称为空或请求本身为空，记录日志并返回错误
		log.Println("Invalid request: CategoryName is empty or nil")
	}

	// Step 3: Initialize response object
	// 创建一个空的响应对象，并初始化产品列表
	resp = &product.ListProductsResp{}
	resp.Products = []*product.Product{}

	// Step 4: Create category query object
	// 创建一个新的分类查询对象，用于查询数据库中的产品信息
	categoryQuery := model.NewCategoryQuery(s.ctx, mysql.DB)

	// Step 5: Query products by category
	// 调用查询方法，根据提供的分类名称获取相关的产品
	categories, err := categoryQuery.GetProductsbyCategory(req.CategoryName)
	if err != nil {
		// 如果查询过程中发生错误，返回错误
		return nil, err
	}

	// Step 6: Populate response with products
	// 遍历获取到的分类数据，依次将每个产品的详细信息添加到响应列表中
	for _, category := range categories {
		for _, productRange := range category.Product {
			resp.Products = append(resp.Products, &product.Product{
				Id:          uint32(productRange.ID),  // 产品ID
				Name:        productRange.Name,        // 产品名称
				Description: productRange.Description, // 产品描述
				Picture:     productRange.Picture,     // 产品图片链接
				Price:       productRange.Price,       // 产品价格
			})
		}
	}

	// Step 7: Return the populated response
	// 返回包含产品列表的响应对象
	return resp, err
}
