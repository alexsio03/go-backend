package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/alexsio03/go-backend/internal/db"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user db.User) {
	type params struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}
	decoder := json.NewDecoder(r.Body)

	parameters := params{}
	err := decoder.Decode(&parameters)
	if err != nil {
		respondWithError(w, 400, "Invalid request payload")
		return
	}

	feed, err := apiCfg.DB.CreateFeed(r.Context(), db.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      parameters.Name,
		Url:       parameters.URL,
		UserID:    user.ID,
	})
	if err != nil {
		respondWithError(w, 400, "Error creating user")
		return
	}

	respondWithJSON(w, 200, dbFeedToFeed(feed))
}

func (apiCfg *apiConfig) handlerGetFeeds(w http.ResponseWriter, r *http.Request) {
	feed, err := apiCfg.DB.GetFeeds(r.Context())
	if err != nil {
		respondWithError(w, 400, "Error getting feeds")
		return
	}

	respondWithJSON(w, 200, dbFeedsToFeed(feed))
}
