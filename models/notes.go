package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Note struct {
	Id        primitive.ObjectID `bson:"id,omitempty"`
	Title     string             `bson:"title"`
	Note      string             `bson:"note"`
	CreatedAt time.Time          `bson:"created_at, omitempty"`
	UpdatedAt time.Time          `bson:"updated_at, omitempty"`
}
