package utils

import (
	"context"
	"time"

	"blog-go/database"
)

func SetCacheObject(ctx context.Context, key string, value interface{}) error {
	return database.RedisClient.Set(ctx, key, value, 0).Err()
}

func SetCacheObjectWithExpire(ctx context.Context, key string, value interface{}, expire time.Duration) error {
	return database.RedisClient.Set(ctx, key, value, expire).Err()
}

func GetCacheObject(ctx context.Context, key string) (string, error) {
	return database.RedisClient.Get(ctx, key).Result()
}

func DeleteCacheObject(ctx context.Context, key string) error {
	return database.RedisClient.Del(ctx, key).Err()
}

func IncreaseCacheMapValue(ctx context.Context, key, field string, delta int64) error {
	return database.RedisClient.HIncrBy(ctx, key, field, delta).Err()
}

func GetCacheMapValue(ctx context.Context, key, field string) (string, error) {
	return database.RedisClient.HGet(ctx, key, field).Result()
}
