package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Note struct {
	Title     string             `json:"title,omitempty" validate:"required"`
	Body      string             `json:"body,omitempty" validate:"required"`
	ID        primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at,omitempty"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at,omitempty"`
}
