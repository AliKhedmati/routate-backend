package database

import (
	"context"
	"fmt"
	"github.com/AliKhedmati/routate-backend/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"sync"
	"time"
)

var (
	configs *config.Config
	client  *mongo.Client
	once    sync.Once
)

func Init() {
	configs = config.GetInstance()
	if err := Connect(); err != nil {
		panic(err)
	}
}

// Connect establishes a connection to MongoDB
func Connect() error {
	once.Do(func() {
		// Set client options
		clientOptions := options.Client().ApplyURI(
			fmt.Sprintf("mongodb://%s:%s@%s:%s",
				configs.Get("DB_USERNAME"),
				configs.Get("DB_PASSWORD"),
				configs.Get("DB_HOST"),
				configs.Get("DB_PORT"),
			),
		)

		// Set context
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		// Connect to MongoDB
		client, _ = mongo.Connect(ctx, clientOptions)

		// Check the connection
		_ = client.Ping(ctx, nil)

		defer func() {
			_ = Close()
		}()

		fmt.Println("Connected to MongoDB!")
	})

	return nil
}

// Close closes the connection to MongoDB
func Close() error {
	// Disconnect from MongoDB
	if err := client.Disconnect(context.Background()); err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Disconnected from MongoDB!")

	return nil
}
