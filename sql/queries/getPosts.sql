-- name: GetPosts :many
SELECT id, created_at, updated_at, title, url, description, published_at, feed_id
FROM posts
ORDER BY published_at DESC
LIMIT $1;
