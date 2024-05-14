-- name: CreateStat :one
INSERT INTO stats (id, created_at, updated_at, score)
VALUES ($1, $2, $3, $4)
RETURNING *;