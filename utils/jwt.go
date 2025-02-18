package utils

import (
	"time"
	"travel-server/global"

	"github.com/golang-jwt/jwt/v5"
)

type MyClaims struct {
	UserId int
	jwt.RegisteredClaims
}

func JWTSign(userId int) (string, error) {
	// 密钥
	secret := []byte(global.Config.Jwt.Secret)
	// clamis
	claims := MyClaims{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(global.Config.Jwt.Expires))), // 过期时间
			Issuer:    global.Config.Jwt.Issuer,                                                                 // 签发人
		},
	}
	// 生成Token，指定签名算法和claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 签名
	sign, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}
	return sign, nil
}

func JWTVerify(sign string) (*MyClaims, error) {
	secret := []byte(global.Config.Jwt.Secret)
	claims := &MyClaims{}
	_, err := jwt.ParseWithClaims(sign, claims, func(t *jwt.Token) (any, error) {
		return secret, nil
	})
	return claims, err
}
