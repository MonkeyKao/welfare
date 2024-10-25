package utils

import (
	//"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

var Secret = []byte("q5QEligUIlqOFCIQ")

type UserClaims struct {
	UserID uint
	jwt.StandardClaims
}

func GenerateToken(UserID uint, expireTime time.Duration) (string, error) {
	cla := UserClaims{
		UserID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(expireTime).Unix(), // 过期时间
			Issuer:    "SMS",                             // 签发人
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, cla)
	tokenString, _ := token.SignedString(Secret)

	return tokenString, nil
}
