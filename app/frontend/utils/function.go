package utils

import (
	"context"
	"github.com/7qing/gomall/api/frontend/infra/rpc"
	"github.com/7qing/gomall/rpc_gen/kitex_gen/auth"
	"github.com/cloudwego/kitex/pkg/klog"
)

func GetUserIdFromCtx(ctx context.Context) int32 {
	token := ctx.Value(UserIdKey)
	if token == nil {
		return 0
	}
	tokenStr, ok := token.(string)
	if !ok {
		// 如果 token 不是字符串类型，返回默认值或处理错误
		klog.Fatal("Token is not a string")
		return 0
	}
	tokenResp, err := rpc.AuthClient.VerifyTokenByRPC(ctx, &auth.VerifyTokenReq{
		Token: tokenStr,
	})
	if err != nil {
		klog.Fatal("AuthClient.VerifyTokenByRPC err:", err)
	}
	userid := tokenResp.UserId
	return userid
}
