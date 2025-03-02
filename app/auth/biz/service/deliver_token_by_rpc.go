package service

import (
	"context"
	auth "github.com/7qing/gomall/rpc_gen/kitex_gen/auth"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

type DeliverTokenByRPCService struct {
	ctx context.Context
} // NewDeliverTokenByRPCService new DeliverTokenByRPCService
func NewDeliverTokenByRPCService(ctx context.Context) *DeliverTokenByRPCService {
	return &DeliverTokenByRPCService{ctx: ctx}
}

// Run create note info
func (s *DeliverTokenByRPCService) Run(req *auth.DeliverTokenReq) (resp *auth.DeliveryResp, err error) {
	// 创建一个新的JWT Token实例，使用HS256签名方法
	token := jwt.New(jwt.SigningMethodHS256)

	// 获取token中的Claims部分（载荷）
	claims := token.Claims.(jwt.MapClaims)

	// 设置Token中的userId字段
	claims["userId"] = req.UserId

	// 根据不同的用户ID设置角色（role）
	// 如果用户ID为2，则角色为alice；否则角色为bob
	if req.UserId == 2 {
		claims["role"] = "alice"
	} else {
		claims["role"] = "bob"
	}

	// 设置Token的过期时间为7天后
	claims["exp"] = time.Now().Add(time.Hour * 24 * 7).Unix()

	// 使用环境变量中的密钥（ACCESS_SECRET）来签署Token，并生成最终的token字符串
	tokenString, _ := token.SignedString([]byte(os.Getenv("ACCESS_SECRET")))

	// 构建响应对象，将生成的Token字符串传递给客户端
	resp = &auth.DeliveryResp{
		Token: tokenString,
	}

	// 返回响应对象
	return resp, nil
}
