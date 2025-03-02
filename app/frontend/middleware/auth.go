package middleware

import (
	"context"
	"github.com/7qing/gomall/api/frontend/utils"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/sessions"
)

// 方便业务逻辑，获取身份认证相关内容
func GlobalAuth() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		session := sessions.Default(c)
		session.Get("user_id")
		ctx = context.WithValue(ctx, utils.UserIdKey, session.Get("user_id"))
		c.Next(ctx)
	}
}

// 个别用户要使用（需要登录）
func Auth() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		session := sessions.Default(c)
		user_id := session.Get("user_id")
		if user_id == nil {
			c.Redirect(302, []byte("/sign-in?next="+c.FullPath()))
			c.Abort()
			return
		}
		ctx = context.WithValue(ctx, utils.UserIdKey, session.Get("user_id"))
		c.Next(ctx)
	}
}
