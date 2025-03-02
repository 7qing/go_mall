package service

import (
	"context"
	auth "github.com/7qing/gomall/rpc_gen/kitex_gen/auth"
	"github.com/joho/godotenv"
	"testing"
)

func TestVerifyTokenByRPC_Run(t *testing.T) {
	godotenv.Load("../../.env")
	ctx := context.Background()
	s := NewVerifyTokenByRPCService(ctx)
	// init req and assert value

	validRefreshToken := generateValidRefreshToken(t)

	req := &auth.VerifyTokenReq{
		Token: validRefreshToken,
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	// Step 8: Assert the role is "bob" in the response
	if resp == nil {
		t.Errorf("Expected a valid response, but got nil")
		return
	}

	// Assuming the role is inside the response as resp.Role (you may need to adjust based on actual structure)
	if resp.Role != "bob" {
		t.Errorf("Expected role to be 'bob', but got: %v", resp.Role)
	}
}
