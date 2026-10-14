package models

import (
	"time"

	"greenlight.ahmedeid.net/internal/validator"
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

func ValidateMovie(v *validator.Validator, movie Movie) {
	// title validation
	v.Check(movie.Title != "", "title", "must be provided")
	v.Check(len(movie.Title) <= 500, "title", "must not be more than 500 bytes long")

	// year validation
	v.Check(movie.Year != 0, "year", "must be provided")
	v.Check(movie.Year >= 1888, "year", "must be greater than 1888")
	v.Check(movie.Year <= time.Now().Year(), "year", "must not be in the future")

	// runtime validation
	v.Check(movie.Runtime != 0, "runtime", "must be provided")
	v.Check(movie.Runtime > 0, "runtime", "must be positive")

	// genres validation
	v.Check(movie.Genres != nil, "genres", "must be provided")
	v.Check(len(movie.Genres) >= 1, "genres", "must contain at least 1 genre")
	v.Check(len(movie.Genres) <= 5, "genres", "must not contain more than 5 genres")
	v.Check(validator.Unique(movie.Genres), "genres", "must not contain duplicate values")
}
