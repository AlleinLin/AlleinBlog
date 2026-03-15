package database

import (
	"context"
	"fmt"

	"blog-go/config"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func InitRedis() error {
	cfg := config.AppConfig.Redis
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	ctx := context.Background()
	_, err := RedisClient.Ping(ctx).Result()
	return err
}

func GetRedis() *redis.Client {
	return RedisClient
}
