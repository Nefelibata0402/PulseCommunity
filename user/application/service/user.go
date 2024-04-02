package service

import (
	"context"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"newsCenter/common/errs"
	"newsCenter/idl/userGrpc"
	"newsCenter/user/application/code"
	"newsCenter/user/application/pkg/snowflake"
	"newsCenter/user/domain/repository"
	"newsCenter/user/infrastructure/persistence/dal"
	"newsCenter/user/infrastructure/persistence/userData"
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
	UserNameExits, err := user.userRepo.GetUserByUserName(c, req.Username)
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
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (user *UserService) TokenAuth(c context.Context, req *userGrpc.LoginRequest) (resp *userGrpc.LoginResponse, err error) {
	return nil, status.Errorf(codes.Unimplemented, "method TokenAuth not implemented")
}
