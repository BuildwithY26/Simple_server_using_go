-- +goose Up

ALTER TABLE feeds ADD COLUMN last_fetched_at TIMESTAMP;

-- +goose Down
ALTER TABLE feeds DROP COLUMN last_fetched_at;

-- name: MarkFeedAsFetched :one
UPDATE feeds SET last_fetched_at = NOW(),
updated_at = NOW() 
WHERE id = $1 
RETURNING *;