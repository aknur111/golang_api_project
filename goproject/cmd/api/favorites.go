package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (app *application) respondWithError(w http.ResponseWriter, code int, message string) {
	app.respondWithJSON(w, code, map[string]string{"error": message})
}

func (app *application) respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func (app *application) createFavoriteHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		UserID int `json:"user_id"`
		SongID int `json:"song_id"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	err = app.models.Favorites.Add(input.UserID, input.SongID)
	if err != nil {
		fmt.Println("ERROR while adding favorite:", err)

		app.respondWithError(w, http.StatusInternalServerError, "the server encountered a problem and could not process your request")
		return
	}

	app.respondWithJSON(w, http.StatusCreated, map[string]string{"message": "favorite added successfully"})
}

func (app *application) listFavoritesHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	userIDStr := query.Get("user_id")
	if userIDStr == "" {
		app.badRequestResponse(w, r, errMissingParameter("user_id"))
		return
	}

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	songs, err := app.models.Favorites.GetAll(userID)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	app.writeJSON(w, http.StatusOK, envelope{"favorites": songs}, nil)
}

func (app *application) deleteFavoriteHandler(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	userIDStr := params.ByName("user_id")
	songIDStr := params.ByName("song_id")

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		app.respondWithError(w, http.StatusBadRequest, "Invalid user_id")
		return
	}

	songID, err := strconv.Atoi(songIDStr)
	if err != nil {
		app.respondWithError(w, http.StatusBadRequest, "Invalid song_id")
		return
	}

	err = app.models.Favorites.Remove(userID, songID)
	if err != nil {
		app.respondWithError(w, http.StatusInternalServerError, "Failed to delete favorite")
		return
	}

	app.respondWithJSON(w, http.StatusOK, map[string]string{"message": "Favorite deleted successfully"})
}

func errMissingParameter(param string) error {
	return &ErrorString{Message: "missing parameter: " + param}
}

type ErrorString struct {
	Message string
}

func (e *ErrorString) Error() string {
	return e.Message
}
