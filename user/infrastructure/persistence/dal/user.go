package dal

import (
	"context"
	"newsCenter/user/infrastructure/persistence/database/gorms"
	"newsCenter/user/infrastructure/persistence/userData"
)

type UserDao struct {
	conn *gorms.GormConn
}

func NewUserDao() *UserDao {
	return &UserDao{
		conn: gorms.New(),
	}
}

func (u *UserDao) GetUserByUserName(c context.Context, userName string) (bool, error) {
	var count int64
	err := u.conn.Session(c).Model(&userData.UserInfo{}).Where("username", userName).Count(&count).Error
	return count > 0, err
}

func (u *UserDao) SaveUserInfo(c context.Context, userInfo *userData.UserInfo) error {
	return u.conn.Session(c).Create(userInfo).Error
}
