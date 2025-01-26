package service

import (
	"context"
	"github.com/go-redis/redis/v8"
)

type (
	User struct {
		ID    int    `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	JwtResponse struct {
		Token string `json:"token"`
	}

	LoginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
)

const (
	constUser123  = "user123"
	constPass123  = "pass123"
	redisHost     = "106.15.33.153:6379"
	redisDB       = 0
	redisUsername = "mystic"
	redisPassword = "Zaq1@wsx"
)

var secretKey = []byte("abcd1234") // 密钥，用于签名 JWT
var ctx = context.Background()

type RedisClient struct {
	client *redis.Client
}

func NewRedisClient() *RedisClient {
	addr := redisHost
	username := redisUsername
	password := redisPassword
	db := redisDB
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Username: username,
		Password: password,
		DB:       db,
	})
	return &RedisClient{client: client}
}
