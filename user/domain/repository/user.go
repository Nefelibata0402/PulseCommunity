package repository

import (
	"context"
	"pulseCommunity/user/domain/entity"
)

type UserRepository interface {
	// FindUserName 用户名是否存在
	FindUserName(c context.Context, userName string) (bool, error)
	// SaveUserInfo SaveUserName 将注册信息插入数据库
	SaveUserInfo(c context.Context, userInfo *entity.UserInfo) error
	FindUsernameAndPassword(c context.Context, userName, password string) (userInfo *entity.UserInfo, err error)
	GetUserInfo(c context.Context, id int64) (userInfo *entity.UserInfo, err error)
}
