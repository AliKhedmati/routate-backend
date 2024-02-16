package drivers

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sync"
	"time"
)

type MongoDB struct {
	Username string
	Password string
	Host     string
	Port     string
	Database string
}

var (
	client *mongo.Client
	db     *mongo.Database
	once   sync.Once
)

// Connect establishes a connection to MongoDB
func (MongoDB *MongoDB) Connect() error {

	var err error

	once.Do(func() {
		// Set client options
		clientOptions := options.Client().ApplyURI(
			fmt.Sprintf("mongodb://%s:%s@%s:%s",
				MongoDB.Username,
				MongoDB.Password,
				MongoDB.Host,
				MongoDB.Port,
			),
		)

		// Set context
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		// Connect to MongoDB
		if client, err = mongo.Connect(ctx, clientOptions); err != nil {
			return
		}

		// Check the connection
		if err := client.Ping(ctx, nil); err != nil {
			return
		}

		// Set the db.
		db = client.Database(MongoDB.Database)

		fmt.Println("Connected to MongoDB!")
	})

	return err
}

// Close closes the connection to MongoDB
func (MongoDB *MongoDB) Close() error {
	// Disconnect from MongoDB
	if err := client.Disconnect(context.Background()); err != nil {
		return err
	}

	fmt.Println("Disconnected from MongoDB!")

	return nil
}

// GetDatabase returns an instance of MongoDB database
func (MongoDB *MongoDB) GetDatabase() *mongo.Database {
	return db
}

// GetClient returns an instance of MongoDB client
func (MongoDB *MongoDB) GetClient() *mongo.Client {
	return client
}
