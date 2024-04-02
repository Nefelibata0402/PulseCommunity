package repository

import (
	"context"
	"newsCenter/user/infrastructure/persistence/userData"
)

type UserRepository interface {
	// GetUserByUserName 用户名是否存在
	GetUserByUserName(c context.Context, userName string) (bool, error)
	// SaveUserInfo SaveUserName 将注册信息插入数据库
	SaveUserInfo(c context.Context, userInfo *userData.UserInfo) error
}
