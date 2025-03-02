package service

import (
	"context"
	"errors"
	auth "github.com/7qing/gomall/rpc_gen/kitex_gen/auth"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

type RenewTokenByRPCService struct {
	ctx context.Context
} // NewRenewTokenByRPCService new RenewTokenByRPCService
func NewRenewTokenByRPCService(ctx context.Context) *RenewTokenByRPCService {
	return &RenewTokenByRPCService{ctx: ctx}
}

// Run create note info
func (s *RenewTokenByRPCService) Run(req *auth.RenewTokenReq) (resp *auth.DeliveryResp, err error) {
	// Finish your business logic.
	token, err := jwt.Parse(req.RefreshToken, func(token *jwt.Token) (interface{}, error) {
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
		token := jwt.New(jwt.SigningMethodHS256)
		claims2 := token.Claims.(jwt.MapClaims)
		claims2["userId"] = claims["userId"].(float64)
		claims2["role"] = claims["role"].(string)

		claims2["exp"] = time.Now().Add(time.Hour * 24).Unix() // 设置Token的过期时间
		tokenString, _ := token.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
		resp = &auth.DeliveryResp{
			Token: tokenString,
		}
		return resp, nil
	}
	return nil, errors.New("无效的Token")
}
