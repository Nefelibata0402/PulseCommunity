package jwts

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type JwtToken struct {
	AccessToken  string
	RefreshToken string
	AccessExp    int64
	RefreshExp   int64
}

func CreateToken(val string, exp time.Duration, secret string, refreshExp time.Duration, refreshSecret string, ssid string) (*JwtToken, error) {
	aExp := time.Now().Add(exp).Unix()
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"token": val,
		"exp":   aExp,
		"ssid":  ssid,
	})
	aToken, err := accessToken.SignedString([]byte(secret))
	if err != nil {
		return nil, err
	}
	rExp := time.Now().Add(refreshExp).Unix()
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"token": val,
		"exp":   rExp,
		"ssid":  ssid,
	})
	rToken, err := refreshToken.SignedString([]byte(refreshSecret))
	if err != nil {
		return nil, err
	}
	return &JwtToken{
		AccessExp:    aExp,   //访问时间
		AccessToken:  aToken, //访问令牌
		RefreshExp:   rExp,   //刷新时间
		RefreshToken: rToken, //刷新令牌
	}, nil
}

func ParseToken(tokenString string, secret string) (string, string, error) {
	//解析token 验证签名
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return "", "", err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		val := claims["token"].(string)
		exp := int64(claims["exp"].(float64))
		Ssid := claims["ssid"].(string)
		if exp <= time.Now().Unix() {
			return "", "", errors.New("token过期了")
		}
		return val, Ssid, nil
	} else {
		return "", "", err
	}
}
