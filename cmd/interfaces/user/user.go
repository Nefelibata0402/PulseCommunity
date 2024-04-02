package user

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"newsCenter/cmd/interfaces/rpc"
	"newsCenter/cmd/model/userModel"
	"newsCenter/common/code"
	"newsCenter/common/errs"
	"newsCenter/idl/userGrpc"
	"time"
)

func Register(c *gin.Context) {
	//1.绑定参数
	result := &code.Result{}
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
		Username:        registerReq.UserName,
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
	c.JSON(200, gin.H{
		"status_code": 200,
	})
}
