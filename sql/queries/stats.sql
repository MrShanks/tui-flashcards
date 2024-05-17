-- name: CreateStat :one
INSERT INTO stats (id, created_at, updated_at, score)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetStats :many
SELECT * FROM stats
ORDER BY created_at;

-- name: GetMaxScore :one
SELECT MAX(score) 
FROM stats;