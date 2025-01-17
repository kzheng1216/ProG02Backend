package service

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// ValidateUser 方法验证用户名和密码，生成并返回 JWT
func ValidateUser(username string, password string) (string, error) {
	fmt.Println("POST /auth/login")

	// 验证用户名和密码
	if username != constUser123 || password != constPass123 {
		return "", fmt.Errorf("Invalid username or password")
	}

	// 生成 JWT
	token, err := generateJWT(username)
	if err != nil {
		fmt.Errorf("Error generating token" + err.Error())
		return "", err
	}
	return token, nil
}

// 生成 JWT
func generateJWT(username string) (string, error) {
	// 设置 JWT payload
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 1).Unix(), // 设置过期时间为 1 小时
	}

	// 创建 token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 使用 secret key 签名生成 token
	return token.SignedString(secretKey)
}
