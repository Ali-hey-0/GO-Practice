-- Active: 1742858015285@@127.0.0.1@5432@postgres@public
-- +goose Up

CREATE TABLE users (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name TEXT NOT NULL
);

-- +goose Down
DROP TABLE users;

