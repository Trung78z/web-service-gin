package cache

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

var RD *redis.Client

func InitCache(url string) {
	opts, err := redis.ParseURL(url)
	if err != nil {
		log.Fatalf("failed to parse Redis URL: %v", err)
	}

	client := redis.NewClient(opts)

	// Ping Redis to check the connection
	ctx := context.Background()
	if err := client.Ping(ctx).Err(); err != nil {
		log.Fatalf("failed to connect to Redis: %v", err)
	}

	log.Println("Redis client initialized successfully")
	RD = client
}
