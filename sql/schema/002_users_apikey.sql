-- +goose Up
-- Add pgcrypto extension for UUID and hash functions
CREATE EXTENSION IF NOT EXISTS pgcrypto;

-- Add api_key column with a default value using UUID and SHA256
ALTER TABLE users ADD COLUMN api_key VARCHAR(64) UNIQUE NOT NULL DEFAULT 
    encode(digest(gen_random_uuid()::text, 'sha256'), 'hex');

-- +goose Down
-- Simply drop the column
ALTER TABLE users DROP COLUMN IF EXISTS api_key;

