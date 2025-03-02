package service

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"

	auth "github.com/7qing/gomall/rpc_gen/kitex_gen/auth"
	"github.com/joho/godotenv"
	"testing"
)

func TestDeliverTokenByRPC_Run(t *testing.T) {
	godotenv.Load("../../.env")
	ctx := context.Background()
	s := NewDeliverTokenByRPCService(ctx)
	// init req and assert value

	req := &auth.DeliverTokenReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test
	// 检查是否有错误返回
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	// 检查Token是否生成
	if resp.Token == "" {
		t.Errorf("Expected a token to be generated, but got an empty token")
	}

	// 解析Token并验证字段
	token, _ := jwt.Parse(resp.Token, func(token *jwt.Token) (interface{}, error) {
		// 获取密钥来验证Token签名
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})

	// 验证Token Claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// 验证userId是否正确
		if claims["userId"] != float64(req.UserId) {
			t.Errorf("Expected userId %v, but got %v", req.UserId, claims["userId"])
		}

		// 验证角色是否正确
		expectedRole := "bob"
		if req.UserId == 2 {
			expectedRole = "alice"
		}

		if claims["role"] != expectedRole {
			t.Errorf("Expected role %v, but got %v", expectedRole, claims["role"])
		}

		// 验证Token过期时间是否设置正确
		exp := int64(claims["exp"].(float64))
		expectedExp := time.Now().Add(time.Hour * 24 * 7).Unix()
		if exp != expectedExp {
			t.Errorf("Expected expiration time %v, but got %v", expectedExp, exp)
		}
	} else {
		t.Errorf("Failed to parse or validate the token")
	}
}
