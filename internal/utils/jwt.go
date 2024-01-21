package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type CustomClaims struct {
	UserID int64
	jwt.RegisteredClaims
}

const secret = "etov-jwt-secret-key"

// TokenExpireDuration Token过期时间为一个月
const TokenExpireDuration = time.Hour * 24 * 30

func GenerateTokenUsingHs256(userId int64) (string, error) {
	claim := CustomClaims{
		UserID: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "Auth_Server",                                           // 签发者
			Subject:   "etov-server",                                           // 签发对象
			Audience:  jwt.ClaimStrings{"ETOV-WEB"},                            // 签发受众
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenExpireDuration)), // 过期时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),                          // 签发时间
			ID:        GenerateSalt(8),                                         // 类似于盐值
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claim).SignedString([]byte(secret))
	return token, err
}

func ParseTokenHs256(token string) (*CustomClaims, error) {
	tk, err := jwt.ParseWithClaims(token, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil //返回签名密钥
	})
	if err != nil {
		return nil, err
	}

	if !tk.Valid {
		return nil, errors.New("claim invalid")
	}

	claims, ok := tk.Claims.(*CustomClaims)
	if !ok {
		return nil, errors.New("invalid claim type")
	}

	return claims, nil
}
