package main

import (
	"fmt"
	"net/http"

	"github.com/alexsio03/go-backend/internal/auth"
	"github.com/alexsio03/go-backend/internal/db"
)

type authedHandler func(http.ResponseWriter, *http.Request, db.User)

func (apiCfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("Auth error: %v", err))
			return
		}

		user, err := apiCfg.DB.GetUserByAPI(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, 400, "Error getting user")
			return
		}

		handler(w, r, user)
	}
}
