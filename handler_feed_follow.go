package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/wags-lane/rssagg/internal/database"
)

func (apiCfg *apiConfig) handlerFollowsFeed(w http.ResponseWriter, r *http.Request,user database.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}

	err := decoder.Decode(&params)	
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error Parsing Json: %v", err))
		return
	}

	feedFollows, err := apiCfg.DB.CreateFeedFollows(r.Context(), database.CreateFeedFollowsParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID: user.ID,
		FeedID: params.FeedID,
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error Creating feed follow: %v", err))
		return
	}

	respondWithJson(w, 201, databaseFeedsFollowToFeeds(feedFollows))
}

func (apiCfg *apiConfig) handlerGetsFeedFollows(w http.ResponseWriter, r *http.Request,user database.User) {
	feedFollow, err := apiCfg.DB.GetFeedFollows(r.Context(),user.ID)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error Getting Feed Follows: %v", err))
		return
	}

	respondWithJson(w, 201, databaseFeedsToFollowsFeedsFollows(feedFollow))
}

func (apiCfg *apiConfig) handlerDeleteFeedFollow(w http.ResponseWriter, r *http.Request,user database.User) {
	feedFollowIDStr := chi.URLParam(r, "feedFollowID")
	feedFollowID, err := uuid.Parse(feedFollowIDStr)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error Parsing UUID: %v", err))
		return
	} 

	err = apiCfg.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		ID: feedFollowID,
		UserID: user.ID,
	})

	if err != nil {
		respondWithError(w,400,fmt.Sprintf("Error Deleting Users Feed Follows: %v", err))
		return
	} 
	
	respondWithJson(w, 200, struct{}{})
}