package models

import (
	"time"
)

type Movie struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"-"` // use the - directive to hide the field from appearing
	Title     string    `json:"title"`
	Year      int       `json:"year,omitzero"`
	Runtime   Runtime   `json:"runtime,omitzero"`
	Genres    []string  `json:"genres,omitempty"`
	Version   int       `json:"version"`
}
