// Code generated by sqlc. DO NOT EDIT.
// source: query.sql

package db

import (
	"context"
	"database/sql"
)

const upsertAuthor = `-- name: UpsertAuthor :exec
INSERT INTO authors (name, bio)
VALUES (?, ?)
ON DUPLICATE KEY
    UPDATE bio = ?
`

type UpsertAuthorParams struct {
	Name  string
	Bio   sql.NullString
	Bio_2 sql.NullString
}

func (q *Queries) UpsertAuthor(ctx context.Context, arg UpsertAuthorParams) error {
	_, err := q.db.ExecContext(ctx, upsertAuthor, arg.Name, arg.Bio, arg.Bio_2)
	return err
}
