package drivers

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"sync"
	"time"
)

type Redis struct {
	Host     string
	Port     string
	Password string
	Database int
}

var (
	client *redis.Client
	once   sync.Once
)

// Connect establishes a connection to Redis
func (r *Redis) Connect() error {
	once.Do(func() {
		// Create Redis client
		client = redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%s", r.Host, r.Port),
			Password: r.Password,
			DB:       0,
		})

		// Ping Redis server to ensure connection
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		_, err := client.Ping(ctx).Result()
		if err != nil {
			return
		}

		fmt.Println("Connected to Redis!")
	})

	return nil
}

// Close closes the connection to Redis
func (r *Redis) Close() error {
	err := client.Close()
	if err != nil {
		return err
	}

	fmt.Println("Disconnected from Redis!")

	return nil
}

func (r *Redis) GetCache() *redis.Client {
	return client
}
