-- name: FeedByURL :one
SELECT *
FROM feed
WHERE url=$1;
