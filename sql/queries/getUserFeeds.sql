-- name: GetUserFeeds :many
WITH feedfollows AS (
    SELECT * FROM feed_follows
    WHERE feed_follows.user_id = $1
)
SELECT 
    feeds.id,
    feeds.created_at,
    feeds.updated_at,
    feeds.name,
    feeds.url,
    feeds.user_id,
    feeds.last_fetched_at
FROM feedfollows
INNER JOIN feeds ON feedfollows.feed_id = feeds.id
INNER JOIN users ON feedfollows.user_id = users.id;