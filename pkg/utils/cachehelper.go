package utils

import (
	"context"
	"errors"
	"log"

	"github.com/Trung78z/web-service-gin/pkg/cache"
	"github.com/redis/go-redis/v9"
)

// SetCache: Save data to Redis
func SetCache(ctx context.Context, key string, value string, ttl int) error {
	err := cache.RD.Set(ctx, key, value, 0).Err() // TTL 0 is for no expiration
	if err != nil {
		log.Printf("failed to set cache for key %s: %v", key, err)
		return err
	}
	// log.Printf("cache set for key %s successfully", key)
	return nil
}

// GetCache: Get data from Redis
func GetCache(ctx context.Context, key string) (string, error) {
	val, err := cache.RD.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", errors.New("key not found in cache")
	} else if err != nil {
		log.Printf("failed to get cache for key %s: %v", key, err)
		return "", err
	}
	// log.Printf("cache hit for key %s: %s", key, val)
	return val, nil
}
