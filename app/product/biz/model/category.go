package model

import (
	"context"
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Product     []Product `json:"product" gorm:"many2many:product_category"`
}

func (Category) TableName() string {
	return "category"
}

// 实现查询
type CategoryQuery struct {
	ctx context.Context
	db  *gorm.DB
}

// 创建一个分类

func (c CategoryQuery) GetProductsbyCategory(category string) (categories []Category, err error) {
	//preload 预加载
	// 真实电商项目，这里查询很复杂
	err = c.db.WithContext(c.ctx).Model(&Category{Name: category}).Preload("Product").Find(&categories).Error
	return
}

func (c CategoryQuery) FindCategoryByName(name string, Category *Category) (err error) {
	err = c.db.WithContext(c.ctx).Where("name = ?", name).First(Category).Error
	return err
}
func (c CategoryQuery) CreateCategory(name string, Descriptor string) (err error) {
	category := Category{
		Name:        name,
		Description: Descriptor,
	}
	err = c.db.WithContext(c.ctx).Create(&category).Error
	return err
}
func NewCategoryQuery(ctx context.Context, db *gorm.DB) *CategoryQuery {
	return &CategoryQuery{
		ctx: ctx,
		db:  db,
	}
}
