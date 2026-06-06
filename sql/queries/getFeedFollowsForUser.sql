-- name: GetFeedFollowsForUser :many
SELECT id, created_at, updated_at, user_id, feed_id
FROM feed_follows
WHERE user_id = $1;
