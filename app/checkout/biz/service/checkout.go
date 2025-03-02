package service

import (
	"context"
	"github.com/7qing/gomall/app/checkout/infra/rpc"
	checkout "github.com/7qing/gomall/rpc_gen/kitex_gen/checkout"
	"github.com/7qing/gomall/rpc_gen/kitex_gen/order"
	payment "github.com/7qing/gomall/rpc_gen/kitex_gen/payment"
	"strconv"

	"github.com/7qing/gomall/rpc_gen/kitex_gen/cart"
	"github.com/7qing/gomall/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
)

type CheckoutService struct {
	ctx context.Context
} // NewCheckoutService new CheckoutService
func NewCheckoutService(ctx context.Context) *CheckoutService {
	return &CheckoutService{ctx: ctx}
}

// Run create note info
func (s *CheckoutService) Run(req *checkout.CheckoutReq) (resp *checkout.CheckoutResp, err error) {
	// Step 0: Check if the request is valid
	if req == nil || req.Address == nil || req.CreditCard == nil {
		return nil, kerrors.NewGRPCBizStatusError(5005000, "Invalid request: Missing required fields")
	}
	// Step 1: Get the cart details
	// 调用 Cart 服务获取用户购物车中的商品

	cartResp, err := rpc.CartClient.GetCart(s.ctx, &cart.GetCartReq{UserId: req.UserId})
	if err != nil {
		return nil, kerrors.NewGRPCBizStatusError(5005001, err.Error())
	}

	// Step 2: Check if cart is empty
	// 如果购物车为空或者没有商品，返回错误码 5004001
	if cartResp == nil || cartResp.Cart == nil || cartResp.Cart.Items == nil {
		return nil, kerrors.NewGRPCBizStatusError(5004001, "Cart is empty or no items")
	}

	// Step 3: Calculate total price and prepare order items
	// 初始化总价和订单项数组
	var total float32
	var orderItems []*order.OrderItem
	for _, cartItem := range cartResp.Cart.Items {
		// 真实环境中避免在for循环中，循环使用rpc调用
		productresp, resultErr := rpc.ProductcatalogClient.GetProduct(s.ctx, &product.GetProductReq{
			Id: cartItem.ProductId})
		if resultErr != nil {
			continue
		}
		if productresp.Product == nil {
			continue
		}
		total += productresp.Product.Price * float32(cartItem.Quantity)
		orderItems = append(orderItems, &order.OrderItem{
			Item: &cart.CartItem{
				ProductId: productresp.Product.Id,
				Quantity:  cartItem.Quantity,
			},
			Cost: productresp.Product.Price * float32(cartItem.Quantity),
		})
	}
	// Step 4: Place the order
	// 创建订单
	var orderId string
	x, _ := strconv.Atoi(req.Address.ZipCode)
	orderResp, err := rpc.OrdermentClient.PlaceOrder(s.ctx, &order.PlaceOrderReq{
		UserId: req.UserId,
		Address: &order.Address{
			StreetAddress: req.Address.StreetAddress,
			City:          req.Address.City,
			State:         req.Address.State,
			Country:       req.Address.Country,
			ZipCode:       int32(x),
		},
		Email:      req.Email,
		OrderItems: orderItems,
	})
	if err != nil {
		return nil, kerrors.NewGRPCBizStatusError(5004002, err.Error())
	}

	orderId = orderResp.Order.OrderId

	// Step 5: Empty the cart
	// 清空用户购物车
	_, err = rpc.CartClient.EmptyCart(s.ctx, &cart.EmptyCartReq{UserId: req.UserId})
	if err != nil {
		klog.Errorf(err.Error())
	}

	// Step 6: Charge the payment
	// 发起支付请求
	paymmentResp, err := rpc.PaymentClient.Charge(s.ctx, &payment.ChargeReq{
		Amount: total,
		CreditCard: &payment.CreditCardInfo{
			CreditCardNumber:          req.CreditCard.CreditCardNumber,
			CreditCardCvv:             req.CreditCard.CreditCardCvv,
			CreditCardExpirationYear:  req.CreditCard.CreditCardExpirationYear,
			CreditCardExpirationMonth: req.CreditCard.CreditCardExpirationMonth,
		},
		OrderId: orderId,
		UserId:  req.UserId,
	})

	// Step 7: Handle payment failure
	// 如果支付失败，返回错误
	if err != nil {
		return nil, err
	}
	// Step 8: Log the payment response
	// 打印支付响应日志
	klog.Info(paymmentResp)
	// Step 9: Return the response
	// 返回包含订单ID和交易ID的响应
	resp = &checkout.CheckoutResp{
		OrderId:       orderId,
		TransactionId: paymmentResp.TransactionId,
	}
	return
}
