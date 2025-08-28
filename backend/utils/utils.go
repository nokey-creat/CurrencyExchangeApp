// 存放一些工具函数
package utils

import (
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

// 生成JWT。接受用户名，返回token
func GenerateJWT(username string) (string, error) {

	claims := MyCustomClaims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)), //这是个时间戳
		},
	}

	mySecret := []byte("secret")

	//获得未签名的jwt的token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//签名
	signedToken, err := token.SignedString(mySecret)

	return "Bearer " + signedToken, err

}

// 比对加密后的密码,密码正确返回true
func CheckPassword(password string, hashPassword string) bool {
	//Returns nil on success, or an error on failure.
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	return err == nil
}
