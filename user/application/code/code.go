package code

import "newsCenter/common/errs"

var (
	DbError               = errs.NewErrCore(2000, "db错误")
	UserNameExist         = errs.NewErrCore(2001, "用户名已存在")
	SonyFlakeNotInit      = errs.NewErrCore(2002, "雪花算法未初始化")
	UsernameOrPasswordErr = errs.NewErrCore(2003, "用户名或密码错误")
)
