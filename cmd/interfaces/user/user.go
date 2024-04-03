package user

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"newsCenter/cmd/interfaces/rpc"
	"newsCenter/cmd/model/userModel"
	"newsCenter/common/errs"
	"newsCenter/common/returnCode"
	"newsCenter/idl/userGrpc"
	"time"
)

func Register(c *gin.Context) {
	//1.绑定参数
	result := &returnCode.Result{}
	var registerReq userModel.RegisterRequest
	if err := c.ShouldBind(&registerReq); err != nil {
		c.JSON(http.StatusOK, result.Fail(http.StatusBadRequest, "请求参数格式有误"))
		return
	}
	//2.校验参数
	err := registerReq.Verify()
	if err != nil {
		c.JSON(http.StatusOK, result.Fail(http.StatusBadRequest, err.Error()))
		return
	}
	//3.调用grpc服务
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	req := &userGrpc.RegisterRequest{
		Username:        registerReq.Username,
		Password:        registerReq.Password,
		ConfirmPassword: registerReq.ConfirmPassword,
	}
	resp, err := rpc.UserServiceClient.Register(ctx, req)
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		c.JSON(http.StatusOK, result.Fail(code, msg))
		return
	}
	//4.返回结果
	c.JSON(http.StatusOK, result.Success(resp))
}

func Login(c *gin.Context) {
	//1.绑定参数
	result := &returnCode.Result{}
	var LoginReq userModel.LoginRequest
	if err := c.ShouldBind(&LoginReq); err != nil {
		c.JSON(http.StatusOK, result.Fail(http.StatusBadRequest, "请求参数格式有误"))
		return
	}
	//2.校验参数

	//3.调用grpc服务
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	req := &userGrpc.LoginRequest{
		Username: LoginReq.Username,
		Password: LoginReq.Password,
	}
	resp, err := rpc.UserServiceClient.Login(ctx, req)
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		c.JSON(http.StatusOK, result.Fail(code, msg))
		return
	}
	//4.返回结果
	c.JSON(http.StatusOK, result.Success(resp))
}
