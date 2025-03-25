-- +goose Up

ALTER TABLE feeds ADD COLUMN last_fetched_at TIMESTAMP ;




-- +goose Down

ALTER TABLE feeds ADD COLUMN last_fetched_at;



