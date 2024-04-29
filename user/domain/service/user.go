package service

import (
	"newsCenter/common/jwts"
	"newsCenter/user/domain/entity"
	"newsCenter/user/infrastructure/config"
	"strconv"
	"time"
)

func CreateToken(userInfo *entity.UserInfo) (token *jwts.JwtToken, err error) {
	userIdString := strconv.FormatInt(int64(userInfo.UserId), 10)
	expirationTime := time.Duration(config.UserConfig.JwtConfig.AccessExp*3600*24) * time.Second
	refreshExpirationTime := time.Duration(config.UserConfig.JwtConfig.RefreshExp*3600*24) * time.Second
	token, err = jwts.CreateToken(userIdString, expirationTime, config.UserConfig.JwtConfig.AccessSecret, refreshExpirationTime, config.UserConfig.JwtConfig.RefreshSecret)
	return token, err
}
