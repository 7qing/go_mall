package service

import (
	"context"
	"errors"
	"github.com/7qing/gomall/app/user/biz/dal/mysql"
	"github.com/7qing/gomall/app/user/biz/model"
	user "github.com/7qing/gomall/rpc_gen/kitex_gen/user"
)

type DeleteService struct {
	ctx context.Context
} // NewDeleteService new DeleteService
func NewDeleteService(ctx context.Context) *DeleteService {
	return &DeleteService{ctx: ctx}
}

// Run create note info
func (s *DeleteService) Run(req *user.DeleteReq) (resp *user.DeleteResp, err error) {
	// Finish your business logic.
	err = model.DeleteUser(mysql.DB, uint(req.UserId))
	if err != nil {
		resp = &user.DeleteResp{
			Res: false,
		}
		return resp, errors.New("delete user failed")
	}
	resp = &user.DeleteResp{
		Res: true,
	}
	return resp, nil
}
