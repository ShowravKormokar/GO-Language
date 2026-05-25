package main

import (
	"context"
	"fmt"
	"log"

	redis "github.com/redis/go-redis/v9"
)

func main() {

	// Create Context for Radis Operation
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis server address
		Password: "",               // No password by default
		DB:       0,                // Use default DB
	})

	err := rdb.Set(ctx, "key:", "value", 0).Err()
	if err != nil {
		log.Fatalf("Could not set value: %v", err)
	}

	val, err := rdb.Get(ctx, "key:").Result()
	if err != nil {
		log.Fatalf("Could not get value: %v", err)
	}

	fmt.Println("Value from Redis:", val)

}
