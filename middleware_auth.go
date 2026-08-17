package main

import (
	"net/http"

	"github.com/BuildwithY26/go/internal/database"
	"github.com/BuildwithY26/go/internal/auth"
	
)

type authHandler func(http.ResponseWriter, *http.Request,database.User)

func (apiCfg *apiConfig) authMiddleware(handler authHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		 apikey, err := auth.GetAPIKey(&r.Header)
		if err != nil {
			respondwithError(w, http.StatusUnauthorized, "API key is missing or invalid")
			return
		}
	user,err := apiCfg.DB.GetUserByAPIKey(r.Context(), apikey)
	if err != nil {
		respondwithError(w, http.StatusUnauthorized, "API key is missing or invalid")
		return
	}

	handler(w, r, user)
}
}
