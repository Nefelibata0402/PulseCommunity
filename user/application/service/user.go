package service

import (
	"context"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"newsCenter/common/snowflake"
	"newsCenter/common/unierr"
	"newsCenter/idl/userGrpc"
	"newsCenter/user/domain/entity"
	"newsCenter/user/domain/service"
	"newsCenter/user/infrastructure/pkg/encrypts"
)

type UserService struct {
	userGrpc.UnimplementedUserServiceServer
}

func New() *UserService {
	return &UserService{}
}

func (user *UserService) Register(c context.Context, req *userGrpc.RegisterRequest) (resp *userGrpc.RegisterResponse, err error) {
	//1.校验业务逻辑-用户名是否存在及已存在返回的情况
	userService := service.New()
	resp, err = userService.CheckUserNameExits(c, req.Username)
	if err != nil {
		return resp, err
	}
	if resp != nil {
		return resp, nil
	}
	//2.雪花算法生成user_id
	userId, err := snowflake.GetID()
	if err != nil {
		zap.L().Error("Register snowflake.GetID Fail", zap.Error(err))
		return resp, err
	}
	//3.密码加密
	encryptPassword := encrypts.EncryptPassword(req.Password)
	userInfo := &entity.UserInfo{
		Id:       userId,
		Username: req.Username,
		Password: encryptPassword,
	}
	//4.执行业务-将注册信息插入数据库
	err = userService.SaveUserInfo(c, userInfo)
	if err != nil {
		return resp, err
	}
	//5.创建token
	ssid := uuid.New().String()
	token, err := service.CreateToken(userInfo, ssid)
	if err != nil {
		zap.L().Error("Register 创建token失败", zap.Error(err))
		return resp, err
	}
	tokenList := &userGrpc.TokenMessage{
		AccessToken:    token.AccessToken,
		RefreshToken:   token.RefreshToken,
		TokenType:      "bear",
		AccessTokenExp: token.AccessExp,
	}
	resp = &userGrpc.RegisterResponse{
		StatusCode: unierr.Success.ErrCode,
		StatusMsg:  unierr.Success.ErrMsg,
		UserId:     userInfo.Id,
		Token:      tokenList,
	}
	//6.返回
	return resp, nil
}
func (user *UserService) Login(c context.Context, req *userGrpc.LoginRequest) (resp *userGrpc.LoginResponse, err error) {
	//1.校验业务逻辑-验证用户名密码是否正确
	encryptPassword := encrypts.EncryptPassword(req.Password)
	userService := service.New()
	resp, userInfo, err := userService.CheckUsernameAndPassword(c, req.Username, encryptPassword)
	if err != nil {
		zap.L().Error("Login CheckUsernameAndPassword Fail", zap.Error(err))
		return resp, err
	}
	if resp != nil {
		return resp, nil
	}
	//生成token ssid
	ssid := uuid.New().String()
	token, err := service.CreateToken(userInfo, ssid)
	if err != nil {
		zap.L().Error("Login 创建token失败", zap.Error(err))
		return resp, err
	}
	tokenList := &userGrpc.TokenMessage{
		AccessToken:    token.AccessToken,
		RefreshToken:   token.RefreshToken,
		TokenType:      "bear",
		AccessTokenExp: token.AccessExp,
	}
	//用户信息放入缓存中
	err = userService.CacheUserInfo(c, userInfo)
	if err != nil {
		return resp, err
	}
	resp = &userGrpc.LoginResponse{
		StatusCode: unierr.Success.ErrCode,
		StatusMsg:  unierr.Success.ErrMsg,
		UserId:     userInfo.Id,
		Token:      tokenList,
		Ssid:       ssid,
	}
	return resp, nil
}

func (user *UserService) TokenAuth(c context.Context, req *userGrpc.TokenRequest) (resp *userGrpc.TokenResponse, err error) {
	//解析token 获得userId
	userService := service.New()
	parseToken, ssid, err := userService.ParseToken(req.Token)
	if err != nil {
		return resp, err
	}
	//验证ssid
	err = userService.CheckSsid(c, ssid)
	if err != nil {
		return nil, err
	}
	//从缓存中获取用户信息
	resp, userInfo, err := userService.GetCacheUserInfo(c, parseToken)
	if err != nil {
		zap.L().Error("TokenAuth GetCacheUserInfo Fail", zap.Error(err))
		return resp, err
	}
	if resp != nil {
		return resp, nil
	}
	return &userGrpc.TokenResponse{UserId: userInfo.Id, Ssid: ssid}, nil
}

func (user *UserService) GetUserinfo(c context.Context, req *userGrpc.UserInfoRequest) (resp *userGrpc.UserInfoResponse, err error) {
	userService := service.New()
	userinfo, err := userService.GetUserInfo(c, int64(req.UserId))
	if err != nil {
		zap.L().Error("GetUserinfo GetUserInfo Fail", zap.Error(err))
		return nil, err
	}
	resp = &userGrpc.UserInfoResponse{
		StatusCode: unierr.Success.ErrCode,
		StatusMsg:  unierr.Success.ErrMsg,
		User: &userGrpc.User{
			Id:   int64(userinfo.Id),
			Name: userinfo.Username,
		},
	}
	return resp, nil
}

func (user *UserService) LogoutJWT(c context.Context, req *userGrpc.LogoutJWTRequest) (resp *userGrpc.LogoutJWTResponse, err error) {
	userService := service.New()
	err = userService.ClearToken(c, req.Ssid)
	if err != nil {
		zap.L().Error("LogoutJWT ClearToken Fail", zap.Error(err))
		return nil, err
	}
	resp = &userGrpc.LogoutJWTResponse{
		StatusCode: unierr.Success.ErrCode,
		StatusMsg:  unierr.Success.ErrMsg,
	}
	return resp, nil
}
