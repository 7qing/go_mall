package service

import (
	"context"
	"github.com/7qing/gomall/api/frontend/infra/rpc"
	frontendUtils "github.com/7qing/gomall/api/frontend/utils"
	"github.com/7qing/gomall/rpc_gen/kitex_gen/user"
	"github.com/hertz-contrib/sessions"

	common "github.com/7qing/gomall/api/frontend/hertz_gen/frontend/common"
	"github.com/cloudwego/hertz/pkg/app"
)

type DeleteService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewDeleteService(Context context.Context, RequestContext *app.RequestContext) *DeleteService {
	return &DeleteService{RequestContext: RequestContext, Context: Context}
}

func (h *DeleteService) Run(req *common.Empty) (resp *common.Empty, err error) {
	_, err = rpc.UserClient.Delete(h.Context, &user.DeleteReq{
		UserId: frontendUtils.GetUserIdFromCtx(h.Context),
	})
	session := sessions.Default(h.RequestContext)
	session.Clear()
	err = session.Save()
	if err != nil {
		return nil, err
	}
	
	return
}
