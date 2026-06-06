-- name: GetUserById :one
SELECT id, created_at, updated_at, name
FROM users
WHERE id = $1;
