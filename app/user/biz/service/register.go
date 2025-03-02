package service

import (
	"context"
	"errors"
	"github.com/7qing/gomall/app/user/biz/dal/mysql"
	"github.com/7qing/gomall/app/user/biz/model"
	user "github.com/7qing/gomall/rpc_gen/kitex_gen/user"
	"golang.org/x/crypto/bcrypt"
)

type RegisterService struct {
	ctx context.Context
} // NewRegisterService new RegisterService
func NewRegisterService(ctx context.Context) *RegisterService {
	return &RegisterService{ctx: ctx}
}

// Run create note info
func (s *RegisterService) Run(req *user.RegisterReq) (resp *user.RegisterResp, err error) {
	// Step 1: Validate input
	// 检查电子邮件是否为空，若为空则返回错误
	if req.Email == "" {
		return nil, errors.New("email is required")
	}

	// 检查密码是否为空，若为空则返回错误
	if req.Password == "" {
		return nil, errors.New("password is required")
	}

	// 检查确认密码是否为空，若为空则返回错误
	if req.ConfirmPassword == "" {
		return nil, errors.New("confirm password is required")
	}

	// 检查密码和确认密码是否一致，若不一致则返回错误
	if req.Password != req.ConfirmPassword {
		return nil, errors.New("password not match")
	}

	// Step 2: Hash the password
	// 使用bcrypt进行密码加密，生成加密后的密码
	passwordHashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		// 如果密码加密过程中发生错误，返回该错误
		return nil, err
	}

	// Step 3: Create a new user object
	// 创建一个新的User对象，并将电子邮件和加密后的密码存储在User结构体中
	newUser := &model.User{
		Email:          req.Email,
		PasswordHashed: string(passwordHashed), // 存储加密后的密码
	}

	// Step 4: Store user in the database
	// 将新创建的用户信息存储到数据库
	model.CreateUser(mysql.DB, newUser)

	// Step 5: Return the user ID in the response
	// 返回一个包含用户ID的注册响应
	return &user.RegisterResp{UserId: int32(newUser.ID)}, nil
}
