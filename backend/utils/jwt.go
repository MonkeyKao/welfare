package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

var Secret = []byte("q5QEligUIlqOFCIQ")

type UserClaims struct {
	UserID uint
	Code   string
	jwt.StandardClaims
}

func GenerateToken(UserID uint, Code string, expireTime time.Duration) (string, error) {
	cla := UserClaims{
		UserID,
		Code,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(expireTime).Unix(), // 过期时间
			Issuer:    "SMS",                             // 签发人
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, cla)
	tokenString, _ := token.SignedString(Secret)

	return tokenString, nil
}

func ParseToken(tokenString string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return Secret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
