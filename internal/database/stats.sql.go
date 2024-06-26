// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: stats.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createStat = `-- name: CreateStat :one
INSERT INTO stats (id, created_at, updated_at, score)
VALUES ($1, $2, $3, $4)
RETURNING id, created_at, updated_at, score
`

type CreateStatParams struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Score     int32
}

func (q *Queries) CreateStat(ctx context.Context, arg CreateStatParams) (Stat, error) {
	row := q.db.QueryRowContext(ctx, createStat,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.Score,
	)
	var i Stat
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Score,
	)
	return i, err
}

const getMaxScore = `-- name: GetMaxScore :one
SELECT MAX(score) 
FROM stats
`

func (q *Queries) GetMaxScore(ctx context.Context) (interface{}, error) {
	row := q.db.QueryRowContext(ctx, getMaxScore)
	var max interface{}
	err := row.Scan(&max)
	return max, err
}

const getStats = `-- name: GetStats :many
SELECT id, created_at, updated_at, score FROM stats
ORDER BY created_at
`

func (q *Queries) GetStats(ctx context.Context) ([]Stat, error) {
	rows, err := q.db.QueryContext(ctx, getStats)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Stat
	for rows.Next() {
		var i Stat
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Score,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
