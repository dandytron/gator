-- name: CreateFeed :one

INSERT INTO feeds(id, created_at, updated_at, name, url, user_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
)
RETURNING *;

-- name: GetFeedByURL :one

SELECT *
FROM feeds
WHERE url = $1;

-- name: DeleteFeeds :exec

DELETE FROM feeds;

-- name: GetFeeds :many

SELECT feeds.name AS name, feeds.url AS url, users.name AS user
FROM feeds
INNER JOIN users
ON feeds.user_id = users.id;

-- name: MarkFeedFetched :one

UPDATE feeds
SET last_fetched_at = NOW(),
updated_at = NOW()
WHERE feeds.id = $1
RETURNING *;

-- name: GetNextFeedToFetch :one

SELECT *
FROM feeds
ORDER BY last_fetched_at ASC NULLS FIRST
LIMIT 1;