package database

import (
	"github.com/AliKhedmati/routate-backend/src/config"
	databaseDrivers "github.com/AliKhedmati/routate-backend/src/database/drivers"
	"go.mongodb.org/mongo-driver/mongo"
)

type Database interface {
	Connect() error
	Close() error
	GetDatabase() *mongo.Database //Todo: should be more dynamic.
	GetClient() *mongo.Client     //Todo: should be more dynamic.
}

var (
	configs *config.Config
	db      Database
)

func Init() error {
	if err := Connect(); err != nil {
		return err
	}
	return nil
}

func Connect() error {
	var err error
	configs = config.GetInstance()
	db = &databaseDrivers.MongoDB{
		Username: configs.Get("DB_USERNAME"),
		Password: configs.Get("DB_PASSWORD"),
		Host:     configs.Get("DB_HOST"),
		Port:     configs.Get("DB_PORT"),
		Database: configs.Get("DB_DATABASE"),
	}

	// Connect to the database
	if err = db.Connect(); err != nil {
		return err
	}

	return err
}

func Close() error {
	if err := db.Close(); err != nil {
		return err
	}
	return nil
}

func GetDatabase() *mongo.Database {
	return db.GetDatabase()
}

func GetClient() *mongo.Client {
	return db.GetClient()
}
