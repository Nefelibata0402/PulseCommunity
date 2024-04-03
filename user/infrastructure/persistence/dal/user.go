package dal

import (
	"context"
	"errors"
	"gorm.io/gorm"
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

func (u *UserDao) FindUserName(c context.Context, userName string) (bool, error) {
	var count int64
	err := u.conn.Session(c).Model(&userData.UserInfo{}).Where("username", userName).Count(&count).Error
	return count > 0, err
}

func (u *UserDao) SaveUserInfo(c context.Context, userInfo *userData.UserInfo) error {
	return u.conn.Session(c).Create(userInfo).Error
}

func (u *UserDao) FindUsernameAndPassword(c context.Context, userName, password string) (userInfo *userData.UserInfo, err error) {
	err = u.conn.Session(c).Where("username = ? and password = ? ", userName, password).First(&userInfo).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return userInfo, err
}
