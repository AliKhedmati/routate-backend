package database

import (
	"github.com/AliKhedmati/routate-backend/config"
	"github.com/AliKhedmati/routate-backend/src/cache/drivers"
	drivers2 "github.com/AliKhedmati/routate-backend/src/database/drivers"
	"log"
)

type Database interface {
	Connect() error
	Close() error
}

var database Database

// GetDB returns a database instance.
func GetDB() Database {
	return database
}

func Init() {
	if err := config.Init(); err != nil {
		log.Fatalf("Failed to initialize configuration: %v", err)
	}
	configs := config.GetInstance()

	switch configs.Get("DB_CONNECTION") {
	case "mongodb":
		database = &drivers2.MongoDB{
			Username: configs.Get("DB_USERNAME"),
			Password: configs.Get("DB_PASSWORD"),
			Host:     configs.Get("DB_HOST"),
			Port:     configs.Get("DB_PORT"),
		}
	case "redis":
		database = &drivers.Redis{}
	default:
		log.Fatalf("Unknown database connection.")
	}

	// Connect to the database
	if err := database.Connect(); err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	defer func() {
		// Ensure the database connection is closed when the program exits
		if err := database.Close(); err != nil {
			log.Printf("Failed to close the database connection: %v", err)
		}
	}()
}
