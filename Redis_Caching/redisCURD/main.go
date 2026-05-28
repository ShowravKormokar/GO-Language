package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

// Create (Write)
func WriteDataOnRedis(ctx context.Context, rdb *redis.Client, key string, value string, ttl time.Duration) {
	err := rdb.Set(ctx, key, value, ttl*time.Second).Err()
	if err != nil {
		log.Fatalf("Could not set value: %v", err.Error())
	}
	fmt.Println("Value set successfully on redis.")
}

// Read
func ReadDataFromRedis(ctx context.Context, rdb *redis.Client, key string) {
	val, err := rdb.Get(ctx, key).Result()
	if err != nil {
		log.Fatalf("Couldn't read value: %v", err.Error())
	}
	fmt.Println("Value from redis: ", val)
}

// Update (Overwrite value)
func UpdateDataOnRedis(ctx context.Context, rdb *redis.Client, key string, value string, ttl time.Duration) {
	upErr := rdb.Set(ctx, key, value, ttl*time.Second).Err()
	if upErr != nil {
		log.Fatalf("Could not update value: %v", upErr.Error())
	}
	fmt.Println("Value updated successfully on redis.")
}

// Delete
func DeleteDataFromRedis(ctx context.Context, rdb *redis.Client, key string) {
	delCnt, err := rdb.Del(ctx, key).Result()
	if err != nil {
		log.Fatalf("Couldn't delete value: %v", err.Error())
	}
	if delCnt > 0 {
		fmt.Println("Key deleted successfully.")
	} else {
		fmt.Println("Key not found to delete.")
	}
}

// Exists
func CheckKeyExists(ctx context.Context, rdb *redis.Client, key string) {
	sts, err := rdb.Exists(ctx, key).Result()
	if err != nil {
		log.Fatalf("Couldn't read value: %v", err.Error())
	}
	if sts > 0 {
		fmt.Println("Key exists in Redis.")
	} else {
		fmt.Println("Key does not exist.")
	}
}

func main() {
	ctx := context.Background()

	rbd := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	// Create(Write)
	WriteDataOnRedis(ctx, rbd, "key:name:", "Showrav Kormokar", 60)

	// Exists
	CheckKeyExists(ctx, rbd, "key:name:")

	// Read
	ReadDataFromRedis(ctx, rbd, "key:name:")

	time.Sleep(20 * time.Second) // Wait for the key to update

	// Update (Overwrite value)
	UpdateDataOnRedis(ctx, rbd, "key:name:", "Showrav Kormokar Up", 30)
	ReadDataFromRedis(ctx, rbd, "key:name:") // Read after update

	time.Sleep(20 * time.Second) // Wait for the key to delete

	// Delete
	DeleteDataFromRedis(ctx, rbd, "key:name:")
	// Check existence after delete key
	CheckKeyExists(ctx, rbd, "key:name:")

}
