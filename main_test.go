package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func main() {
	errLoadEnv := godotenv.Load()
	if errLoadEnv != nil {
		log.Fatal("Error loading .env file")
	}
	opts, errLoadURL := redis.ParseURL(os.Getenv("REDIS_URL"))
	if errLoadURL != nil {
		panic(errLoadURL)
	}
	client := redis.NewClient(opts)
	defer client.Close()

	err := client.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := client.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := client.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
}
