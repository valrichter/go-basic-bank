-- name: CreateEntry :one
INSERT INTO entrie (
  account_id,
  amount
) VALUES (
  $1, $2
) RETURNING *;

-- name: GetEntry :one
SELECT * FROM entrie
WHERE id = $1 LIMIT 1;

-- name: ListEntrie :many
SELECT * FROM entrie
WHERE account_id = $1
ORDER BY id
LIMIT $2
OFFSET $3;