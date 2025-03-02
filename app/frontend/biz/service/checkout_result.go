package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/utils"

	common "github.com/7qing/gomall/api/frontend/hertz_gen/frontend/common"
	"github.com/cloudwego/hertz/pkg/app"
)

type CheckoutResultService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCheckoutResultService(Context context.Context, RequestContext *app.RequestContext) *CheckoutResultService {
	return &CheckoutResultService{RequestContext: RequestContext, Context: Context}
}

func (h *CheckoutResultService) Run(req *common.Empty) (resp map[string]any, err error) {
	// todo
	return utils.H{}, nil
}
