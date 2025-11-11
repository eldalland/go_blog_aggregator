-- name: DropLastFetchedAt :exec
ALTER TABLE feeds
DROP COLUMN last_fetched_at;