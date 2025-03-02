package service

import (
	"context"
	"github.com/7qing/gomall/api/frontend/infra/rpc"
	rpcauth "github.com/7qing/gomall/rpc_gen/kitex_gen/auth"
	"github.com/7qing/gomall/rpc_gen/kitex_gen/user"
	"github.com/hertz-contrib/sessions"

	auth "github.com/7qing/gomall/api/frontend/hertz_gen/frontend/auth"
	common "github.com/7qing/gomall/api/frontend/hertz_gen/frontend/common"
	"github.com/cloudwego/hertz/pkg/app"
)

type RegisterService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewRegisterService(Context context.Context, RequestContext *app.RequestContext) *RegisterService {
	return &RegisterService{RequestContext: RequestContext, Context: Context}
}

func (h *RegisterService) Run(req *auth.RegisterReq) (resp *common.Empty, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	userresp, err := rpc.UserClient.Register(h.Context, &user.RegisterReq{
		Email:           req.Email,
		Password:        req.Password,
		ConfirmPassword: req.ConfirmPassword,
	})

	if err != nil {
		return nil, err
	}

	TokenResp, err := rpc.AuthClient.DeliverTokenByRPC(h.Context, &rpcauth.DeliverTokenReq{
		UserId: userresp.UserId,
	})
	if err != nil {
		return nil, err
	}
	// session api
	//注册处理业务逻辑
	session := sessions.Default(h.RequestContext)
	session.Set("user_id", TokenResp.Token)
	err = session.Save()
	if err != nil {
		return nil, err
	}
	return
}
