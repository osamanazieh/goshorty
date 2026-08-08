-- name: CreateUrl :many
INSERT INTO shorty(id, url, short_code, created_at, updated_at) 
VALUES(
    $1,
    $2,
    $3,
    $4,
    $5
)
RETURNING id, url, short_code, created_at, updated_at;


-- name: GetUrl :one 
SELECT * FROM shorty
WHERE short_code = $1;

-- name: UpdateUrl :one
UPDATE shorty
SET url = $1, updated_at = $2
WHERE short_code = $3
RETURNING id, url, short_code, created_at, updated_at;

-- name: GetHits :one 
SELECT hits FROM shorty
WHERE short_code = $1; 

-- name: SetHits :exec
UPDATE shorty
SET hits = $1
WHERE short_code= $2;

-- name: DeleteUrl :execrows
DELETE FROM shorty
WHERE short_code = $1; 

