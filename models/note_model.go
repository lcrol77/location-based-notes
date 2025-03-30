package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Note struct {
	Title string             `json:"title,omitempty" validate:"required"`
	Body  string             `json:"body,omitempty" validate:"required"`
	ID    primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
}
