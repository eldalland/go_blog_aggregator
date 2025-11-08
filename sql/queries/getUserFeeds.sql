-- name: GetUserFeeds :many
WITH feedfollows AS (
    SELECT * FROM feed_follows
    WHERE feed_follows.user_id = $1
)
SELECT 
    feedfollows.*,
    feeds.name AS feed_name,
    users.name AS user_name
FROM feedfollows
INNER JOIN feeds ON feedfollows.feed_id = feeds.id
INNER JOIN users ON feedfollows.user_id = users.id;