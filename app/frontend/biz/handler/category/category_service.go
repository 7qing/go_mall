package category

import (
	"context"

	"github.com/7qing/gomall/api/frontend/biz/service"
	"github.com/7qing/gomall/api/frontend/biz/utils"
	category "github.com/7qing/gomall/api/frontend/hertz_gen/frontend/category"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// Category .
// @router /category/:category [GET]
func Category(ctx context.Context, c *app.RequestContext) {
	var err error
	var req category.CategoryReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewCategoryService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	c.HTML(consts.StatusOK, "category", utils.WarpResponse(ctx, c, resp))
}
