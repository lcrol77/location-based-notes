package models

import "time"

type Note struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Note  string `json:"note"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
