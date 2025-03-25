package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetApiKey extracts the API key from the Authorization header
// Expected format: "ApiKey <key>"
// Returns the API key and any error encountered during extraction
func GetApiKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("no authorization header")
	}

	vals := strings.Split(val, " ")

	if len(vals) != 2 {
		return "", errors.New("invalid authorization header")
	}

	if vals[0] != "ApiKey" {
		return "", errors.New("invalid first part of authorization header")
	}

	return vals[1], nil
}