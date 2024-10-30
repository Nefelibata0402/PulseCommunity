package dao

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"pulseCommunity/user/domain/entity"
	"pulseCommunity/user/infrastructure/persistence/database/gorms"
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
	err := u.conn.Session(c).Model(&entity.UserInfo{}).Where("username", userName).Count(&count).Error
	return count > 0, err
}

func (u *UserDao) SaveUserInfo(c context.Context, userInfo *entity.UserInfo) error {
	return u.conn.Session(c).Create(userInfo).Error
}

func (u *UserDao) FindUsernameAndPassword(c context.Context, userName, password string) (userInfo *entity.UserInfo, err error) {
	err = u.conn.Session(c).Where("username = ? and password = ? ", userName, password).First(&userInfo).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return userInfo, err
}

func (u *UserDao) GetUserInfo(c context.Context, id int64) (userInfo *entity.UserInfo, err error) {
	err = u.conn.Session(c).Where("id = ?", id).First(&userInfo).Error
	if err != nil {
		return nil, err
	}
	return userInfo, nil
}
