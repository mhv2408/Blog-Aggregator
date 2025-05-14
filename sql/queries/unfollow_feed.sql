
-- name: UnfollowFeed :exec
DELETE 
FROM feed_follows ff
USING users, feed
WHERE feed_follows.user_id = users.id
AND feed_follows.feed_id = feed.id
AND users.name = $1 
AND feed.url = $2;
