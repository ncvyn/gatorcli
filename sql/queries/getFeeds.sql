-- name: GetFeeds :many
SELECT id, created_at, updated_at, name, url, user_id
FROM feeds;
