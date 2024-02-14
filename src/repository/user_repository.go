package repository

import (
	"context"
	"github.com/AliKhedmati/routate-backend/src/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	FindByID(ctx context.Context, id string) (*model.User, error)
	Create(ctx context.Context, user *model.User) error
	Update(user *model.User) error
}

type MongoUserRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(db *mongo.Database) *MongoUserRepository {
	return &MongoUserRepository{db.Collection("users")}
}

func (repository *MongoUserRepository) FindByID(ctx context.Context, id string) (*model.User, error) {
	var user model.User
	err := repository.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (repository *MongoUserRepository) Create(ctx context.Context, user *model.User) error {
	_, err := repository.collection.InsertOne(ctx, user)
	return err
}
