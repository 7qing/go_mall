package middleware

import (
	"context"
	"github.com/7qing/gomall/api/frontend/infra/rpc"
	"github.com/7qing/gomall/rpc_gen/kitex_gen/auth"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/hertz-contrib/casbin"
	"github.com/hertz-contrib/sessions"
)

func Register(h *server.Hertz) {
	h.Use(GlobalAuth())
}

var Casbinauth *casbin.Middleware

func CasbinRegister(h *server.Hertz) {
	// 使用 session 存储用户信息.
	var err error

	Casbinauth, err = casbin.NewCasbinMiddleware("config/model.conf", "config/policy.csv", parseToken)
	if err != nil {
		klog.Fatal(err)
	}
}

// subjectFromSession 从 session 中获取访问实体.
func parseToken(ctx context.Context, c *app.RequestContext) string {
	// 获取访问实体
	session := sessions.Default(c)
	if tokenString, ok := session.Get("user_id").(string); !ok {
		return ""
	} else {
		Verifyresp, err := rpc.AuthClient.VerifyTokenByRPC(ctx, &auth.VerifyTokenReq{
			Token: tokenString,
		})
		if err != nil {
			klog.Fatal(err)
		}
		if Verifyresp.Res == false {
			klog.Fatal("解析错误")
		}
		return Verifyresp.Role
	}
}
