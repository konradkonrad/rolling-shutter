// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: meta-query.sql

package commondb

import (
	"context"
)

const getMeta = `-- name: GetMeta :one
SELECT value FROM meta_inf WHERE key = $1
`

func (q *Queries) GetMeta(ctx context.Context, key string) (string, error) {
	row := q.db.QueryRow(ctx, getMeta, key)
	var value string
	err := row.Scan(&value)
	return value, err
}

const insertMeta = `-- name: InsertMeta :exec
INSERT INTO meta_inf (key, value) VALUES ($1, $2)
`

type InsertMetaParams struct {
	Key   string
	Value string
}

func (q *Queries) InsertMeta(ctx context.Context, arg InsertMetaParams) error {
	_, err := q.db.Exec(ctx, insertMeta, arg.Key, arg.Value)
	return err
}
