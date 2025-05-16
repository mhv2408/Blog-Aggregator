-- +goose UP
ALTER TABLE feed ADD COLUMN last_fetched_at TIMESTAMP;

-- +goose DOWN
ALTER TABLE feed DROP COLUMN last_fetched_at;