-- name: GetUserById :one
SELECT name
FROM users
WHERE id = $1;