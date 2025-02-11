package main

import (
	"fmt"
	"net/http"

	"github.com/Mathis-zls/RSS/internal/database"
	"github.com/Mathis-zls/RSS/internal/database/auth"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middelwareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apikey, err := auth.GetAPIkey(r.Header)

		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("Auth error%v", err))
			return
		}
		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apikey)
		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("Counldn't get user error%v", err))
			return
		}
		handler(w, r, user)
	}
}
