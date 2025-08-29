// 存放一些工具函数
package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// bcrypt算法加密的封装
func HashPassword(pwd string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	return string(hash), err
}

// JWT中传递的声明（claims）
type MyCustomClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// JWT签名密钥
var mySecret = []byte("secret")

// 生成JWT。接受用户名，返回token
func GenerateJWT(username string) (string, error) {

	claims := MyCustomClaims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)), //这是个时间戳
		},
	}

	//获得未签名的jwt的token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//签名
	signedToken, err := token.SignedString(mySecret)

	return "Bearer " + signedToken, err

}

// 验证并解析jwt，如果jwt有效则返回用户名
func ParseJWT(tokenString string) (string, error) {
	//去掉bearer头
	if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
		tokenString = tokenString[7:]
	}

	//解析并验证jwt
	//如果签名无效或过期，会返回error，token.valid被设为false
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (any, error) {
		//限制签名算法
		if token.Method.Alg() != jwt.SigningMethodHS256.Alg() {
			return nil, errors.New("unexpected Signing Method")
		}
		return mySecret, nil
	})
	if err != nil {
		return "", err
	} else if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		//验证成功
		return claims.Username, nil
	} else {
		return "", errors.New("unknown claims type, cannot proceed")
	}
}

// 比对加密后的密码,密码正确返回true
func CheckPassword(password string, hashPassword string) bool {
	//Returns nil on success, or an error on failure.
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	return err == nil
}
