-- name: GetPostsForUser :many
SELECT * FROM posts
WHERE feed_id = ANY($1::uuid[])
ORDER BY published_at DESC
LIMIT $2;