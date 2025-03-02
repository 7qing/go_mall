package model

import (
	"context"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Product struct {
	gorm.Model
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Picture     string     `json:"picture"`
	Price       float32    `json:"price"`
	Categories  []Category `json:"categories" gorm:"many2many:product_category"` //多对多关系
}

func (Product) TableName() string {
	return "product"
}

// 实现查询
type ProductQuery struct {
	ctx context.Context
	db  *gorm.DB
}

// 创建一个商品
func (p ProductQuery) CreateProduct(product *Product) error {
	return p.db.WithContext(p.ctx).Create(&product).Error
}

// GetByID 查询 api
func (p ProductQuery) GetByID(ID int) (product Product, err error) {
	//err = db.Where("ID = ?", ID).First(&product).Error
	//return product, err
	//调用withcontext方法，方便后续调用链路追踪
	err = p.db.WithContext(p.ctx).Model(&Product{}).First(&product, ID).Error
	return
}

func (p ProductQuery) SearchProducts(search string) (products []Product, err error) {
	err = p.db.WithContext(p.ctx).Model(&Product{}).Find(&products, "name like ? or description like ?",
		"%"+search+"%", "%"+search+"%").Error
	return
}

func NewProductQuery(ctx context.Context, db *gorm.DB) *ProductQuery {
	return &ProductQuery{ctx: ctx, db: db}
}

func (p ProductQuery) DeleteProduct(Name string) error {
	// 删除与产品相关的所有 product_category 记录
	DelProduct := &Product{}
	err := p.FindProductByName(DelProduct, Name)
	if err != nil {
		return err
	}

	// 然后删除产品表中的产品记录
	err = p.db.WithContext(p.ctx).Unscoped().Select(clause.Associations).Delete(DelProduct).Error
	return err
}
func (p ProductQuery) FindProductByName(product *Product, Name string) error {
	err := p.db.WithContext(p.ctx).Model(&Product{}).Where("name = ?", Name).First(product).Error
	return err
}
