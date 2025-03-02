package main

import (
	"context"
	"github.com/7qing/gomall/app/payment/biz/service"
	"github.com/7qing/gomall/app/payment/conf"
	payment "github.com/7qing/gomall/rpc_gen/kitex_gen/payment"
	"github.com/joho/godotenv"
	"testing"
)

func TestCharge_Run(t *testing.T) {
	godotenv.Load(".env")

	conf.InitConf()
	ctx := context.Background()
	s := service.NewChargeService(ctx)
	// init req and assert value

	req := &payment.ChargeReq{
		Amount: 2,
		CreditCard: &payment.CreditCardInfo{
			CreditCardNumber:          "1",
			CreditCardCvv:             01,
			CreditCardExpirationYear:  2030,
			CreditCardExpirationMonth: 11,
		},
		OrderId: "dsanldsajlk;dj",
		UserId:  2,
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
