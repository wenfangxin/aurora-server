package core

import (
	"aurora/config"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type MyCustomClaims struct {
	UserId  int32
	RoleKey string
	jwt.StandardClaims
}

var jwtSecret = []byte(config.JwtSinge)

// CreateJwtToken 创建token
func CreateJwtToken(userId int32, roleKey string) (string, error) {
	nowTime := time.Now()
	expiredTime := nowTime.Add(72 * time.Hour)
	claims := MyCustomClaims{
		UserId:  userId,
		RoleKey: roleKey,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiredTime.Unix(),
		},
	}
	claimsToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return claimsToken.SignedString(jwtSecret)
}

// ParseToken 解析token
func ParseToken(token string) (*MyCustomClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	//解析存入的jwt信息
	claims, ok := tokenClaims.Claims.(*MyCustomClaims)
	if ok && tokenClaims.Valid {
		return claims, nil
	}
	return nil, err

}
