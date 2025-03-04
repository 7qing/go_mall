package service

import (
	"context"
	"github.com/7qing/gomall/app/product/biz/dal/mysql"
	"github.com/7qing/gomall/app/product/biz/model"
	product "github.com/7qing/gomall/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/kerrors"
)

type GetProductService struct {
	ctx context.Context
} // NewGetProductService new GetProductService
func NewGetProductService(ctx context.Context) *GetProductService {
	return &GetProductService{ctx: ctx}
}

// Run create note info
func (s *GetProductService) Run(req *product.GetProductReq) (resp *product.GetProductResp, err error) {

	// Finish your business logic.
	if req.Id == 0 {
		return nil, kerrors.NewGRPCBizStatusError(2004001, "product id is required")
	}
	productQuery := model.NewProductQuery(s.ctx, mysql.DB)

	newproduct, err := productQuery.GetByID(int(req.Id))

	if err != nil {
		return nil, err
	}

	resp = &product.GetProductResp{
		Product: &product.Product{
			Id:          uint32(newproduct.ID),
			Name:        newproduct.Name,
			Description: newproduct.Description,
			Picture:     newproduct.Picture,
			Price:       newproduct.Price,
		},
	}
	return resp, nil
}
