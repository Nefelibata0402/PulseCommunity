package tokenVerify

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"newsCenter/cmd/interfaces/user"
	"newsCenter/common/unierr"
	"newsCenter/idl/userGrpc"
	"time"
)

func TokenVerify() func(*gin.Context) {
	return func(c *gin.Context) {
		//1. 从header中获取token
		token := c.GetHeader("Authorization")
		//2. 调用user服务进行token认证
		//做一个超时180秒不响应
		ctx, cancel := context.WithTimeout(context.Background(), 180*time.Second)
		defer cancel()
		response, err := user.UserServiceClient.TokenAuth(ctx, &userGrpc.TokenRequest{Token: token})
		if err != nil {
			zap.L().Error("TokenVerify TokenAuth Fail", zap.Error(err))
			c.JSON(http.StatusOK, gin.H{
				"status_code": unierr.NoLogin.ErrCode,
				"status_msg":  unierr.NoLogin.ErrMsg,
			})
			c.Abort()
			return
		}
		//3. 处理结果，认证通过 将信息放入gin上下文，失败返回未登录
		c.Set("userId", response.UserId)
		c.Set("Ssid", response.Ssid)
		c.Next()
	}
}
