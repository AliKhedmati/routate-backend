package cache

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

// Redis represents a Redis database connection
type Redis struct {
	client *redis.Client
}

// Connect establishes a connection to Redis
func (db *Redis) Connect() error {
	// Create Redis client
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// Ping Redis server to ensure connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := client.Ping(ctx).Result()
	if err != nil {
		return err
	}

	fmt.Println("Connected to Redis!")

	db.client = client

	return nil
}

// Close closes the connection to Redis
func (db *Redis) Close() error {
	err := db.client.Close()
	if err != nil {
		return err
	}

	fmt.Println("Disconnected from Redis!")

	return nil
}
