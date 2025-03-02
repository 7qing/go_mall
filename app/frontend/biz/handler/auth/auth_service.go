package auth

import (
	"context"

	"github.com/7qing/gomall/api/frontend/biz/service"
	"github.com/7qing/gomall/api/frontend/biz/utils"
	auth "github.com/7qing/gomall/api/frontend/hertz_gen/frontend/auth"
	common "github.com/7qing/gomall/api/frontend/hertz_gen/frontend/common"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// Register .
// @router /auth/register [POST]
func Register(ctx context.Context, c *app.RequestContext) {
	var err error
	var req auth.RegisterReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	_, err = service.NewRegisterService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	c.Redirect(consts.StatusOK, []byte("/"))

}

// Login .
// @router /auth/login [POST]
func Login(ctx context.Context, c *app.RequestContext) {
	var err error
	var req auth.LoginReq
	var redirect string
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	//resp := &common.Empty{}
	//resp, err = service.NewLoginService(ctx, c).Run(&req)
	redirect, err = service.NewLoginService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	// 常量包中的OK状态consts.StatusOK
	//处理完成后重定向
	c.Redirect(consts.StatusOK, []byte(redirect))

}

// Logout .
// @router /auth/logout [POST]
func Logout(ctx context.Context, c *app.RequestContext) {
	var err error
	var req common.Empty
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	_, err = service.NewLogoutService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	// 常量包中的OK状态consts.StatusOK
	//处理完成后重定向
	c.Redirect(consts.StatusOK, []byte("/"))
}

// Delete .
// @router /auth/delete [POST]
func Delete(ctx context.Context, c *app.RequestContext) {
	var err error
	var req common.Empty
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	_, err = service.NewDeleteService(ctx, c).Run(&req)

	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	c.Redirect(consts.StatusOK, []byte("/"))
}
