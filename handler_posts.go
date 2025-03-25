package main

import (
	"net/http"
	"strconv"

	"github.com/wags-lane/rssagg/internal/database"
)

func (cfg *apiConfig) handlerPostsGet(w http.ResponseWriter, r *http.Request, user database.User) {
	limitStr := r.URL.Query().Get("limit")
	limit := 10
	if specifiedLimit, err := strconv.Atoi(limitStr); err == nil {
		limit = specifiedLimit
	}

	// Log the limit for now to avoid unused variable warning
	_ = limit

	// Since we don't have the actual database function, create a temporary placeholder
	// In a real implementation, you would call the database function to get posts
	respondWithError(w, http.StatusNotImplemented, "Posts functionality not yet implemented")
}
