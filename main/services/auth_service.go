package services

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"time"
)

var secretKey = []byte("abcd1234") // 密钥，用于签名 JWT

const (
	constUser123 = "user123"
	constPass123 = "pass123"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// JWT 响应结构
type JwtResponse struct {
	Token string `json:"token"`
}

// ValidateUser 方法验证用户名和密码，生成并返回 JWT
func ValidateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("POST /auth/login")
	// 解析请求体
	var loginRequest LoginRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&loginRequest); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// 验证用户名和密码
	if loginRequest.Username != constUser123 || loginRequest.Password != constPass123 {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	// 生成 JWT
	token, err := generateJWT(loginRequest.Username)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	// 返回 JWT
	response := JwtResponse{Token: token}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
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
