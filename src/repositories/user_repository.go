package repositories

import (
	"context"
	"github.com/AliKhedmati/routate-backend/src/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// UserRepository Represents the repository.
type UserRepository struct {
	collection *mongo.Collection
}

// NewUserRepository Returns an instance of UserRepository
func NewUserRepository(collection *mongo.Collection) *UserRepository {
	return &UserRepository{
		collection: collection,
	}
}

// FindByID Finds user from database by its id.
func (repository *UserRepository) FindByID(ctx context.Context, id string) (*models.User, error) {
	var user models.User
	idObj, _ := primitive.ObjectIDFromHex(id)
	err := repository.collection.FindOne(ctx, bson.M{"_id": idObj}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Create Creates new user.
func (repository *UserRepository) Create(ctx context.Context, user *models.User) error {
	_, err := repository.collection.InsertOne(ctx, user)
	return err
}
