package repository

import (
	"context"
	"log/slog"
	"time"

	"github.com/redis/go-redis/v9"
)

var (
	ctx         = context.Background()
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})
	logger = slog.Default()
)

// InitRedis initializes the Redis client.
func InitRedis(redisAddr string) {
	// Update Redis address if provided
	RedisClient.Options().Addr = redisAddr

	// Test the connection with retry logic
	for i := 0; i < 5; i++ { // Try 5 times
		_, err := RedisClient.Ping(ctx).Result()
		if err != nil {
			logger.Error("Failed to connect to Redis", "attempt", i+1, "error", err)
			time.Sleep(2 * time.Second) // Wait before retrying
		} else {
			logger.Info("Connected to Redis successfully", "address", redisAddr)
			return // Exit the function if connection is successful
		}
	}

	// Panic if all attempts fail
	logger.Error("Could not connect to Redis after multiple attempts")
	panic("Failed to connect to Redis")
}

// GetCachedWeather checks if weather data is cached in Redis.
func GetCachedWeather(address string) (string, error) {
	val, err := RedisClient.Get(ctx, address).Result()
	if err == redis.Nil {
		// Cache miss
		return "", nil
	} else if err != nil {
		logger.Error("Failed to get cached weather data", "error", err, "address", address)
		return "", err
	}
	logger.Info("Cache hit", "address", address)
	return val, nil
}

// SetCachedWeather sets the weather data in Redis cache with an expiration time of 10 minutes.
func SetCachedWeather(address string, data string) error {
	err := RedisClient.Set(ctx, address, data, 10*time.Minute).Err()
	if err != nil {
		logger.Error("Failed to set cached weather data", "error", err, "address", address)
		return err
	}
	return nil
}
