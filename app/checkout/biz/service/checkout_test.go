package service

import (
	"context"
	"github.com/7qing/gomall/app/checkout/infra/rpc"

	"github.com/7qing/gomall/app/checkout/conf"
	checkout "github.com/7qing/gomall/rpc_gen/kitex_gen/checkout"
	"github.com/7qing/gomall/rpc_gen/kitex_gen/payment"
	"github.com/joho/godotenv"
	"testing"
)

func TestCheckout_Run(t *testing.T) {
	godotenv.Load(".env")

	conf.InitConf()
	rpc.Init()
	ctx := context.Background()
	s := NewCheckoutService(ctx)
	// init req and assert value

	req := &checkout.CheckoutReq{
		UserId:    2,
		Firstname: "Y",
		Lastname:  "X",
		Email:     "33@33.com",
		Address: &checkout.Address{
			StreetAddress: "1",
			City:          "1",
			State:         "1",
			Country:       "1",
			ZipCode:       "1",
		},
		CreditCard: &payment.CreditCardInfo{
			CreditCardNumber:          "1",
			CreditCardCvv:             1,
			CreditCardExpirationYear:  2030,
			CreditCardExpirationMonth: 12,
		},
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
