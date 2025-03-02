package service

import (
	"context"
	"github.com/7qing/gomall/api/frontend/infra/rpc"
	"github.com/7qing/gomall/rpc_gen/kitex_gen/user"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/hertz-contrib/sessions"

	auth "github.com/7qing/gomall/api/frontend/hertz_gen/frontend/auth"
	rpcauth "github.com/7qing/gomall/rpc_gen/kitex_gen/auth"
	"github.com/cloudwego/hertz/pkg/app"
)

type LoginService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewLoginService(Context context.Context, RequestContext *app.RequestContext) *LoginService {
	return &LoginService{RequestContext: RequestContext, Context: Context}
}

func (h *LoginService) Run(req *auth.LoginReq) (redirect string, err error) {
	// TODO user svc api
	// 我们所调用的RPC微服务的user的api
	resp, err := rpc.UserClient.Login(h.Context, &user.LoginReq{
		Email:    req.Email,
		Password: req.Password,
	})

	if err != nil {
		return "", err
	}

	klog.Info("login success1")
	// session api
	//登录处理业务逻辑
	TokenResp, err := rpc.AuthClient.DeliverTokenByRPC(h.Context, &rpcauth.DeliverTokenReq{
		UserId: resp.UserId,
	})

	if err != nil {
		return "", err
	}
	klog.Info("login success2")
	session := sessions.Default(h.RequestContext)
	session.Set("user_id", TokenResp.Token)
	err = session.Save()
	if err != nil {
		//hlog.CtxErrorf(h.Context, "Failed to save session: %v", err)
		return "/", err
	}
	redirect = "/"
	if req.Next != "" {
		redirect = req.Next
	}
	return redirect, nil
}
