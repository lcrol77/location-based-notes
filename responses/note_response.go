package responses

import (
	"lbn/models"
)

type NoteResponse struct {
	Status   int           `json:"status"`
	Message  string        `json:"message"`
	Note     models.Note   `json:"note,omitempty"`
	Errors   []string      `json:"errors,omitempty"`
	Inserted []interface{} `json:"inserted,omitempty"`
}

type NotesResponse struct {
	Status  int           `json:"status"`
	Message string        `json:"message"`
	Notes   []models.Note `json:"Notes"`
	Errors  []string      `json:"errors,omitempty"`
}
