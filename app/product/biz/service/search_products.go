package service

import (
	"context"
	"github.com/7qing/gomall/app/product/biz/dal/mysql"
	"github.com/7qing/gomall/app/product/biz/model"
	product "github.com/7qing/gomall/rpc_gen/kitex_gen/product"
)

type SearchProductsService struct {
	ctx context.Context
} // NewSearchProductsService new SearchProductsService
func NewSearchProductsService(ctx context.Context) *SearchProductsService {
	return &SearchProductsService{ctx: ctx}
}

// Run create note info
func (s *SearchProductsService) Run(req *product.SearchProductsReq) (resp *product.SearchProductsResp, err error) {
	// Step 1: Initialize product query object
	// 创建一个新的产品查询对象，用于执行数据库查询
	productQuery := model.NewProductQuery(s.ctx, mysql.DB)

	// Step 2: Search products by query string
	// 使用查询字符串来搜索产品，调用SearchProducts方法进行搜索
	queryProducts, err := productQuery.SearchProducts(req.Query)
	if err != nil {
		// 如果查询过程中发生错误，返回错误
		return nil, err
	}

	// Step 3: Process search results
	// 遍历查询结果，将每个产品的信息转换成返回的产品结构体
	var products []*product.Product
	for _, product_range := range queryProducts {
		products = append(products, &product.Product{
			Id:          uint32(product_range.ID),  // 产品ID
			Name:        product_range.Name,        // 产品名称
			Description: product_range.Description, // 产品描述
			Picture:     product_range.Picture,     // 产品图片链接
			Price:       product_range.Price,       // 产品价格
		})
	}

	// Step 4: Construct the response
	// 创建并填充响应对象，将处理后的产品列表添加到响应结果中
	resp = &product.SearchProductsResp{
		Results: products, // 搜索结果（产品列表）
	}

	// Step 5: Return the response
	// 返回包含搜索结果的响应对象
	return
}
