package common

import (
	"gin-vue/model"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// 定义密钥
var jwtKey = []byte("andklankdla")

// 自定义jwt加密的结构体，这里加上了用户id
type Claims struct {
	UserId uint
	jwt.StandardClaims
}

func ReleaseToken(user model.User) (string, error) {
	expirationTime := time.Now().Add(7 * 24 * time.Hour)
	claims := &Claims{
		UserId: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "gin",
			Subject:   "user token",
		},
	}
	// 初始化对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 生成token字符串
	signedString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return signedString, nil
}

func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtKey, nil

	})

	return token, claims, err
}
