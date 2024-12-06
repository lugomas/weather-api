package repository

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

var ctx = context.Background()

// Connect to Redis
var RedisClient = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "",
	DB:       0,
})

// InitRedis Init initializes the Redis client.
func InitRedis(redisAddr string) {

	// Test the connection
	_, err := RedisClient.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
	log.Printf("INFO: connected to redis successfully")
}

// GetCachedWeather checks if weather data is cached in Redis.
func GetCachedWeather(address string) (string, error) {
	val, err := RedisClient.Get(ctx, address).Result()
	if err == redis.Nil {
		// Cache miss
		return "", nil
	} else if err != nil {
		return "", err
	}
	return val, nil
}

// SetCachedWeather sets the weather data in Redis cache with an expiration time of 10 minutes.
func SetCachedWeather(address string, data string) error {
	err := RedisClient.Set(ctx, address, data, 10*time.Minute).Err()
	if err != nil {
		return err
	}
	return nil
}
