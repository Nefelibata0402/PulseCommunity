package repository

import (
	"context"
	"newsCenter/user/infrastructure/persistence/userData"
)

type UserRepository interface {
	// FindUserName 用户名是否存在
	FindUserName(c context.Context, userName string) (bool, error)
	// SaveUserInfo SaveUserName 将注册信息插入数据库
	SaveUserInfo(c context.Context, userInfo *userData.UserInfo) error
	FindUsernameAndPassword(c context.Context, userName, password string) (userInfo *userData.UserInfo, err error)
}
