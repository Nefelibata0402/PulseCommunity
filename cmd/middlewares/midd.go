package middlewares

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"newsCenter/cmd/interfaces/rpc"
	"newsCenter/common/code"
	"newsCenter/common/errs"
	"newsCenter/idl/userGrpc"
	"time"
)

func TokenVerify() func(*gin.Context) {
	result := &code.Result{}
	return func(c *gin.Context) {
		//1. 从header中获取token
		token := c.GetHeader("Authorization")
		//2. 调用user服务进行token认证
		//做一个超时180秒不响应
		ctx, cancelFunc := context.WithTimeout(context.Background(), 180*time.Second)
		defer cancelFunc()
		response, err := rpc.UserServiceClient.TokenAuth(ctx, &userGrpc.LoginRequest{Token: token})
		if err != nil {
			code, msg := errs.ParseGrpcError(err)
			c.JSON(http.StatusOK, result.Fail(code, msg))
			c.Abort()
			return
		}
		//3. 处理结果，认证通过 将信息放入gin上下文，失败返回未登录
		c.Set("memberId", response.Token)
		c.Next()
	}
}
