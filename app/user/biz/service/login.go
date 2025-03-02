package service

import (
	"context"
	"errors"
	"github.com/7qing/gomall/app/user/biz/dal/mysql"
	"github.com/7qing/gomall/app/user/biz/model"
	"github.com/7qing/gomall/rpc_gen/kitex_gen/user"

	"golang.org/x/crypto/bcrypt"
)

type LoginService struct {
	ctx context.Context
} // NewLoginService new LoginService
func NewLoginService(ctx context.Context) *LoginService {
	return &LoginService{ctx: ctx}
}

// Run create note info
func (s *LoginService) Run(req *user.LoginReq) (resp *user.LoginResp, err error) {
	// Step 1: Validate input
	// 检查电子邮件和密码是否为空，若为空则返回错误
	if req.Email == "" || req.Password == "" {
		return nil, errors.New("email, password are required")
	}

	// Step 2: Get user by email
	// 根据用户输入的电子邮件查找用户
	newuser, err := model.GetByEmail(mysql.DB, req.Email)
	if err != nil {
		// 如果没有找到对应的电子邮件，则返回错误
		return nil, errors.New("no such email")
	}

	// Step 3: Compare password with stored hashed password
	// 使用bcrypt比较输入的密码与数据库中存储的加密密码
	err = bcrypt.CompareHashAndPassword([]byte(newuser.PasswordHashed), []byte(req.Password))
	if err != nil {
		// 如果密码不匹配，则返回错误
		return nil, errors.New("wrong password")
	}

	// Step 4: Return successful login response
	// 如果密码匹配，返回一个包含用户ID的登录响应
	return &user.LoginResp{UserId: int32(newuser.ID)}, nil
}
