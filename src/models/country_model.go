package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Country struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name" bson:"name"`
	MobileRegex string             `json:"mobile_regex" bson:"mobile_prefix"`
	CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at" bson:"updated_at"`
}
