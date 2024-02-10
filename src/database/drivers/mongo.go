package drivers

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

// MongoDB represents a MongoDB database connection
type MongoDB struct {
	Host     string
	Port     string
	Username string
	Password string

	client   *mongo.Client
	database *mongo.Database
}

// Connect establishes a connection to MongoDB
func (db *MongoDB) Connect() error {
	// Set client options
	clientOptions := options.Client().ApplyURI(
		fmt.Sprintf("mongodb://%s:%s@%s:%s",
			db.Username,
			db.Password,
			db.Host,
			db.Port,
		),
	)

	// Set context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return err
	}

	// Check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return err
	}

	fmt.Println("Connected to MongoDB!")

	db.client = client
	db.database = client.Database("mydb") // Replace "mydb" with your database name

	return nil
}

// Close closes the connection to MongoDB
func (db *MongoDB) Close() error {
	// Disconnect from MongoDB
	err := db.client.Disconnect(context.Background())
	if err != nil {
		return err
	}

	fmt.Println("Disconnected from MongoDB!")

	return nil
}
