package service

import (
	"context"
	auth "github.com/7qing/gomall/rpc_gen/kitex_gen/auth"
	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"os"
	"testing"
	"time"
)

func TestRenewTokenByRPC_Run(t *testing.T) {
	godotenv.Load("../../.env")
	ctx := context.Background()
	s := NewRenewTokenByRPCService(ctx)
	// init req and assert value

	validRefreshToken := generateValidRefreshToken(t) // Helper function to generate a valid refresh token

	req := &auth.RenewTokenReq{
		RefreshToken: validRefreshToken,
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

}

func generateValidRefreshToken(t *testing.T) string {
	// Create a valid JWT token using a known secret and claims
	claims := jwt.MapClaims{
		"userId": 123,   // Replace with your userId
		"role":   "bob", // Replace with your role
		"exp":    time.Now().Add(time.Hour * 24 * 7).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		t.Fatalf("Failed to generate token: %v", err)
	}
	return tokenString
}
