package main

import (
	"fmt"
	"net/http"

	"github.com/wags-lane/rssagg/internal/auth"
	"github.com/wags-lane/rssagg/internal/database"
)

type authedHandler func(http.ResponseWriter,*http.Request,database.User)


func (apiCfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetApiKey(r.Header)
	if err != nil {
		respondWithError(w, 403, fmt.Sprintf("Auth Error: %v", err))
		return
	}

	user, err := apiCfg.DB.GetUserByApiKey(r.Context(), apiKey)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error Getting User: %v", err))
		return
	}
	handler(w,r,user)

	}
}

