package service

import (
	"context"
	"errors"
	auth "github.com/7qing/gomall/rpc_gen/kitex_gen/auth"
	"github.com/dgrijalva/jwt-go"
	"os"
)

type VerifyTokenByRPCService struct {
	ctx context.Context
} // NewVerifyTokenByRPCService new VerifyTokenByRPCService
func NewVerifyTokenByRPCService(ctx context.Context) *VerifyTokenByRPCService {
	return &VerifyTokenByRPCService{ctx: ctx}
}

// Run create note info
func (s *VerifyTokenByRPCService) Run(req *auth.VerifyTokenReq) (resp *auth.VerifyResp, err error) {
	// Finish your business logic.
	token, err := jwt.Parse(req.Token, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})

	if err != nil {
		return nil, err
	}

	// 验证Token的签名方法是否有效
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, errors.New("无效的签名方法")
	}

	// 返回Token中的声明部分
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return &auth.VerifyResp{
			Res:    true,
			UserId: int32(claims["userId"].(float64)),
			Role:   claims["role"].(string),
		}, nil
	}
	return nil, errors.New("无效的Token")
}
