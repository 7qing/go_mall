package service

import (
	"context"
	"github.com/7qing/gomall/app/payment/biz/dal/mysql"
	"github.com/7qing/gomall/app/payment/biz/model"
	payment "github.com/7qing/gomall/rpc_gen/kitex_gen/payment"
	"github.com/cloudwego/kitex/pkg/kerrors"
	creditcard "github.com/durango/go-credit-card"
	"github.com/google/uuid"
	"strconv"
	"time"
)

type ChargeService struct {
	ctx context.Context
} // NewChargeService new ChargeService
func NewChargeService(ctx context.Context) *ChargeService {
	return &ChargeService{ctx: ctx}
}

// Run create note info
func (s *ChargeService) Run(req *payment.ChargeReq) (resp *payment.ChargeResp, err error) {
	// Finish your business logic.
	// Step 1: Construct the credit card object
	// 构造信用卡对象
	card := creditcard.Card{
		Number: req.CreditCard.CreditCardNumber,
		Cvv:    strconv.Itoa(int(req.CreditCard.CreditCardCvv)),
		Month:  strconv.Itoa(int(req.CreditCard.CreditCardExpirationMonth)),
		Year:   strconv.Itoa(int(req.CreditCard.CreditCardExpirationYear)),
	}
	// Step 2: Validate the credit card
	// 验证信用卡信息
	err = card.Validate(true)
	if err != nil {
		return nil, kerrors.NewGRPCBizStatusError(4004001, err.Error())
	}
	// Step 3: Generate a unique transaction ID
	// 生成唯一的交易ID
	transactionId, err := uuid.NewRandom()
	if err != nil {
		return nil, kerrors.NewGRPCBizStatusError(4005001, err.Error())
	}
	// Step 4: Log the payment transaction
	// 将支付日志插入数据库
	err = model.CreatePaymentLog(s.ctx, mysql.DB, &model.PaymentLog{
		UserId:        req.UserId,
		OrderId:       req.OrderId,
		TransactionId: transactionId.String(),
		Amount:        req.Amount,
		PayAt:         time.Now(),
	})
	// Step 5: Handle errors in payment log creation
	// 如果支付日志插入失败，返回错误 5005002
	if err != nil {
		return nil, kerrors.NewGRPCBizStatusError(5005002, err.Error())
	}
	// Step 6: Prepare the response
	// 创建并返回响应对象，包含交易ID
	resp = &payment.ChargeResp{
		TransactionId: transactionId.String(),
	}
	return resp, nil
}
