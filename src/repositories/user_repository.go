package repositories

import (
	"context"
	"github.com/AliKhedmati/routate-backend/src/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(collection *mongo.Collection) *UserRepository {
	return &UserRepository{
		collection: collection,
	}
}

func (repository *UserRepository) FindByID(ctx context.Context, id string) (*models.User, error) {
	var user models.User
	err := repository.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (repository *UserRepository) Create(ctx context.Context, user *models.User) error {
	_, err := repository.collection.InsertOne(ctx, user)
	return err
}
