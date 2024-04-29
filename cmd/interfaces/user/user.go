package user

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"newsCenter/cmd/model/userModel"
	"newsCenter/common/unierr"
	"newsCenter/idl/userGrpc"
	"time"
)

func Register(c *gin.Context) {
	//1.绑定参数
	var registerReq userModel.RegisterRequest
	if err := c.ShouldBind(&registerReq); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status_code": unierr.ErrorParams.ErrCode,
			"status_msg":  unierr.ErrorParams.ErrMsg,
		})
		return
	}
	//2.校验参数
	err := registerReq.Verify()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status_code": unierr.DifferentPassword.ErrCode,
			"status_msg":  unierr.DifferentPassword.ErrMsg,
		})
		return
	}
	//3.调用grpc服务
	ctx, cancel := context.WithTimeout(context.Background(), 180*time.Second)
	defer cancel()
	req := &userGrpc.RegisterRequest{
		Username:        registerReq.Username,
		Password:        registerReq.Password,
		ConfirmPassword: registerReq.ConfirmPassword,
	}
	resp, err := UserServiceClient.Register(ctx, req)
	if err != nil {
		c.JSON(http.StatusOK, unierr.ErrorInternal)
		return
	}
	//4.返回结果
	c.JSON(http.StatusOK, gin.H{
		"status_code": resp.StatusCode,
		"status_msg":  resp.StatusMsg,
		"user_id":     resp.UserId,
		"token":       resp.Token,
	})
}

func Login(c *gin.Context) {
	//1.绑定参数
	var LoginReq userModel.LoginRequest
	if err := c.ShouldBind(&LoginReq); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status_code": unierr.ErrorParams.ErrCode,
			"status_msg":  unierr.ErrorParams.ErrMsg,
		})
	}
	//2.校验参数
	//3.调用grpc服务
	ctx, cancel := context.WithTimeout(context.Background(), 180*time.Second)
	defer cancel()
	req := &userGrpc.LoginRequest{
		Username: LoginReq.Username,
		Password: LoginReq.Password,
	}
	resp, err := UserServiceClient.Login(ctx, req)
	if err != nil {
		c.JSON(http.StatusOK, unierr.ErrorInternal)
		return
	}
	//4.返回结果
	c.JSON(http.StatusOK, gin.H{
		"status_code": resp.StatusCode,
		"status_msg":  resp.StatusMsg,
		"user_id":     resp.UserId,
		"token":       resp.Token,
	})
}
