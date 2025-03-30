package models

type Note struct {
	Title string `json:"title,omitempty" validate:"required"`
	Body  string `json:"body,omitempty" validate:"required"`
}
