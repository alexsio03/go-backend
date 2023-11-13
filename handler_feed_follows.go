package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/alexsio03/go-backend/internal/db"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user db.User) {
	type params struct {
		FeedID uuid.UUID `json:"feed_id"`
	}
	decoder := json.NewDecoder(r.Body)

	parameters := params{}
	err := decoder.Decode(&parameters)
	if err != nil {
		respondWithError(w, 400, "Invalid request payload")
		return
	}

	feedFollow, err := apiCfg.DB.CreateFeedFollow(r.Context(), db.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    parameters.FeedID,
	})
	if err != nil {
		respondWithError(w, 400, "Error creating user")
		return
	}

	respondWithJSON(w, 200, dbFeedFollowToFeedFollow(feedFollow))
}

func (apiCfgg *apiConfig) handlerGetFeedFollows(w http.ResponseWriter, r *http.Request, user db.User) {
	feedFollows, err := apiCfgg.DB.GetFeedFollows(r.Context(), user.ID)
	if err != nil {
		respondWithError(w, 400, "Couldn't get feed follows")
		return
	}

	respondWithJSON(w, 200, dbFeedFollowsToFeedFollows(feedFollows))
}

func (apiCfg *apiConfig) handlerDeleteFeedFollow(w http.ResponseWriter, r *http.Request, user db.User) {
	feedFollowIDStr := chi.URLParam(r, "feedFollowID")
	feedFollowID, err := uuid.Parse(feedFollowIDStr)
	if err != nil {
		respondWithError(w, 400, "Invalid feed follow ID")
		return
	}
	err = apiCfg.DB.DeleteFeedFollow(r.Context(), db.DeleteFeedFollowParams{ID: feedFollowID, UserID: user.ID})
	if err != nil {
		respondWithError(w, 400, "Couldn't delete feed follow")
		return
	}
	respondWithJSON(w, 200, nil)
}
