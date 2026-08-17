package main

import (
	"encoding/json"
	"net/http"

	"time"

	"github.com/BuildwithY26/go/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct{
		Name string `json:"name"`
		URL string `json:"url"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondwithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	feed,err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name: params.Name,
		Url: params.URL,
		UserID: user.ID,
	})
	if err != nil {
		respondwithError(w, http.StatusInternalServerError, "Failed to create feed")
		return
	}
	respondwithJSON(w, http.StatusCreated, databaseFeedToFeed(feed))
}

func (apiCfg *apiConfig) handlerGetFeeds(w http.ResponseWriter, r *http.Request, user database.User) {
	feeds, err := apiCfg.DB.GetFeeds(r.Context())
	if err != nil {
		respondwithError(w, http.StatusInternalServerError, "Failed to retrieve feeds")
		return
	}
	respondwithJSON(w, http.StatusOK, feeds)
}