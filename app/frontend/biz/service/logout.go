package service

import (
	"context"
	"github.com/hertz-contrib/sessions"

	common "github.com/7qing/gomall/api/frontend/hertz_gen/frontend/common"
	"github.com/cloudwego/hertz/pkg/app"
)

type LogoutService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewLogoutService(Context context.Context, RequestContext *app.RequestContext) *LogoutService {
	return &LogoutService{RequestContext: RequestContext, Context: Context}
}

func (h *LogoutService) Run(req *common.Empty) (resp *common.Empty, err error) {
	// session api
	//推出处理业务逻辑
	session := sessions.Default(h.RequestContext)
	session.Clear()
	err = session.Save()
	if err != nil {
		return nil, err
	}
	return
}
