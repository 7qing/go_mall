package service

import (
	"context"
	"errors"
	"github.com/7qing/gomall/app/product/biz/dal/mysql"
	"github.com/7qing/gomall/app/product/biz/model"
	product "github.com/7qing/gomall/rpc_gen/kitex_gen/product"
	"gorm.io/gorm"
)

type CreateProductService struct {
	ctx context.Context
} // NewCreateProductService new CreateProductService
func NewCreateProductService(ctx context.Context) *CreateProductService {
	return &CreateProductService{ctx: ctx}
}

// Run create note info
func (s *CreateProductService) Run(req *product.CreateProductReq) (resp *product.CreateProductResp, err error) {
	// Finish your business logic.
	if req.Product.Name == "" {
		return nil, errors.New("product name is required")
	}
	Newproduct := model.Product{
		Name:        req.Product.Name,
		Description: req.Product.Description,
		Picture:     req.Product.Picture,
		Price:       req.Product.Price,
	}
	Query := model.NewCategoryQuery(s.ctx, mysql.DB)
	var categories []model.Category
	// 查找类别是否已经存在

	for _, categoryName := range req.Product.Categories {
		var category model.Category
		// 查找类别是否已经存在
		err := Query.FindCategoryByName(categoryName, &category)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err // 查询数据库时发生错误
		}

		// 如果没有找到，创建新的 Category
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err := Query.CreateCategory(categoryName, "")
			if err != nil {
				return nil, err
			}
			// 新创建类别后，重新查询数据库获取该类别
			err = Query.FindCategoryByName(categoryName, &category)
			if err != nil {
				return nil, err
			}
		}
		// 将类别添加到 categories 列表中
		categories = append(categories, category)
	}
	// 4. 将找到的或创建的类别与 Product 关联
	Newproduct.Categories = categories
	ProductQuery := model.NewProductQuery(s.ctx, mysql.DB)
	// 5. 保存产品
	err = ProductQuery.CreateProduct(&Newproduct)
	if err != nil {
		return nil, err
	}

	// 6. 返回创建的产品
	resp = &product.CreateProductResp{
		Res: true,
	}
	return
}
