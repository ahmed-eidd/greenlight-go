package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"greenlight.ahmedeid.net/internal/models"
	"greenlight.ahmedeid.net/internal/validator"
)

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title   string         `json:"title"`
		Year    int            `json:"year"`
		Runtime models.Runtime `json:"runtime"`
		Genres  []string       `json:"genres"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		var badRequestErr *badRequestError

		switch {
		case errors.As(err, &badRequestErr):
			app.badRequestResponse(w, r, err)
		default:
			app.serverErrorResponse(w, r, err)
		}

		return
	}

	v := validator.New()

	movie := models.Movie{
		Title:   input.Title,
		Year:    input.Year,
		Runtime: input.Runtime,
		Genres:  input.Genres,
	}

	if models.ValidateMovie(v, movie); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	movie, err = app.models.Movies.Insert(movie)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("v1/movies/%d", movie.ID))

	err = app.writeJSON(w, http.StatusCreated, envelope{"movie": movie}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {

	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	movie := models.Movie{
		ID:        id,
		CreatedAt: time.Now(),
		Title:     "Casablanca",
		Genres:    []string{"romance", "drama", "war"},
		Runtime:   models.Runtime(time.Hour + 42*time.Minute),
		Version:   1,
	}

	err = app.writeJSON(w, http.StatusOK, envelope{
		"movie": movie,
	}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
