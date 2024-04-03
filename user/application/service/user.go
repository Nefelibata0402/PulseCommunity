package service

import (
	"context"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"newsCenter/common/errs"
	"newsCenter/common/jwts"
	"newsCenter/idl/userGrpc"
	"newsCenter/user/application/code"
	"newsCenter/user/application/pkg/snowflake"
	"newsCenter/user/config"
	"newsCenter/user/domain/repository"
	"newsCenter/user/infrastructure/persistence/dal"
	"newsCenter/user/infrastructure/persistence/userData"
	"strconv"
	"time"
)

type UserService struct {
	userGrpc.UnimplementedUserServiceServer
	userRepo repository.UserRepository
}

func New() *UserService {
	return &UserService{
		userRepo: dal.NewUserDao(),
	}
}

func (user *UserService) Register(c context.Context, req *userGrpc.RegisterRequest) (resp *userGrpc.RegisterResponse, err error) {
	//1.校验业务逻辑-用户名是否存在
	UserNameExits, err := user.userRepo.FindUserName(c, req.Username)
	if err != nil {
		zap.L().Error("Register db get error", zap.Error(err))
		return nil, errs.GrpcError(code.DbError)
	}
	if UserNameExits {
		return nil, errs.GrpcError(code.UserNameExist)
	}
	userId, err := snowflake.GetID()
	if err != nil {
		return nil, errs.GrpcError(code.SonyFlakeNotInit)
	}
	userInfo := &userData.UserInfo{
		UserId:   userId,
		Username: req.Username,
		Password: req.Password,
	}
	//2.执行业务-将注册信息插入数据库
	err = user.userRepo.SaveUserInfo(c, userInfo)
	if err != nil {
		zap.L().Error("Register db get error", zap.Error(err))
		return nil, errs.GrpcError(code.DbError)
	}
	//3.返回
	return resp, nil
}
func (user *UserService) Login(c context.Context, req *userGrpc.LoginRequest) (resp *userGrpc.LoginResponse, err error) {
	//1.校验业务逻辑-验证用户名密码是否正确
	userInfo, err := user.userRepo.FindUsernameAndPassword(c, req.Username, req.Password)
	if err != nil {
		zap.L().Error("Login db get error", zap.Error(err))
		return nil, errs.GrpcError(code.DbError)
	} else if userInfo == nil {
		errs.GrpcError(code.UsernameOrPasswordErr)
	}
	//生成token
	userIdString := strconv.FormatInt(userInfo.Id, 10)
	expirationTime := time.Duration(config.UserConfig.JwtConfig.AccessExp*3600*24) * time.Second
	refreshExpirationTime := time.Duration(config.UserConfig.JwtConfig.RefreshExp) * time.Second
	token := jwts.CreateToken(userIdString, expirationTime, config.UserConfig.JwtConfig.AccessSecret, refreshExpirationTime, config.UserConfig.JwtConfig.RefreshSecret)
	tokenList := &userGrpc.TokenRequest{
		AccessToken:    token.AccessToken,
		RefreshToken:   token.RefreshToken,
		TokenType:      "bear",
		AccessTokenExp: token.AccessExp,
	}
	return &userGrpc.LoginResponse{
		UserId: int64(userInfo.UserId),
		Token:  tokenList,
	}, nil
}
func (user *UserService) TokenAuth(c context.Context, req *userGrpc.LoginRequest) (resp *userGrpc.LoginResponse, err error) {
	return nil, status.Errorf(codes.Unimplemented, "method TokenAuth not implemented")
}
