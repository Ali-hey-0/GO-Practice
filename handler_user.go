package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/wags-lane/rssagg/internal/auth"
	"github.com/wags-lane/rssagg/internal/database"
)

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}

	err := decoder.Decode(&params)	
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error Parsing Json: %v", err))
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name: params.Name,
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error Creating User: %v", err))
		return
	}

	respondWithJson(w, 200, databaseUserToUser(user))
}

func (apiCfg *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request,user database.User) {
	apiKey, err := auth.GetApiKey(r.Header)
	if err != nil {
		respondWithError(w, 403, fmt.Sprintf("Auth Error: %v", err))
		return
	}

	user, err = apiCfg.DB.GetUserByApiKey(r.Context(), apiKey)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error Getting User: %v", err))
		return
	}
	respondWithJson(w, 200, databaseUserToUser(user))
}