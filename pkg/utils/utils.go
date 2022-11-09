package utils

import (
	"github.com/dgrijalva/jwt-go"
	"strings"
	"time"
)

var JWTsecret = []byte("ABAB") // 私钥

type Claims struct {
	Id       uint   `json:"id"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
	jwt.StandardClaims
}

// GenerateToken 签发token
func GenerateToken(id uint, username string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(24 * time.Hour)
	claims := Claims{
		Id:       id,
		UserName: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "todo_list",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(JWTsecret)
	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaim, err := jwt.ParseWithClaims(strings.Split(token, "Bearer ")[1], &Claims{},
		func(token *jwt.Token) (interface{}, error) {
			return JWTsecret, nil
		})
	if tokenClaim != nil {
		if claims, ok := tokenClaim.Claims.(*Claims); ok && tokenClaim.Valid {
			return claims, nil
		}
	}
	return nil, err
}
