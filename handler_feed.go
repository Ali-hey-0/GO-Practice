package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/wags-lane/rssagg/internal/database"
)

func (apiCfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request,user database.User) {
	type parameters struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}

	err := decoder.Decode(&params)	
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error Parsing Json: %v", err))
		return
	}

	feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name: params.Name,
		Url: params.URL,
		UserID: user.ID,
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error Creating User: %v", err))
		return
	}

	respondWithJson(w, 201, databaseFeedToFeed(feed))
}





func (apiCfg *apiConfig) handlerGetsFeeds(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}

	err := decoder.Decode(&params)	
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error Parsing Json: %v", err))
		return
	}

	feeds, err := apiCfg.DB.GetFeedsByID(r.Context())

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error Creating User: %v", err))
		return
	}

	respondWithJson(w, 201, databaseFeedsToFeeds(feeds))
}
