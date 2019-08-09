package models

import (
	"time"
)

type Music struct {
	ID        int       `json:"id,omitempty"`
	Name      string    `json:"name"`
	Year      string    `json:"year"`
	Style     string    `json:"style"`
	CreatedAt time.Time `json:"createAt"`
	UpdatedAt time.Time `json:"updateAt"`
}
