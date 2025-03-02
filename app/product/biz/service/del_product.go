package service

import (
	"context"
	"github.com/7qing/gomall/app/product/biz/dal/mysql"
	"github.com/7qing/gomall/app/product/biz/model"
	product "github.com/7qing/gomall/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/klog"
)

type DelProductService struct {
	ctx context.Context
} // NewDelProductService new DelProductService
func NewDelProductService(ctx context.Context) *DelProductService {
	return &DelProductService{ctx: ctx}
}

// Run create note info
func (s *DelProductService) Run(req *product.DelProductReq) (resp *product.DelProductResp, err error) {
	// Finish your business logic.
	ProductQuery := model.NewProductQuery(s.ctx, mysql.DB)
	err = ProductQuery.DeleteProduct(req.Name)
	if err != nil {
		klog.Fatal("delete product err:", err)
	}
	resp = &product.DelProductResp{
		Res: true,
	}
	return
}
