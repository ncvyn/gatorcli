-- name: GetNextFeedToFetch :one
SELECT id, created_at, updated_at, name, url, user_id
FROM feeds
WHERE last_fetched_at IS NULL OR last_fetched_at < NOW() - INTERVAL '1 minute'
ORDER BY updated_at ASC NULLS FIRST
LIMIT 1;
