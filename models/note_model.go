package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Note struct {
	Id    primitive.ObjectID `json:"id,omitempty"` // FIXME: pretty sure this is duoplicated and _id is fine
	Title string             `json:"title,omitempty" validate:"required"`
	Body  string             `json:"body,omitempty" validate:"required"`
	// Location string             `json:"location,omitempty" validate:"required"`
}
