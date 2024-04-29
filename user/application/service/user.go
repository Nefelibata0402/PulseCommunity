package service

import (
	"context"
	"encoding/json"
	"go.uber.org/zap"
	"newsCenter/common/jwts"
	"newsCenter/common/snowflake"
	"newsCenter/common/unierr"
	"newsCenter/idl/userGrpc"
	"newsCenter/user/domain/entity"
	"newsCenter/user/domain/repository"
	"newsCenter/user/domain/service"
	"newsCenter/user/infrastructure/code"
	"newsCenter/user/infrastructure/config"
	"newsCenter/user/infrastructure/persistence/dao"
	"newsCenter/user/infrastructure/pkg/encrypts"
	"strconv"
	"strings"
	"sync"
	"time"
)

type UserService struct {
	userGrpc.UnimplementedUserServiceServer
	userRepo repository.UserRepository
	cache    repository.Cache
}

func New() *UserService {
	return &UserService{
		userRepo: dao.NewUserDao(),
		cache:    dao.Rc,
	}
}

func (user *UserService) Register(c context.Context, req *userGrpc.RegisterRequest) (resp *userGrpc.RegisterResponse, err error) {
	//1.校验业务逻辑-用户名是否存在
	UserNameExits, err := user.userRepo.FindUserName(c, req.Username)
	if err != nil {
		zap.L().Error("Register 数据库错误", zap.Error(err))
		resp = &userGrpc.RegisterResponse{
			StatusCode: unierr.ErrorInternal.ErrCode,
			StatusMsg:  unierr.ErrorInternal.ErrMsg,
		}
		return resp, nil
	}
	if UserNameExits {
		resp = &userGrpc.RegisterResponse{
			StatusCode: unierr.UserNameExist.ErrCode,
			StatusMsg:  unierr.UserNameExist.ErrMsg,
		}
		return resp, nil
	}
	//雪花算法生成user_id
	userId, err := snowflake.GetID()
	if err != nil {
		zap.L().Error("Register 雪花算法初始化失败", zap.Error(err))
		resp = &userGrpc.RegisterResponse{
			StatusCode: unierr.ErrorInternal.ErrCode,
			StatusMsg:  unierr.ErrorInternal.ErrMsg,
		}
	}
	encryptPassword := encrypts.EncryptPassword(req.Password)
	userInfo := &entity.UserInfo{
		UserId:   userId,
		Username: req.Username,
		Password: encryptPassword,
	}
	//2.执行业务-将注册信息插入数据库
	err = user.userRepo.SaveUserInfo(c, userInfo)
	if err != nil {
		zap.L().Error("Register 数据库错误", zap.Error(err))
		resp = &userGrpc.RegisterResponse{
			StatusCode: unierr.ErrorInternal.ErrCode,
			StatusMsg:  unierr.ErrorInternal.ErrMsg,
		}
		return resp, nil
	}
	//创建token
	token, err := service.CreateToken(userInfo)
	if err != nil {
		zap.L().Error("Register 创建token失败", zap.Error(err))
		resp = &userGrpc.RegisterResponse{
			StatusCode: unierr.ErrorInternal.ErrCode,
			StatusMsg:  unierr.ErrorInternal.ErrMsg,
		}
		return resp, nil
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
		UserId:     userInfo.UserId,
		Token:      tokenList,
	}
	//3.返回
	return resp, nil
}
func (user *UserService) Login(c context.Context, req *userGrpc.LoginRequest) (resp *userGrpc.LoginResponse, err error) {
	//1.校验业务逻辑-验证用户名密码是否正确
	encryptPassword := encrypts.EncryptPassword(req.Password)
	userInfo, err := user.userRepo.FindUsernameAndPassword(c, req.Username, encryptPassword)
	if err != nil {
		zap.L().Error("Login 数据库错误", zap.Error(err))
		resp = &userGrpc.LoginResponse{
			StatusCode: unierr.ErrorInternal.ErrCode,
			StatusMsg:  unierr.ErrorInternal.ErrMsg,
		}
		return resp, nil
	} else if userInfo == nil {
		resp = &userGrpc.LoginResponse{
			StatusCode: unierr.UsernameOrPasswordErr.ErrCode,
			StatusMsg:  unierr.UsernameOrPasswordErr.ErrMsg,
		}
		return resp, nil
	}
	//生成token
	token, err := service.CreateToken(userInfo)
	if err != nil {
		zap.L().Error("Login 创建token失败", zap.Error(err))
		resp = &userGrpc.LoginResponse{
			StatusCode: unierr.ErrorInternal.ErrCode,
			StatusMsg:  unierr.ErrorInternal.ErrMsg,
		}
		return resp, nil
	}
	tokenList := &userGrpc.TokenMessage{
		AccessToken:    token.AccessToken,
		RefreshToken:   token.RefreshToken,
		TokenType:      "bear",
		AccessTokenExp: token.AccessExp,
	}
	// 声明一个用于传递错误信息的通道
	errChan := make(chan error, 1)
	var wg sync.WaitGroup
	wg.Add(1)
	//放入缓存 用户信息
	go func() {
		defer wg.Done()
		marshalUserinfo, err := json.Marshal(userInfo)
		if err != nil {
			errChan <- err
			return
		}
		userIdString := strconv.FormatInt(int64(userInfo.UserId), 10)
		expirationTime := time.Duration(config.UserConfig.JwtConfig.AccessExp*3600*24) * time.Second
		err = user.cache.Put(c, code.User+"::"+userIdString, string(marshalUserinfo), expirationTime)
		if err != nil {
			errChan <- err
			return
		}
	}()
	wg.Wait()
	select {
	case err = <-errChan:
		zap.L().Error("Login 用户信息放入缓存失败", zap.Error(err))
		resp = &userGrpc.LoginResponse{
			StatusCode: unierr.ErrorInternal.ErrCode,
			StatusMsg:  unierr.ErrorInternal.ErrMsg,
		}
		return resp, nil
	default:
		// 没有错误发生，继续执行其他操作
	}
	resp = &userGrpc.LoginResponse{
		StatusCode: unierr.Success.ErrCode,
		StatusMsg:  unierr.Success.ErrMsg,
		UserId:     userInfo.UserId,
		Token:      tokenList,
	}
	return resp, nil
}

func (user *UserService) TokenAuth(c context.Context, req *userGrpc.TokenRequest) (resp *userGrpc.TokenResponse, err error) {
	token := req.Token
	if strings.Contains(token, "bearer ") {
		token = strings.ReplaceAll(token, "bearer ", "")
	}
	parseToken, err := jwts.ParseToken(token, config.UserConfig.JwtConfig.AccessSecret)
	if err != nil {
		zap.L().Error("TokenAuth ParseToken error", zap.Error(err))
		resp = &userGrpc.TokenResponse{
			StatusCode: unierr.ErrorInternal.ErrCode,
			StatusMsg:  unierr.ErrorInternal.ErrMsg,
		}
	}
	userJson, err := user.cache.Get(context.Background(), code.User+"::"+parseToken)
	if err != nil {
		zap.L().Error("TokenAuth cache get user Fail, 用户未登陆 或 被攻击 或 redis崩了 redis崩了结束 防止把数据库也打崩了", zap.Error(err))
		resp = &userGrpc.TokenResponse{
			StatusCode: unierr.NoLogin.ErrCode,
			StatusMsg:  unierr.NoLogin.ErrMsg,
		}
	}
	//正常是数据库过期了
	//Todo 可能是redis崩了 放过去可能会打崩数据库，先结束。
	//上面保证一定存入，否则一直失败。
	//如果打到不同实例可能会多次登陆，保证每次打在同一实例上。可以用一致性哈希等等
	if userJson == "" {
		zap.L().Error("TokenAuth cache get user expire，过期 或 被攻击")
		resp = &userGrpc.TokenResponse{
			StatusCode: unierr.NoLogin.ErrCode,
			StatusMsg:  unierr.NoLogin.ErrMsg,
		}
	}
	userInfo := &entity.UserInfo{}
	err = json.Unmarshal([]byte(userJson), userInfo)
	if err != nil {
		zap.L().Error("TokenAuth Unmarshal userJson Fail", zap.Error(err))
		resp = &userGrpc.TokenResponse{
			StatusCode: unierr.ErrorInternal.ErrCode,
			StatusMsg:  unierr.ErrorInternal.ErrMsg,
		}
	}
	return &userGrpc.TokenResponse{UserId: userInfo.UserId}, nil
}
