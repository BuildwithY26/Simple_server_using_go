package main

import (
	"encoding/json"
	"net/http"

	"time"

	"github.com/BuildwithY26/go/internal/database"
	"github.com/google/uuid"
	"github.com/go-chi/chi"
)

func (apiCfg *apiConfig) handlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct{
		Name string `json:"name"`
		FeedID uuid.UUID `json:"feed_id"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondwithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	feedFollow, err := apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID: user.ID,
		FeedID: params.FeedID,
	})
	if err != nil {
		respondwithError(w, http.StatusInternalServerError, "Failed to create feed follow")
		return
	}
	respondwithJSON(w, http.StatusCreated, databaseFeedFollowToFeedFollow(feedFollow))
}

func (apiCfg *apiConfig) handlerGetFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollows, err := apiCfg.DB.GetFeedFollows(r.Context(), user.ID)
	if err != nil {
		respondwithError(w, http.StatusInternalServerError, "Failed to retrieve feed follows")
		return
	}
	respondwithJSON(w, http.StatusOK, feedFollows)
}

func (apiCfg *apiConfig) handlerDeleteFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowIDStr := chi.URLParam(r, "FeedFollowID")
	feedFollowID, err := uuid.Parse(feedFollowIDStr)
	if err != nil {
		respondwithError(w, http.StatusBadRequest, "Invalid FeedFollowID")
		return
	}

	err = apiCfg.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		ID: feedFollowID,
		UserID: user.ID,
	})
	if err != nil {
		respondwithError(w, http.StatusInternalServerError, "Failed to delete feed follow")
		return
	}
	respondwithJSON(w, http.StatusOK, map[string]string{"message": "Unfollowed successfully"})
}