package service

import (
	"context"
	"fmt"
	"github.com/7qing/gomall/api/frontend/hertz_gen/frontend/checkout"
	"github.com/7qing/gomall/api/frontend/infra/rpc"
	frontendUtils "github.com/7qing/gomall/api/frontend/utils"
	rpccheckout "github.com/7qing/gomall/rpc_gen/kitex_gen/checkout"
	"github.com/7qing/gomall/rpc_gen/kitex_gen/payment"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/kitex/pkg/klog"

	"github.com/cloudwego/hertz/pkg/app"
)

type CheckoutWaitingService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCheckoutWaitingService(Context context.Context, RequestContext *app.RequestContext) *CheckoutWaitingService {
	return &CheckoutWaitingService{RequestContext: RequestContext, Context: Context}
}

func (h *CheckoutWaitingService) Run(req *checkout.CheckoutReq) (resp map[string]any, err error) {
	// 检查请求参数是否完整
	if req == nil || req.Firstname == "" || req.Lastname == "" || req.Email == "" || req.Street == "" ||
		req.City == "" || req.Province == "" || req.Country == "" || req.Zipcode == "" || req.CardNum == "" ||
		req.Cvv == 0 || req.ExpirationYear == 0 || req.ExpirationMonth == 0 {
		return nil, fmt.Errorf("Invalid request, some fields are missing")
	}

	userId := frontendUtils.GetUserIdFromCtx(h.Context)

	// 增加日志记录请求信息，方便调试
	klog.Infof("Calling Checkout with UserId: %v, Email: %s", userId, req.Email)

	// 调用 CheckoutClient
	_, err = rpc.CheckoutClient.Checkout(h.Context, &rpccheckout.CheckoutReq{
		UserId:    uint32(userId),
		Firstname: req.Firstname,
		Lastname:  req.Lastname,
		Email:     req.Email,
		Address: &rpccheckout.Address{
			StreetAddress: req.Street,
			City:          req.City,
			State:         req.Province,
			Country:       req.Country,
			ZipCode:       req.Zipcode,
		},
		CreditCard: &payment.CreditCardInfo{
			CreditCardNumber:          req.CardNum,
			CreditCardCvv:             req.Cvv,
			CreditCardExpirationYear:  req.ExpirationYear,
			CreditCardExpirationMonth: req.ExpirationMonth,
		},
	})

	// 错误处理
	if err != nil {
		klog.Errorf("Error during Checkout: %v", err)
		return nil, err
	}
	return utils.H{
		"title":    "waiting",
		"redirect": "/checkout/result",
	}, nil
}
