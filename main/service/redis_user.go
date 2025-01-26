package service

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
)

func (r *RedisClient) SetUser(key string, user User) error {
	userBytes, err := json.Marshal(user)
	if err != nil {
		return err
	}
	return r.client.Set(ctx, key, userBytes, 0).Err()
}

func (r *RedisClient) GetUser(userID string) (User, error) {
	var user User
	value, err := r.client.Get(ctx, userID).Result()
	if err == redis.Nil {
		return user, fmt.Errorf("user %s not found", userID)
	} else if err != nil {
		return user, fmt.Errorf("failed to retrieve user %s: %w", userID, err)
	}

	err = json.Unmarshal([]byte(value), &user)
	if err != nil {
		return user, fmt.Errorf("failed to parse user data for %s: %w", userID, err)
	}
	fmt.Println(user)
	return user, nil
}
