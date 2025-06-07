package main

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

var (
	redisClient *redis.Client
	ctx         = context.Background()
)

func InitRedis(addr string) {
	redisClient = redis.NewClient(&redis.Options{Addr: addr})
}

// getOrSetCache handles caching JSON responses as strings
func getOrSetCache(key string, ttl time.Duration, fetch func() (string, error)) (string, error) {
	if val, err := redisClient.Get(ctx, key).Result(); err == nil {
		return val, nil
	}
	data, err := fetch()
	if err != nil {
		return "", err
	}
	redisClient.Set(ctx, key, data, ttl)
	return data, nil
}
