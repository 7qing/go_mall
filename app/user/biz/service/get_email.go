package service

import (
	"context"
	"github.com/7qing/gomall/app/user/biz/dal/mysql"
	"github.com/7qing/gomall/app/user/biz/model"
	user "github.com/7qing/gomall/rpc_gen/kitex_gen/user"
)

type GetEmailService struct {
	ctx context.Context
} // NewGetEmailService new GetEmailService
func NewGetEmailService(ctx context.Context) *GetEmailService {
	return &GetEmailService{ctx: ctx}
}

// Run create note info
func (s *GetEmailService) Run(req *user.DeleteReq) (resp *user.GetEmailResp, err error) {
	// Finish your business logic.
	Userresp, _ := model.GetByID(mysql.DB, uint(req.UserId))
	resp = &user.GetEmailResp{
		Email: Userresp.Email,
	}
	return resp, nil
}
