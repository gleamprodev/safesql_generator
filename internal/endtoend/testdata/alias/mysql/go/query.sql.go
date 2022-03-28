// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.12.0
// source: query.sql

package querytest

import (
	"context"
)

const aliasBar = `-- name: AliasBar :exec
DELETE FROM bar b
WHERE b.id = ?
`

func (q *Queries) AliasBar(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, aliasBar, id)
	return err
}
