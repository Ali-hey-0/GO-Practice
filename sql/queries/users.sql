-- name: CreateUser :one
-- Creates a new user with an auto-generated API key
INSERT INTO users(id, created_at, updated_at, name, api_key)
VALUES ($1, $2, $3, $4, encode(digest(gen_random_uuid()::text, 'sha256'), 'hex'))
RETURNING *;

-- name: GetUserByApiKey :one
-- Retrieves a user by their API key
SELECT * FROM users WHERE api_key = $1;