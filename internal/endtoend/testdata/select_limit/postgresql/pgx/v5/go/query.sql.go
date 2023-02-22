// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.1
// source: query.sql

package querytest

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const fooLimit = `-- name: FooLimit :many
SELECT a FROM foo
LIMIT $1
`

func (q *Queries) FooLimit(ctx context.Context, limit int32) ([]pgtype.Text, error) {
	rows, err := q.db.Query(ctx, fooLimit, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []pgtype.Text
	for rows.Next() {
		var a pgtype.Text
		if err := rows.Scan(&a); err != nil {
			return nil, err
		}
		items = append(items, a)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const fooLimitOffset = `-- name: FooLimitOffset :many
SELECT a FROM foo
LIMIT $1 OFFSET $2
`

type FooLimitOffsetParams struct {
	Limit  int32
	Offset int32
}

func (q *Queries) FooLimitOffset(ctx context.Context, arg FooLimitOffsetParams) ([]pgtype.Text, error) {
	rows, err := q.db.Query(ctx, fooLimitOffset, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []pgtype.Text
	for rows.Next() {
		var a pgtype.Text
		if err := rows.Scan(&a); err != nil {
			return nil, err
		}
		items = append(items, a)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
