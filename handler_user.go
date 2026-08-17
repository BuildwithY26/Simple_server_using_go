package main

import (
	"encoding/json"
	"net/http"

	"time"

	"github.com/BuildwithY26/go/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct{
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondwithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	user,err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Username: params.Name,
	})
	if err != nil {
		respondwithError(w, http.StatusInternalServerError, "Failed to create user")
		return
	}
	respondwithJSON(w, http.StatusCreated, databaseUserToUser(user))
}


func (apiCfg *apiConfig) handlerGetUserByAPIKey(w http.ResponseWriter, r *http.Request) {
	type parameters struct{
		Name string `json:"api_key"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondwithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	user,err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Username: params.Name,
	})
	if err != nil {
		respondwithError(w, http.StatusInternalServerError, "Failed to create user")
		return
	}
	respondwithJSON(w, http.StatusCreated, databaseUserToUser(user))
}

func (apiCfg *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request, user database.User) {
		respondwithJSON(w, http.StatusOK, databaseUserToUser(user))
}

func (apiCfg *apiConfig) handlerGetPostsForUser(w http.ResponseWriter, r *http.Request, user database.User) {
		posts, err := apiCfg.DB.GetPostsForUser(r, Context(), GetPostsForUserParams(
			user.ID,
			10,
		))
		if err != nil {
		respondwithError(w, http.StatusInternalServerError, "Failed to get post")
		return
	}

	return
}